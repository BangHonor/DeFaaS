// SPDX-License-Identifier: MIT

pragma solidity^0.6.0;

import "./FaaSTokenPay.sol";
import "./ProviderPool.sol";
import "./SimpleAuction.sol";
import "./WitnessPool.sol";
import "./SLA.sol";
import "./SafeMath.sol";
import "./Owned.sol";
import "./FaaSLevel.sol";

contract Market is Owned, FaaSTokenPay, FaaSLevel, ProviderPool {

    // ------------------------------------------------------------------------------------

    using SafeMath for uint;   // 注意到 uint 默认是 uint256

    // ------------------------------------------------------------------------------------

    // 部署订单的有限状态机
    enum OrderStates {
        Bidding,    // 竞价中，以产生新的部署订单和新的竞价为开始，以竞价结束为结束
        Confirming, // 确认中，以竞价匹配成功为开始，以租户和供应商都确认并提供部署信息为结束
        Deploying,  // 部署中，以部署信息填充完毕，租户获取部署信息为开始，租户向供应商发出部署请求，以租户触发Fulfilling状态为结束
        Fulfilling, // 履行中，以触发履行Fulfilling为开始，证人监督供应商，以任意人触发Settled状态完成结算为结束
        Settled,    // 已结算，以触发Settled状态为开始，此状态中账目已结算清楚但还没有转账，以任意人触发Finished状态完成转账为结束
        Finished    // 已结束，为终止状态，以触发Finished状态为开始，此状态中已完成转账
    }
 
    function isAtOrderStates(uint _deploymentOrderID, OrderStates _state) private view returns(bool) {
        return (deploymentOrders[_deploymentOrderID].state == _state);
    }

    modifier atOrderState(uint _deploymentOrderID, OrderStates _state) {
        require(
            isAtOrderStates(_deploymentOrderID, _state),
            "Matket: function cannot be called at this state"
        );
        _;
    }

    // ------------------------------------------------------------------------------------

    // 部署订单
    struct DeploymentOrder {
        // 状态
        OrderStates state;        // 状态
        // 订单内容
        address customer;         // 租户地址
        uint faaSLevelID;         // FaaS 规格
        uint faaSDuration;        // 服务时长
        uint highestUnitPrice;    // 租户可接受的最高单价
        // 竞价
        SimpleAuction auctionContract;  // 竞价合约
        bool isMatched;                 // 是否在竞价中匹配成功
    }

    // 部署信息
    struct DeploymentInfo {
        // 状态信息
        address provider;          // 供应商
        bool isProviderConfirmed;  // 供应商已确认
        // 部署地址
        string endAddr;            // 供应商提供的部署端地址
        string funcPath;           // 租户提供的函数路径
        string deployServerAddr;   // 供应商提供的接受部署的服务器地址
        string accessSecretKey;    // 供应商提供给租户的访问密钥，以租户的公钥加密
    }

    // 租约
    struct Lease {
        // 双方
        address customer;  // 租户地址
        address provider;  // 供应商地址
        // 服务内容
        uint faaSLevelID;  // 租用的 FaaS 规格
        uint faaSDuration; // 租用的时间
        uint unitPrice;    // FaaS 规格对应的单价
        // 费用
        uint providerServiceFee;       // 服务成功时，供应商的服务费用
        uint customerLockFee;          // 租户的预付款
        uint customerWithdrawFee;      // 租户取回多出的预付款
        uint customerCompensationFee;  // 服务失败时，租户的补偿款
        uint witnessFee;               // 证人的费用
        uint maintainerFee;            // 区块链维护者的费用
        // 监督
        bool isViolatedSLA;            // 是否违反 SLA 合约
        address SLAContractAddress;    // SLA 执行合约的地址
    }

    // ------------------------------------------------------------------------------------

    // 部署订单数量
    uint private numDeploymentOrders;

    // 部署订单表：deploymentOrderID => DeploymentOrder
    // 部署信息表：deploymentOrderID => DeploymentInfo
    // 租约表:     deploymentOrderID => Lease
    mapping(uint => DeploymentOrder) private deploymentOrders;
    mapping(uint => DeploymentInfo)  private deploymentInfos;
    mapping(uint => Lease)           private leases;

    WitnessPool wpContract;
    
    // ------------------------------------------------------------------------------------

    // 新建部署订单事件
    // 部署订单竞价结束事件
    // 新建部署信息事件
    // 新建租约事件
    event NewDeploymentOrderEvent(uint indexed _deploymentOrderID, address _auctionAddress);
    event AuctionEndEvent        (uint indexed _deploymentOrderID, bool _success, address _provider, uint _unitPrice);
    event NewDeploymentInfoEvent (uint indexed _deploymentOrderID, address _customer, address _provider);
    event NewLeaseEvent          (uint indexed _deploymentOrderID, address _customer, address _provider);

    // ------------------------------------------------------------------------------------

    constructor(address _tokenContractAddress, address _witnessPoolContractAddress) 
        ProviderPool(_tokenContractAddress)
        public
    {        
        // 部署订单参数初始化
        // 部署订单号计数， 有效的部署订单号从 1 开始，0 为无效的部署订单号
        numDeploymentOrders = 1;

        // 交互合约初始化
        wpContract = WitnessPool(_witnessPoolContractAddress);
        // wpContract.setMarketContractAddress(address(this));  // 部署时动作，要求 Market 的部署者和 WitnessPool 的所有者相同
    }
  
    // ------------------------------------------------------------------------------------

    // 租户 API
    // 计算订单的预付款
    function calculateLockFee(uint _highestUnitPrice, uint _faaSDuration) public pure returns (uint) {
        return (_highestUnitPrice * _faaSDuration);
    }

    // API
    // 查询部署订单
    // 返回：（租户，FaaS规格，服务时长，租户可接受的最高出价）
    function getDeploymentOrder(uint _deploymentOrderID) 
        public
        view
        returns (address, uint, uint, uint)
    {
        // 对于引用类型，状态变量向 storage 局部变量赋值时仅仅传递一个引用
        DeploymentOrder storage _order = deploymentOrders[_deploymentOrderID];
        
        return (
            _order.customer,
            _order.faaSLevelID,
            _order.faaSDuration,
            _order.highestUnitPrice
        );
    }

    // 退还预付款
    function freeLockFee(uint _deploymentOrderID) private {
        require(
            isAtOrderStates(_deploymentOrderID, OrderStates.Bidding) || 
            isAtOrderStates(_deploymentOrderID, OrderStates.Confirming) || 
            isAtOrderStates(_deploymentOrderID, OrderStates.Deploying),
            "Market: not a state that can free locked fee"
        );

        DeploymentOrder storage _order = deploymentOrders[_deploymentOrderID];
        uint lockFee = calculateLockFee(_order.highestUnitPrice,  _order.faaSDuration);
        require(
            tokenContract.transfer(_order.customer, lockFee),
            "Market: failed to refund locked fee back to customer"
        );
    }

    // ------------------------------------------------------------------------------------

    // 租户 API
    // 新建部署订单
    function newDeploymentOrder(
        uint _faaSLevelID,
        uint _highestUnitPrice,
        uint _faaSDuration,
        uint _biddingDuration) 
        public 
        returns (uint) 
    {
        // 参数检查
        require(_faaSLevelID < getFaaSLevelNumber(), "");
        require(_faaSDuration < 100 days, "");
        require(_biddingDuration < 1 hours, "");

        // 锁定预付款
        uint lockFee = calculateLockFee(_highestUnitPrice, _faaSDuration);
        require(
            tokenContract.transferFrom(msg.sender, address(this), lockFee) == true,
            "lock fee before creating deployment order"
        );

        // 新订单的 ID
        uint _deploymentOrderID = numDeploymentOrders++;

        // 为新订单创建竞价合约
        SimpleAuction _auction = new SimpleAuction(_highestUnitPrice,_biddingDuration);

        // 新订单
        deploymentOrders[_deploymentOrderID] = DeploymentOrder({
            state: OrderStates.Bidding,
            customer: msg.sender,
            faaSLevelID: _faaSDuration,
            highestUnitPrice: _highestUnitPrice,
            faaSDuration: _faaSDuration,
            auctionContract: _auction,
            isMatched: false
        });

        // 产生事件
        emit NewDeploymentOrderEvent(_deploymentOrderID, address(_auction));

        return (_deploymentOrderID);
    }

    // 供应商 API
    // 对部署订单竞价
    function bid(uint _deploymentOrderID, uint _unitPrice) 
        public
        providerRegistered(msg.sender)
        providerQualified(msg.sender)
        atOrderState(_deploymentOrderID, OrderStates.Bidding)
    {
        deploymentOrders[_deploymentOrderID].auctionContract.bid(msg.sender, _unitPrice);
    }

    // 租户 API
    // 检查部署订单的匹配结果
    // 返回：（是否匹配成功）
    function matchDeploymentOrder(uint _deploymentOrderID) 
        public
        atOrderState(_deploymentOrderID, OrderStates.Bidding)
        returns (bool) 
    {
        bool _success;
        address _provider;
        uint _unitPrice;

        (_success, _provider, _unitPrice) = deploymentOrders[_deploymentOrderID].auctionContract.auctionEnd();
        emit AuctionEndEvent(_deploymentOrderID, _success, _provider, _unitPrice);
        
        // 匹配失败
        if (_success == false) {
                                                                                // 检查
            deploymentOrders[_deploymentOrderID].state = OrderStates.Finished;  // 生效：修改状态
            freeLockFee(_deploymentOrderID);                                    // 交互：退还预付款

            return (false);
        }

        // 匹配成功, 创建新的部署信息
        deploymentInfos[_deploymentOrderID] = DeploymentInfo({
            provider: _provider,
            isProviderConfirmed: false,
            endAddr: "",
            funcPath: "",
            deployServerAddr: "",
            accessSecretKey: ""
        });
        // 修改状态
        deploymentOrders[_deploymentOrderID].state = OrderStates.Confirming;

        return (true);
    }


    // 供应商 API
    // 供应商确认部署订单，同时提交部署信息
    function providerConfirm(
        uint _deploymentOrderID,
        string memory _endAddr,
        string memory _deployServerAddr,
        string memory _accessSecretKey) 
        public
        atOrderState(_deploymentOrderID, OrderStates.Confirming)
    {
        require(
            msg.sender == deploymentInfos[_deploymentOrderID].provider,
            "Matket: only provider can confirm"
        );
        deploymentInfos[_deploymentOrderID].endAddr = _endAddr;
        deploymentInfos[_deploymentOrderID].deployServerAddr = _deployServerAddr;
        deploymentInfos[_deploymentOrderID].accessSecretKey = _accessSecretKey;
        deploymentInfos[_deploymentOrderID].isProviderConfirmed = true;
    }

    // 租户 API
    // 在供应商确认部署订单之后，租户进行部署订单确认，同时获得部署信息
    function customerConfirm(uint _deploymentOrderID, string memory _funcPath) 
        public
        atOrderState(_deploymentOrderID, OrderStates.Confirming)
    {
        require(
            msg.sender == deploymentOrders[_deploymentOrderID].customer,
            "Matket: only customer can confirm"
        );

        // 要求供应商确认在前
        require(deploymentInfos[_deploymentOrderID].isProviderConfirmed == true,
            "Market：provider has not confirmed yet"
        );

        // 提供部署信息
        deploymentInfos[_deploymentOrderID].funcPath = _funcPath;

        // 产生事件
        emit NewDeploymentInfoEvent(_deploymentOrderID, deploymentOrders[_deploymentOrderID].customer, deploymentInfos[_deploymentOrderID].provider);
        
        // TODO 新建租约
        emit NewLeaseEvent(_deploymentOrderID, deploymentOrders[_deploymentOrderID].customer, deploymentInfos[_deploymentOrderID].provider);

        // 修改状态
        deploymentOrders[_deploymentOrderID].state = OrderStates.Deploying;
    }

    // function () {
    //             Lease memory _lease = newLease(_deploymentOrderID, _provider, _unitPrice);
    //     leases[_deploymentOrderID] = _lease;
    //     emit NewLeaseEvent(_deploymentOrderID, _lease.customer, _lease.provider);
    // }

    // // 租户部署完成，开始履行部署订单
    // function fulfillDeploymentOrder(uint _deploymentOrderID, string memory something)
    //     public
    //     atOrderState(_deploymentOrderID, OrderStates.Deploying)
    // {
    //     // 产生 SLA 执行合约
    //     // TODO：详细的 SLA 内容 something
    //     leases[_deploymentOrderID].SLAContractAddress = wpContract.genSLAContract();

    //     // 修改订单状态
    //     deploymentOrderStates[_deploymentOrderID] = OrderStates.Fulfilling;
    // }

    // // 按照 SLA 的返回结果对订单进行结算
    // // 履行时间结束，供应商完成部署订单
    // function settleDeploymentOrder(uint _deploymentOrderID) 
    //     public
    //     atOrderState(_deploymentOrderID, OrderStates.Fulfilling)
    // {
    //     // TODO：控制时间

    //     // 查看是否违约
    //     leases[_deploymentOrderID].isViolatedSLA = CloudSLA(leases[_deploymentOrderID].SLAContractAddress).isViolatedSLA();

    //     // 修改订单状态：进入已结算（未转账）的状态
    //     deploymentOrderStates[_deploymentOrderID] = OrderStates.Settled;
    // }


    // // 结束部署订单，完成转账
    // function finishDeploymemtOrder(uint _deploymentOrderID) 
    //     public
    //     atOrderState(_deploymentOrderID, OrderStates.Settled)
    // {
    //     Lease memory _lease = leases[_deploymentOrderID];

    //     // 支付区块链维护者费用 TODO
    //     address _maintainerAddress = address(0);
    //     require(tokenContract.transfer(_maintainerAddress, _lease.maintainerFee), "");

    //     // 支付证人费用
    //     require(tokenContract.transfer(address(wpContract), _lease.witnessFee), "");
        
    //     // 退回租户预付款多出的部分
    //     require(tokenContract.transfer(_lease.customer, _lease.customerWithdrawFee), "");  
        
    //     if (_lease.isViolatedSLA == true) {
    //         // 如违约，退回租户补偿费
    //         require(tokenContract.transfer(_lease.customer, _lease.customerWithdrawFee), "");
    //     } else {
    //         //  如不违约，支付供应商报酬
    //         require(tokenContract.transfer(_lease.provider, _lease.providerServiceFee), "");
    //     }

    //     // 结算完成的订单状态为 Finished
    //     deploymentOrderStates[_deploymentOrderID] = OrderStates.Finished;
    // }


    // // ------------------------------------------------------------------------------------

    // // 新建租约
    // // 返回租约ID
    // function newLease(
    //     uint _deploymentOrderID,
    //     address _provider,
    //     uint _unitPrice) 
    //     internal
    //     returns (Lease memory)
    // {
    //     uint _leaseID = _deploymentOrderID;  // leaseID 就是 deploymentOrderID
    //     DeploymentOrder memory order = deploymentOrders[_deploymentOrderID];

    //     // 租户的服务款
    //     uint _customerLockFee =  getDeploymentOrderLockFee(order.highestUnitPrice, order.faaSDuration);
    //     // 租户的实际支付
    //     uint _customerPayFee = order.faaSDuration * _unitPrice;
    //     // 租户多出的预付款
    //     uint _customerWithdrawFee = _customerLockFee - _customerPayFee;
    //     // 证人费用 TODO
    //     uint _witnessFee = 0;
    //     // 区块链维护者费用 TODO
    //     uint _maintainerFee = 0;
    //     // 服务成功时，供应商可得服务报酬
    //     uint _providerServiceFee = _customerPayFee - (_witnessFee + _maintainerFee);
    //     // 服务失败时，租户可得的补偿款
    //     uint _customerCompensationFee = _providerServiceFee;

    //     Lease memory _lease = Lease({
    //         customer: order.customer,
    //         provider: _provider,
    //         faaSLevelID: order.faaSLevelID,
    //         faaSDuration: order.faaSDuration,
    //         unitPrice: _unitPrice,
    //         isViolatedSLA: false,
    //         providerServiceFee: _providerServiceFee,
    //         customerLockFee: _customerLockFee,
    //         customerWithdrawFee: _customerWithdrawFee,
    //         customerCompensationFee: _customerCompensationFee,
    //         witnessFee: _witnessFee,
    //         maintainerFee: _maintainerFee,
    //         SLAContractAddress: address(0)  // 在本函数中填充为 0
    //     });
        
    //     return _lease;
    // }
}
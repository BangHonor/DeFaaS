// SPDX-License-Identifier: MIT

pragma solidity>=0.6.0;

import "./FaaSTokenPay.sol";
import "./ProviderManagement.sol";
import "./SimpleAuction.sol";
import "./WitnessPool.sol";
import "./SafeMath.sol";
import "./Owned.sol";
import "./FaaSLevel.sol";

contract Market is Owned, FaaSTokenPay, FaaSLevel, ProviderManagement {

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
        uint faaSDuration;        // 服务时长，该值以秒为单位
        uint highestUnitPrice;    // 租户可接受的最高单价
        // 竞价
        SimpleAuction auctionContract;  // 竞价合约
        bool isMatched;                 // 是否在竞价中匹配成功
    }

    // 部署信息
    struct DeploymentInfo {
        // 状态信息
        uint unitPrice;            // 竞价中成交的单价
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
        uint faaSLevelID;      // 租用的 FaaS 规格
        uint faaSDuration;     // 租用的时长
        uint unitPrice;        // FaaS 规格对应的单价
        uint faaSFulfillTime;  // FaaS 履行开始的时间
        // 费用
        uint providerServiceFee;       // 服务成功时，供应商的服务费用
        uint customerWithdrawFee;      // 租户取回多出的预付款
        uint customerCompensationFee;  // 服务失败时，租户的补偿款
        uint witnessFee;               // 证人的费用
        uint maintainerFee;            // 区块链维护者的费用
        // 监督
        bool isViolatedSLA;            // 是否违反 SLA 合约
        uint curBlockNum;              // 租约产生时的块号
    }

    // ------------------------------------------------------------------------------------

    // 部署订单数量，并且是部署订单 ID 自增器
    uint private numDeploymentOrders;

    // 部署订单表: deploymentOrderID => DeploymentOrder
    // 部署信息表: deploymentOrderID => DeploymentInfo
    // 租约表:     deploymentOrderID => Lease
    mapping(uint => DeploymentOrder) private deploymentOrders;
    mapping(uint => DeploymentInfo)  private deploymentInfos;
    mapping(uint => Lease)           private leases;

    // WitnessPool 合约
    WitnessPool public wpContract;

    // 维护者地址
    address private maintainerAddress = address(0x1);
    
    // ------------------------------------------------------------------------------------

    // 新建部署订单事件
    // 部署订单竞价结束事件
    // 新建部署信息事件
    // 新建租约事件
    // 新建 SLA 监督事件
    event NewDeploymentOrderEvent(uint indexed _deploymentOrderID, address _auctionAddress);
    event AuctionEndEvent        (uint indexed _deploymentOrderID, bool _success, address _provider, uint _unitPrice);
    event NewDeploymentInfoEvent (uint indexed _deploymentOrderID);
    event NewLeaseEvent          (uint indexed _deploymentOrderID, address _customer, address _provider);
    event NewSLAEvent            (uint indexed _deploymentOrderID);

    // ------------------------------------------------------------------------------------

    constructor(address _tokenContractAddress) 
        ProviderManagement(_tokenContractAddress)
    {        
        // 部署订单参数初始化
        numDeploymentOrders = 0;

        // 创建 WitnessPool 合约，Market 合约是 WitnessPool 合约的所有者
        wpContract = new WitnessPool(_tokenContractAddress);
    }
  
    // ------------------------------------------------------------------------------------

    modifier validDeploymentOrder(uint _deploymentOrderID) {
        require(
            _deploymentOrderID < numDeploymentOrders, 
            "Market: invalid deployment order ID"
        );
        _;
    }

    // ------------------------------------------------------------------------------------


    // 查询部署订单
    // 返回：（租户，FaaS规格，服务时长，租户可接受的最高出价）
    function getDeploymentOrder(uint _deploymentOrderID) 
        public
        view
        validDeploymentOrder(_deploymentOrderID)
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

    // 计算部署订单的服务费
    function calculateServiceFee(uint _unitPrice, uint _faaSDuration) public pure returns (uint) {
        // _faaSDuration 以秒为单位
        // FaaS 费用以小时计价，不足整个小时部分按一小时计算
        return (_unitPrice * ( (_faaSDuration + (1 hours) - 1) / (1 hours)) );
    }

    // 计算部署订单的预付款：锁定双倍费用
    function calculateLockFee(uint _highestUnitPrice, uint _faaSDuration) public pure returns (uint) {
        return ( 2 * calculateServiceFee(_highestUnitPrice, _faaSDuration));
    }

    // 退还预付款
    function _freeLockFee(uint _deploymentOrderID) private {
        require(
            isAtOrderStates(_deploymentOrderID, OrderStates.Bidding) || 
            isAtOrderStates(_deploymentOrderID, OrderStates.Confirming) || 
            isAtOrderStates(_deploymentOrderID, OrderStates.Deploying),
            "Market: not a state that can free locked fee"
        );

        DeploymentOrder storage _order = deploymentOrders[_deploymentOrderID];
        require(
            tokenContract.transfer(_order.customer, calculateLockFee(_order.highestUnitPrice,  _order.faaSDuration)),
            "Market: failed to refund locked fee back to customer"
        );
    }

    // ------------------------------------------------------------------------------------

    // 租户 API
    // 新建部署订单
    // 返回：（部署订单 ID）
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
        require(_highestUnitPrice > 0, "");
        require(_faaSDuration >= 1 days &&  _faaSDuration <= 100 days, "");
        require(_biddingDuration >= 1 hours &&  _biddingDuration <= 7 days, "");

        // 锁定预付款
        require(
            tokenContract.transferFrom(msg.sender, address(this), calculateLockFee(_highestUnitPrice, _faaSDuration)),
            "Market: failed to lock fee before creating deployment order"
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
        validDeploymentOrder(_deploymentOrderID)
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
            _freeLockFee(_deploymentOrderID);                                    // 交互：退还预付款

            return (false);
        }

        // 匹配成功, 创建新的待供应商填充的部署信息
        deploymentInfos[_deploymentOrderID] = DeploymentInfo({
            unitPrice: _unitPrice,
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
        string memory _funcPath,
        string memory _deployServerAddr,
        string memory _accessSecretKey) 
        public
        validDeploymentOrder(_deploymentOrderID)
        atOrderState(_deploymentOrderID, OrderStates.Confirming)
    {
        require(
            msg.sender == deploymentInfos[_deploymentOrderID].provider,
            "Matket: only provider can confirm"
        );

        // 参数检查
        require(bytes(_endAddr).length          > 0 && bytes(_endAddr).length          <= 256, "");
        require(bytes(_funcPath).length         > 0 && bytes(_funcPath).length         <= 256, "");
        require(bytes(_deployServerAddr).length > 0 && bytes(_deployServerAddr).length <= 256, "");
        require(bytes(_accessSecretKey).length  > 0 && bytes(_accessSecretKey).length  <= 256, "");

        deploymentInfos[_deploymentOrderID].endAddr = _endAddr;
        deploymentInfos[_deploymentOrderID].funcPath = _funcPath;
        deploymentInfos[_deploymentOrderID].deployServerAddr = _deployServerAddr;
        deploymentInfos[_deploymentOrderID].accessSecretKey = _accessSecretKey;

        // 确认
        deploymentInfos[_deploymentOrderID].isProviderConfirmed = true;

        emit NewDeploymentInfoEvent(_deploymentOrderID);
    }

    // 租户 API
    // 在供应商确认部署订单之后，租户进行部署订单确认，同时获得部署信息
    // 返回（_endAddr，_funcPath，_deployServerAddr，_accessSecretKey）
    function customerConfirm(uint _deploymentOrderID) 
        public
        validDeploymentOrder(_deploymentOrderID)
        atOrderState(_deploymentOrderID, OrderStates.Confirming)
        returns(string memory, string memory, string memory, string memory)
    {
        require(
            msg.sender == deploymentOrders[_deploymentOrderID].customer,
            "Market: only customer can confirm");

        // 要求供应商确认在前
        require(deploymentInfos[_deploymentOrderID].isProviderConfirmed == true,
            "Market: provider has not confirmed yet");

        // 新建租约
        _newLease(_deploymentOrderID);
        // 产生事件
        emit NewLeaseEvent(_deploymentOrderID, deploymentOrders[_deploymentOrderID].customer, deploymentInfos[_deploymentOrderID].provider);
        
        // 修改状态
        deploymentOrders[_deploymentOrderID].state = OrderStates.Deploying;

        DeploymentInfo storage _info = deploymentInfos[_deploymentOrderID];
        return (_info.endAddr, _info.funcPath, _info.deployServerAddr, _info.accessSecretKey);
    }

    // 租户 API
    // 租户部署完成，开始履行部署订单
    function fulfillDeploymentOrder(uint _deploymentOrderID)
        public
        validDeploymentOrder(_deploymentOrderID)
        atOrderState(_deploymentOrderID, OrderStates.Deploying)
    {
        // 生成 SLA 监督，由证人执行对供应商的监视
        wpContract.newSLA(
            leases[_deploymentOrderID].curBlockNum,
            deploymentInfos[_deploymentOrderID].provider, 
            deploymentOrders[_deploymentOrderID].customer,
            _deploymentOrderID,
            deploymentInfos[_deploymentOrderID].funcPath,
            deploymentOrders[_deploymentOrderID].faaSDuration - 1 hours);
        
        emit NewSLAEvent(_deploymentOrderID);

        leases[_deploymentOrderID].faaSFulfillTime = block.timestamp;

        // 修改订单状态
        deploymentOrders[_deploymentOrderID].state= OrderStates.Fulfilling;
    }

    // 按照 SLA 的返回结果对订单进行结算
    // 履行时间结束，供应商完成部署订单
    function settleDeploymentOrder(uint _deploymentOrderID) 
        public
        validDeploymentOrder(_deploymentOrderID)
        atOrderState(_deploymentOrderID, OrderStates.Fulfilling)
    {
        require(
            block.timestamp > leases[_deploymentOrderID].faaSFulfillTime + leases[_deploymentOrderID].faaSDuration,
            "Matket: FaaS is still fulfilling"
        );

        // 查看是否违约
        leases[_deploymentOrderID].isViolatedSLA = wpContract.isViolatedSLA(_deploymentOrderID);

        // 修改订单状态：进入已结算而未转账的状态
        deploymentOrders[_deploymentOrderID].state = OrderStates.Settled;
    }


    // 结束部署订单，完成转账
    function finishDeploymemtOrder(uint _deploymentOrderID) 
        public
        validDeploymentOrder(_deploymentOrderID)
        atOrderState(_deploymentOrderID, OrderStates.Settled)
    {
        Lease storage _lease = leases[_deploymentOrderID];

        // 支付区块链维护者费用
        require(tokenContract.transfer(maintainerAddress, _lease.maintainerFee), "");

        // 支付证人费用
        require(tokenContract.transfer(address(wpContract), _lease.witnessFee), "");
        
        // 退回租户预付款多出的部分
        require(tokenContract.transfer(_lease.customer, _lease.customerWithdrawFee), "");  
        
        if (_lease.isViolatedSLA == true) {
            // 如违约，退回租户补偿款
            require(tokenContract.transfer(_lease.customer, _lease.customerWithdrawFee), "");
        } else {
            //  如不违约，支付供应商报酬
            require(tokenContract.transfer(_lease.provider, _lease.providerServiceFee), "");
        }

        // 转账完成的订单状态为 Finished，整个订单流程结束
        deploymentOrders[_deploymentOrderID].state = OrderStates.Finished;
    }


    // ------------------------------------------------------------------------------------

    function _newLease(uint _deploymentOrderID) 
        private  
        atOrderState(_deploymentOrderID, OrderStates.Confirming)
    {
        DeploymentOrder storage _order = deploymentOrders[_deploymentOrderID];
        DeploymentInfo  storage _info  = deploymentInfos[_deploymentOrderID];

        uint _customerLockFee = calculateLockFee(_order.highestUnitPrice, _order.faaSDuration);  // 租户的预付款
        uint _customerPayFee = calculateServiceFee(_info.unitPrice, _order.faaSDuration );       // 租户的实际应付                 
        uint _customerWithdrawFee = _customerLockFee - _customerPayFee;                          // 租户应取回的预付款
        uint _witnessFee = 0;                                                                    // 证人费用 TODO
        uint _maintainerFee = 0;                                                                 // 区块链维护者费用 TODO
        uint _providerServiceFee = _customerPayFee - (_witnessFee + _maintainerFee);             // 服务成功时，供应商可得服务报酬
        uint _customerCompensationFee = _providerServiceFee;                                     // 服务失败时，租户可得的补偿款

        // 新建租约
        leases[_deploymentOrderID] = Lease({
            //
            customer: _order.customer,
            provider: _info.provider,
            //
            faaSLevelID: _order.faaSLevelID,
            faaSDuration: _order.faaSDuration,
            unitPrice: _info.unitPrice,
            faaSFulfillTime: 0,            // 待在 Deploying 状态填写
            //
            providerServiceFee: _providerServiceFee,
            customerWithdrawFee: _customerWithdrawFee,
            customerCompensationFee: _customerCompensationFee,
            witnessFee: _witnessFee,
            maintainerFee: _maintainerFee,
            //
            isViolatedSLA: false,          // 待在 Fulfilling 状态填写
            curBlockNum: block.number
        });
    }
}
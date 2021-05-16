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
        Fulfilling, // 履行中，以触发履行Fulfilling为开始，证人监督供应商，以任意人触发Finished状态完成结算为结束
        Finished    // 已结束，以触发Finished状态开始,结算部署订单账单并完成转账，为终止状态
    }

    // 状态检查
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
        // 部署订单内容
        address customer;         // 租户地址
        uint faasLevelID;         // FaaS 规格
        uint faasDuration;        // 服务时长，保存的是秒数
        uint highestUnitPrice;    // 租户可接受的最高单价
        string publicKey;         // 公钥
        // 竞价
        SimpleAuction auctionContract;  // 竞价合约
        bool isMatch;
    }

    // 部署信息
    struct DeploymentInfo {
        // 状态
        uint unitPrice;            // 竞价中成交的单价
        bool isProviderPublish;    // 供应商是否已经发布部署信息
        // 部署信息内容
        address provider;          // 供应商
        string funcPath;           // 函数的访问路径
        string deployPath;         // 部署的访问路径
        string accessKey;          // 供应商提供给租户的访问密钥，以租户的公钥加密
    }

    // 租约
    struct Lease {
        // 费用
        uint providerServiceFee;       // 服务成功时，供应商的服务费用
        uint customerWithdrawFee;      // 租户取回多出的预付款
        uint customerCompensationFee;  // 服务失败时，租户的补偿款
        uint platformChargeFee;        // 平台手续费
        // 服务内容
        bytes32 fulfillKey;               // 用于确认租户在供应商上完成部署
        uint faasFulfillStartTime;     // 供应商开始履行服务的时间
        // 监测
        // TODO 添加一些监测参数
        bool isViolatedSLA;            // 是否违反 SLA 合约
        uint curBlockNum;              // 租约产生时的块号
    }

    // ------------------------------------------------------------------------------------

    // WitnessPool 合约
    WitnessPool public wpContract;

    // 部署订单数量，并且是部署订单 ID 自增器
    uint private numDeploymentOrder;

    modifier validDeploymentOrderID(uint _deploymentOrderID) {
        require(
            _deploymentOrderID < numDeploymentOrder, 
            "Market: invalid deployment order ID"
        );
        _;
    }

    // ------------------------------------------------------------------------------------

    // 部署订单表: deploymentOrderID => DeploymentOrder
    // 部署信息表: deploymentOrderID => DeploymentInfo
    // 租约表:     deploymentOrderID => Lease
    mapping(uint => DeploymentOrder) private deploymentOrders;
    mapping(uint => DeploymentInfo)  private deploymentInfos;
    mapping(uint => Lease)           private leases;

    // ------------------------------------------------------------------------------------

    // 新建部署订单事件
    event NewDeploymentOrderEvent(
        address indexed _customer, 
        uint indexed _nonce,
        uint indexed _deploymentOrderID, 
        uint _faasLevelID,
        uint _highestUnitPrice,
        uint _faasDuration,
        uint _biddingDuration,
        string _publicKey);
    
    // 部署订单竞价结束事件
    event BiddingEndEvent(
        uint indexed _deploymentOrderID,
        address indexed _provider,
        bool indexed _success, 
        uint _faasLevelID,
        uint _unitPrice);

    // 新建部署信息事件
    event NewDeploymentInfoEvent (
        uint indexed _deploymentOrderID,
        address indexed _provider,
        string funcPath,
        string deployPath,
        string accessKey);

    // 新建租约事件
    event NewLeaseEvent(
        uint indexed _deploymentOrderID, 
        address indexed _customer, 
        address indexed _provider);
    
    // 新建 SLA 事件
    event NewSLAEvent(
        uint indexed _deploymentOrderID);

    
    // 结束部署订单事件
    event FinishDeploymentOrderEvent(
        uint indexed _deploymentOrderID,
        bool isViolatedSLA);


    // ------------------------------------------------------------------------------------

    constructor(address _tokenContractAddress) 
        ProviderManagement(_tokenContractAddress)
    {        
        // 部署订单参数初始化
        numDeploymentOrder = 0;

        // 创建 WitnessPool 合约，Market 合约是 WitnessPool 合约的所有者
        wpContract = new WitnessPool(_tokenContractAddress);
    }
  
    // ------------------------------------------------------------------------------------
    // 写合约
    // ------------------------------------------------------------------------------------

    // 租户 API
    // 新建部署订单
    // nonce 用于区别同一租户的不同新建请求
    function newDeploymentOrder(
        uint _nonce,
        uint _faasLevelID,
        uint _highestUnitPrice,
        uint _faasDuration,
        uint _biddingDuration,
        string memory _publicKey) 
        public 
        validFaaSLevelID(_faasLevelID)
    {
        // 参数检查
        // require(_highestUnitPrice > 0, "");
        // require(_faasDuration >= 12 hours &&  _faasDuration <= 360 days, "");
        // require(_biddingDuration >= 1 hours &&  _biddingDuration <= 1 days, "");
        // require(bytes(_publicKey).length > 0 && bytes(_publicKey).length < 256);

        // 锁定预付款
        require(
            tokenContract.transferFrom(msg.sender, address(this), calculateLockFee(_highestUnitPrice, _faasDuration)),
            "Market: failed to lock fee before creating deployment order"
        );

        // 新订单的 ID
        uint _deploymentOrderID = numDeploymentOrder++;

        // 为新订单创建竞价合约
        SimpleAuction _auction = new SimpleAuction(_highestUnitPrice, _biddingDuration);

        // 新订单
        deploymentOrders[_deploymentOrderID] = DeploymentOrder({
            state: OrderStates.Bidding,
            customer: msg.sender,
            faasLevelID: _faasDuration,
            highestUnitPrice: _highestUnitPrice,
            faasDuration: _faasDuration,
            publicKey: _publicKey,
            auctionContract: _auction,
            isMatch: false
        });

        DeploymentOrder storage _order = deploymentOrders[_deploymentOrderID];

        emit NewDeploymentOrderEvent(
            _order.customer, 
            _nonce,
            _deploymentOrderID, 
            _order.faasLevelID,
            _order.highestUnitPrice,
            _order.faasDuration,
            _biddingDuration,
            _order.publicKey);
    }

    // 供应商 API
    // 对部署订单竞价
    function bid(uint _deploymentOrderID, uint _unitPrice) 
        public
        validDeploymentOrderID(_deploymentOrderID)
        providerRegistered(msg.sender)
        providerQualified(msg.sender)
        atOrderState(_deploymentOrderID, OrderStates.Bidding)
    {
        deploymentOrders[_deploymentOrderID].auctionContract.bid(msg.sender, _unitPrice);
    }

    // 租户 API
    // 匹配部署订单
    function matchDeploymentOrder(uint _deploymentOrderID) 
        public
        validDeploymentOrderID(_deploymentOrderID)
        atOrderState(_deploymentOrderID, OrderStates.Bidding)
    {
        DeploymentOrder storage _order = deploymentOrders[_deploymentOrderID];

        bool _success;
        address _provider;
        uint _unitPrice;

        (_success, _provider, _unitPrice) = deploymentOrders[_deploymentOrderID].auctionContract.auctionEnd();

        emit BiddingEndEvent(
            _deploymentOrderID,
            _provider,
            _success, 
            _order.faasLevelID,
            _unitPrice);
        
        // 匹配失败
        if (_success == false) {
                                                                                 // 检查
            deploymentOrders[_deploymentOrderID].state = OrderStates.Finished;  // 生效：修改状态
            deploymentOrders[_deploymentOrderID].isMatch = false;
            _freeLockFee(_deploymentOrderID);                                   // 交互：退还预付款

            return;   // 返回
        }

        // 匹配成功
        deploymentOrders[_deploymentOrderID].isMatch = true;

        // 创建新的待供应商填充的部署信息
        deploymentInfos[_deploymentOrderID] = DeploymentInfo({
            unitPrice: _unitPrice,
            isProviderPublish: false,
            provider: _provider,
            funcPath: "",
            deployPath: "",
            accessKey: ""
        });
        
        // 修改状态
        deploymentOrders[_deploymentOrderID].state = OrderStates.Confirming;
    }


    // 供应商 API
    // 供应商确认部署订单，同时提交部署信息
    function publishDeploymentInfo(
        uint _deploymentOrderID,
        string memory _funcPath,
        string memory _deployPath,
        string memory _accessKey) 
        public
        validDeploymentOrderID(_deploymentOrderID)
        atOrderState(_deploymentOrderID, OrderStates.Confirming)
    {
        // TODO 检查超时
        // 如果超时 退还预付款 结束订单流程

        DeploymentInfo storage _info = deploymentInfos[_deploymentOrderID];

        require(
            msg.sender == _info.provider,
            "Matket: only provider can publish deployment information"
        );

        // 参数检查
        require(bytes(_funcPath).length   > 0 && bytes(_funcPath).length   <= 256, "");
        require(bytes(_deployPath).length > 0 && bytes(_deployPath).length <= 256, "");
        require(bytes(_accessKey).length  > 0 && bytes(_accessKey).length  <= 256, "");


        _info.funcPath = _funcPath;
        _info.deployPath = _deployPath;
        _info.accessKey = _accessKey;

        // 确认
        _info.isProviderPublish = true;

        emit NewDeploymentInfoEvent(
            _deploymentOrderID,
            _info.provider,
            _funcPath,
            _deployPath,
            _accessKey);
    }

    // 租户 API
    // 在供应商确认部署订单之后，租户进行确认，生成租约
    function confirmDeploymentInfo(uint _deploymentOrderID, bytes32 _fulfillKey) 
        public
        validDeploymentOrderID(_deploymentOrderID)
        atOrderState(_deploymentOrderID, OrderStates.Confirming)
    {
        // TODO 超时退还预付款

        require(
            msg.sender == deploymentOrders[_deploymentOrderID].customer,
            "Market: only customer can confirm");

        // 要求供应商确认在前
        require(deploymentInfos[_deploymentOrderID].isProviderPublish == true,
            "Market: provider has not confirmed yet");

        // 新建租约
        DeploymentOrder storage _order = deploymentOrders[_deploymentOrderID];
        DeploymentInfo  storage _info  = deploymentInfos[_deploymentOrderID];

        uint _customerLockFee = calculateLockFee(_order.highestUnitPrice, _order.faasDuration);  // 租户的预付款
        uint _customerPayFee = calculateServiceFee(_info.unitPrice, _order.faasDuration );       // 租户的实际应付                 
        uint _customerWithdrawFee = _customerLockFee - _customerPayFee;                          // 租户应取回的预付款
        uint _platformChargeFee = 0;                                                             // 平台手续费
        uint _providerServiceFee = _customerPayFee - _platformChargeFee;                         // 服务成功时，供应商可得服务报酬
        uint _customerCompensationFee = _providerServiceFee;                                     // 服务失败时，租户可得的补偿款

        // 新建租约
        leases[_deploymentOrderID] = Lease({
            providerServiceFee: _providerServiceFee,
            customerWithdrawFee: _customerWithdrawFee,
            customerCompensationFee: _customerCompensationFee,
            platformChargeFee: _platformChargeFee,
            fulfillKey: _fulfillKey,
            faasFulfillStartTime: 0,        // 待填写
            isViolatedSLA: false,      // 待填写
            curBlockNum: block.number
        });

        emit NewLeaseEvent(
            _deploymentOrderID, 
            _order.customer, 
            _info.provider);

        // 修改状态
        deploymentOrders[_deploymentOrderID].state = OrderStates.Deploying;
    }

    // 租户 API
    // 租户部署完成，开始履行部署订单
    function fulfillDeploymentOrder(uint _deploymentOrderID, bytes32 _fulfillSecretKey)
        public
        validDeploymentOrderID(_deploymentOrderID)
        atOrderState(_deploymentOrderID, OrderStates.Deploying)
    {
        // TODO 超时退还预付款

        DeploymentOrder storage _order = deploymentOrders[_deploymentOrderID];
        DeploymentInfo  storage _info  = deploymentInfos[_deploymentOrderID];
        Lease storage _lease = leases[_deploymentOrderID];

        // 检验
        // see: https://docs.soliditylang.org/en/develop/abi-spec.html#non-standard-packed-mode
        require(
            _lease.fulfillKey == sha256(abi.encodePacked(_fulfillSecretKey)),
            "Market: wrong fulfill secret key");

        // 记录时间
        _lease.faasFulfillStartTime = block.timestamp;

        // 生成 SLA 监督，由证人执行对供应商的监视
        wpContract.newSLA(
            _lease.curBlockNum,
            _info.provider, 
            _order.customer,
            _deploymentOrderID,
            _info.funcPath,
            _order.faasDuration - 1 hours);
        
        emit NewSLAEvent(_deploymentOrderID);


        // 修改订单状态
        deploymentOrders[_deploymentOrderID].state= OrderStates.Fulfilling;
    }

    // 结束部署订单，结算并完成转账
    function finishDeploymemtOrder(uint _deploymentOrderID) 
        public
        validDeploymentOrderID(_deploymentOrderID)
        atOrderState(_deploymentOrderID, OrderStates.Fulfilling)
    {
        DeploymentOrder storage _order = deploymentOrders[_deploymentOrderID];
        DeploymentInfo  storage _info  = deploymentInfos[_deploymentOrderID];
        Lease storage _lease = leases[_deploymentOrderID];

        require(
            block.timestamp > _lease.faasFulfillStartTime + _order.faasDuration,
            "Matket: FaaS is still fulfilling"
        );

        // 查看是否违反 SLA
        _lease.isViolatedSLA = wpContract.isViolatedSLA(_deploymentOrderID);

        // 支付平台手续费
        require(tokenContract.transfer(address(this), _lease.platformChargeFee), "");

        // 退回租户预付款多出的部分
        require(tokenContract.transfer(_order.customer, _lease.customerWithdrawFee), "");  
        
        if (_lease.isViolatedSLA == true) {
            // 如违反 SLA，退回租户补偿款
            require(tokenContract.transfer(_order.customer, _lease.customerCompensationFee), "");
        } else {
            //  如不违约 SLA，支付供应商报酬
            require(tokenContract.transfer(_info.provider, _lease.providerServiceFee), "");
        }

        // 转账完成的订单状态为 Finished，整个订单流程结束
        deploymentOrders[_deploymentOrderID].state = OrderStates.Finished;

        emit FinishDeploymentOrderEvent(
            _deploymentOrderID,
            _lease.isViolatedSLA);
    }


    // ------------------------------------------------------------------------------------
    // 辅助函数
    // ------------------------------------------------------------------------------------

    // 计算部署订单的服务费
    function calculateServiceFee(uint _unitPrice, uint _faasDuration) public pure returns (uint) {
        // _faasDuration 以秒为单位
        // FaaS 费用以小时计价，不足整个小时部分按一小时计算
        return (_unitPrice * ( (_faasDuration + (1 hours) - 1) / (1 hours)) );
    }

    // 计算部署订单的预付款：锁定双倍费用
    function calculateLockFee(uint _highestUnitPrice, uint _faasDuration) public pure returns (uint) {
        return ( 2 * calculateServiceFee(_highestUnitPrice, _faasDuration));
    }

    // 退还预付款
    function _freeLockFee(uint _deploymentOrderID) 
        private 
        validDeploymentOrderID(_deploymentOrderID)
    {
        require(
            isAtOrderStates(_deploymentOrderID, OrderStates.Bidding) || 
            isAtOrderStates(_deploymentOrderID, OrderStates.Confirming) || 
            isAtOrderStates(_deploymentOrderID, OrderStates.Deploying),
            "Market: not a state that can free locked fee"
        );

        DeploymentOrder storage _order = deploymentOrders[_deploymentOrderID];
        require(
            tokenContract.transfer(_order.customer, calculateLockFee(_order.highestUnitPrice,  _order.faasDuration)),
            "Market: failed to refund locked fee back to customer"
        );
    }

    // ------------------------------------------------------------------------------------
    // 读合约
    // ------------------------------------------------------------------------------------

    // 查询部署订单
    function getDeploymentOrder(uint _deploymentOrderID) 
        public
        view
        validDeploymentOrderID(_deploymentOrderID)
        returns (address, uint, uint, uint)
    {
        // 对于引用类型，状态变量向 storage 局部变量赋值时仅仅传递一个引用
        DeploymentOrder storage _order = deploymentOrders[_deploymentOrderID];
        
        return (
            _order.customer,
            _order.faasLevelID,
            _order.faasDuration,
            _order.highestUnitPrice
        );
    }

    function queryMatch(uint _deploymentOrderID)
        public  
        view
        validDeploymentOrderID(_deploymentOrderID)
        returns (bool)
    {
        return deploymentOrders[_deploymentOrderID].isMatch;
    }

    function queryLease(uint _deploymentOrderID)
        public
        view
        validDeploymentOrderID(_deploymentOrderID)
        returns (address, address)
    {
        DeploymentOrder storage _order = deploymentOrders[_deploymentOrderID];
        DeploymentInfo  storage _info  = deploymentInfos[_deploymentOrderID];

        return (_order.customer, _info.provider);
    }

}
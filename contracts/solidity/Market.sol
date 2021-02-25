// SPDX-License-Identifier: MIT

pragma solidity^0.6.0;

import "./FaaSToken.sol";
import "./SimpleAuction.sol";
import "./WitnessPool.sol";

contract Market {
    
    // 部署订单
    struct DeploymentOrder {
        address customer;        // 租户地址
        uint faaSLevelID;        // 租用的 FaaS 规格
        uint faaSDuration;       // 租用的时间
        uint highestUnitPrice;   // 租户接受的最高单价
    }
    
    // 租约
    struct Lease {
        address customer;  // 租户地址
        address provider;  // 供应商地址
        
        uint faaSLevelID;  // 租用的 FaaS 规格
        uint faaSDuration; // 租用的时间
        uint unitPrice;    // FaaS 规格对应的单价
        
        bool isViolatedSLA;            // 是否违反 SLA 合约
        uint providerServiceFee;       // 服务成功时，供应商的服务费用
        uint customerLockFee;          // 租户的预付款
        uint customerWithdrawFee;      // 租户取回多出的预付款
        uint customerCompensationFee;  // 服务失败时，租户的补偿款
        uint witnessFee;               // 证人的费用
        uint maintainerFee;            // 区块链维护者的费用
    }

    // ------------------------------------------------------------------------------------

    FaaSToken tokenContract;
    WitnessPool wpContract;

    // 供应商押金
    uint public providerDeposit;
    // 供应商押金表
    mapping(address => uint) private providerDeposits;
    // 供应商信誉初始值，是新注册的供应商的初始信誉值
    uint public providerReputationInit;  
    // 供应商信誉合格值，低于合格值为不合格
    uint public providerReputationQualified;  
    // 供应商信誉值表
    // 信誉值为 0 表示未注册
    mapping (address => uint) private providerReputations;

    // 部署订单数量
    uint public numDeploymentOrders;
    // 部署订单表：deploymentOrderID => DeploymentOrder
    mapping(uint => DeploymentOrder) public deploymentOrders;
    // 部署订单的拍卖地址表: deploymentOrderID => SimpleAuction
    mapping(uint => SimpleAuction) public auctions;
    // 匹配成功的部署订单的租约表: deploymentOrderID(leaseID) => Lease
    mapping(uint => Lease) public leases;
    // 部署订单的状态表：deploymentOrderID => orderState
    mapping(uint => OrderStates) private deploymentOrderStates;

    // ------------------------------------------------------------------------------------

    // 部署订单的有限状态机
    enum OrderStates {
        Bidding,    // 竞价中
        Fulfilling, // 履行中 
        Settling,   // 结算中
        Finished    // 已完成
    }

    modifier atOrderState(uint _deploymentOrderID, OrderStates _state) {
        require(
            deploymentOrderStates[_deploymentOrderID] == _state,
            "Matket: function cannot be called at this order state"
        );
        _;
    }
    
    // ------------------------------------------------------------------------------------


    // 新建部署订单事件
    event newDeploymentOrderEvent(uint _deploymentOrderID, address _auctionAddress);
    // 部署订单竞价结束事件
    event auctionEndEvent(uint _deploymentOrderID, bool _success, address _provider, uint _unitPrice);
    // 新租约事件
    event newLeaseEvent(address _customer, address _provider, uint _deploymentOrderID);

    // ------------------------------------------------------------------------------------

    constructor(address tokenContractAddress) public
    {        
        // 合约地址初始化
        tokenContract = FaaSToken(tokenContractAddress);

        // 供应商参数初始化
        providerDeposit = 100; // 100 token
        providerReputationInit = 5;
        providerReputationQualified = 5;

        // 部署订单参数初始化
        // 部署订单号计数， 有效的部署订单号从 1 开始，0 为无效的部署订单号
        numDeploymentOrders = 1;
    }

    // ------------------------------------------------------------------------------------

    // 查询供应商押金
    function getProviderDeposit() public view returns (uint) {
        return providerDeposit;
    } 

    // 查询供应商信誉合格值
    function getProviderReputationQualified() public view returns (uint) {
        return providerReputationQualified;
    }

    // 查询指定供应商的信誉值
    function getProviderReputation(address provider) public view returns (uint) {
        return providerReputations[provider];
    }

    // 查询是否已注册为供应商
    function isProviderRegistered(address provider) public view returns (bool) {
        // 信誉值为 0 表示供应商未注册, 大于 0 表示供应商已注册
        return (providerReputations[provider] > 0);
    }

    // 查询供应商是否合格
    function isProviderQualified(address provider) public view returns (bool) {
        // 低于合格值为不合格
        // 即，信誉值在区间 (0, providerReputationQualified) 的供应商为不合格的
        // 不合格的供应商既不能注销供应商资格取回押金，也不能参与竞价
        return (providerReputations[msg.sender] >= providerReputationQualified);
    }

    // 已注册
    modifier providerRegistered() {
        require(
            isProviderRegistered(msg.sender) == true,
            "Market: the address had not been registered"
        );
        _;
    }

    // 未注册
    modifier providerUnregistered() {
        require(
            isProviderRegistered(msg.sender) == false,
            "Market: the address had been registered"
        );
        _;
    }

    // 合格
    modifier providerQualified() {
        require(
            isProviderQualified(msg.sender),
            "Market: the provider is not qualified"
        );
        _;
    }

    // 注册供应商资格，支付押金
    function providerLogin() 
        public 
        providerUnregistered 
    {
        // 支付押金
        require(
            tokenContract.transferFrom(msg.sender, address(this), providerDeposit) == true,
            "Market: failed to pay a register deposit"
        );

        providerDeposits[msg.sender] = providerDeposit;
        providerReputations[msg.sender] = providerReputationInit;
    }

    // 注销供应商资格，取回押金
    function providerLogout() 
        public 
        providerRegistered
        providerQualified
    {
        // 信誉值清零
        providerReputations[msg.sender] = 0;

        // 检查：记录押金
        uint _depoist = providerDeposits[msg.sender];
        // 生效：押金清零
        providerDeposits[msg.sender] = 0;
        // 交互：退还押金
        require(
            tokenContract.transfer(msg.sender, _depoist) == true,
            "Market: failed to refund the provider deposit"
        );
    }

    // ------------------------------------------------------------------------------------

    // 计算订单预付款
    function getDeploymentOrderLockFee(uint _highestUnitPrice, uint _faaSDuration) 
        public 
        pure 
        returns (uint) 
    {
        // 安全性 TODO：防止乘法结果溢出
        return (_highestUnitPrice * _faaSDuration);
    }

    // 查询部署订单
    function getDeploymentOrder(uint _deploymentOrderID) 
        public
        view
        returns (address, uint, uint, uint)
    {
        DeploymentOrder memory order = deploymentOrders[_deploymentOrderID];
        
        return (
            order.customer,
            order.faaSLevelID,
            order.faaSDuration,
            order.highestUnitPrice
        );
    }

    // 新建部署订单
    function newDeploymentOrder(
        uint _faaSLevelID,
        uint _highestUnitPrice,
        uint _faaSDuration,
        uint _biddingDuration) 
        public 
        returns (uint, address) 
    {
        // 锁定预付款
        uint lockFee = getDeploymentOrderLockFee(_highestUnitPrice, _faaSDuration);
        require(
            tokenContract.transferFrom(msg.sender, address(this), lockFee) == true,
            "lock fee before creating deployment order"
        );

        // 构造新订单
        uint _deploymentOrderID = numDeploymentOrders++;
        deploymentOrders[_deploymentOrderID] = DeploymentOrder({
            customer: msg.sender,
            faaSLevelID: _faaSDuration,
            highestUnitPrice: _highestUnitPrice,
            faaSDuration: _faaSDuration
        });

        // 为新订单创建竞价合约
        SimpleAuction _auction = new SimpleAuction(
            address(this),
            _deploymentOrderID,
            _highestUnitPrice,
            _biddingDuration
        );
        auctions[_deploymentOrderID] = _auction;

        // 为新订单创建状态
        deploymentOrderStates[_deploymentOrderID] = OrderStates.Bidding;

        // 产生事件
        emit newDeploymentOrderEvent(_deploymentOrderID, address(_auction));

        return (_deploymentOrderID, address(_auction));
    }

    // 对部署订单竞价
    function bid(uint _deploymentOrderID, uint _unitPrice) 
        public
        providerRegistered
        providerQualified
        atOrderState(_deploymentOrderID, OrderStates.Bidding)
    {
        auctions[_deploymentOrderID].bid(msg.sender, _unitPrice);
    }

    // 检查部署订单的匹配结果
    // 返回：（是否匹配成功，供应商信息）
    function matchDeploymentOrder(uint _deploymentOrderID) 
        public
        atOrderState(_deploymentOrderID, OrderStates.Bidding)
        returns (bool, string memory) 
    {
        bool _success;
        address _provider;
        uint _unitPrice;

        (_success, _provider, _unitPrice) = auctions[_deploymentOrderID].auctionEnd();
        emit auctionEndEvent(_deploymentOrderID, _success, _provider, _unitPrice);
        
        // 匹配失败
        if (_success == false) {
            // 修改匹配失败的订单状态为 Finished
            deploymentOrderStates[_deploymentOrderID] = OrderStates.Finished;
            return (false, "");
        }

        // 匹配成功, 创建新租约
        Lease memory _lease = newLease(_deploymentOrderID, _provider, _unitPrice);
        leases[_deploymentOrderID] = _lease;
        emit newLeaseEvent(_lease.customer, _lease.provider, _deploymentOrderID);

        // 修改匹配成功订单状态为 Fulfilling
        deploymentOrderStates[_deploymentOrderID] = OrderStates.Fulfilling;

        return (true, "");
    }

    // 报告部署订单的履行情况
    function reportDeploymemtOrder(uint _delpoymentOrderID, bool _isViolatedSLA) 
        public
        atOrderState(_delpoymentOrderID, OrderStates.Fulfilling)
    {
        require(
            msg.sender == address(wpContract),
            "Market: only WitnessPool Contract can report"
        );

        leases[_delpoymentOrderID].isViolatedSLA = _isViolatedSLA;

        // 报告完成后，订单可以进入结算状态
        deploymentOrderStates[_delpoymentOrderID] = OrderStates.Settling;
    }

    // 结算部署订单
    function settleDeploymemtOrder(uint _deploymentOrderID) 
        public
        atOrderState(_deploymentOrderID, OrderStates.Settling)
    {
        Lease memory lease = leases[_deploymentOrderID];

        // 支付区块链维护者费用 TODO
        address _maintainerAddress = address(0);
        require(tokenContract.transfer(_maintainerAddress, lease.maintainerFee), "");

        // 支付证人费用
        require(tokenContract.transfer(address(wpContract), lease.witnessFee), "");
        
        // 退回租户预付款多出的部分
        require(tokenContract.transfer(lease.customer, lease.customerWithdrawFee), "");  
        
        if (lease.isViolatedSLA == true) {
            // 如违约，退回租户补偿费
            require(tokenContract.transfer(lease.customer, lease.customerWithdrawFee), "");
        } else {
            //  如不违约，支付供应商报酬
            require(tokenContract.transfer(lease.provider, lease.providerServiceFee), "");
        }

        // 结算完成的订单状态为 Finished
        deploymentOrderStates[_deploymentOrderID] = OrderStates.Finished;
    }


    // ------------------------------------------------------------------------------------

    // 新建租约
    // 返回租约ID
    function newLease(
        uint _deploymentOrderID,
        address _provider,
        uint _unitPrice) 
        internal
        returns (Lease memory)
    {
        uint _leaseID = _deploymentOrderID;  // leaseID 就是 deploymentOrderID
        DeploymentOrder memory order = deploymentOrders[_deploymentOrderID];

        // 租户的服务款
        uint _customerLockFee =  getDeploymentOrderLockFee(order.highestUnitPrice, order.faaSDuration);
        // 租户的实际支付
        uint _customerPayFee = order.faaSDuration * _unitPrice;
        // 租户多出的预付款
        uint _customerWithdrawFee = _customerLockFee - _customerPayFee;
        // 证人费用 TODO
        uint _witnessFee = 0;
        // 区块链维护者费用 TODO
        uint _maintainerFee = 0;
        // 服务成功时，供应商可得服务报酬
        uint _providerServiceFee = _customerPayFee - (_witnessFee + _maintainerFee);
        // 服务失败时，租户可得的补偿款
        uint _customerCompensationFee = _providerServiceFee;

        Lease memory _lease = Lease({
            customer: order.customer,
            provider: _provider,
            faaSLevelID: order.faaSLevelID,
            faaSDuration: order.faaSDuration,
            unitPrice: _unitPrice,
            isViolatedSLA: false,
            providerServiceFee: _providerServiceFee,
            customerLockFee: _customerLockFee,
            customerWithdrawFee: _customerWithdrawFee,
            customerCompensationFee: _customerCompensationFee,
            witnessFee: _witnessFee,
            maintainerFee: _maintainerFee
        });

        // TODO 产生 SLA 合约等等
        // wp.....
        
        return _lease;
    }

}
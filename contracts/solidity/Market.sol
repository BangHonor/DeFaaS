// SPDX-License-Identifier: MIT

pragma solidity^0.6.0;

import "./FaaSTokenPay.sol";
import "./ProviderPool.sol";
import "./SimpleAuction.sol";
import "./WitnessPool.sol";

contract Market is FaaSTokenPay, ProviderPool {
    
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

    WitnessPool wpContract;

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
    event NewDeploymentOrderEvent(uint _deploymentOrderID, address _auctionAddress);
    // 部署订单竞价结束事件
    event AuctionEndEvent(uint _deploymentOrderID, bool _success, address _provider, uint _unitPrice);
    // 新租约事件
    event NewLeaseEvent(address _customer, address _provider, uint _deploymentOrderID);

    // ------------------------------------------------------------------------------------

    constructor(address _tokenContractAddress) 
        FaaSTokenPay(_tokenContractAddress)
        public
    {        
        // 部署订单参数初始化
        // 部署订单号计数， 有效的部署订单号从 1 开始，0 为无效的部署订单号
        numDeploymentOrders = 1;
    }

  
    // ------------------------------------------------------------------------------------

    // 注册供应商资格，支付押金
    function providerLogin() 
        public
        providerUnregistered 
    {
        // 注册
        _providerLogin();

        // 支付押金
        require(
            tokenContract.transferFrom(msg.sender, address(this), stdProviderDeposit) == true,
            "Market: failed to pay a register deposit"
        );    
    }

    // 注销供应商资格，取回押金
    function providerLogout() 
        public
        providerRegistered
        providerQualified 
    {
        // 检查：记录押金
        uint _depoist = getProviderDeposit(msg.sender);
        // 生效：注销
        _providerLogout();
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
        DeploymentOrder memory _order = deploymentOrders[_deploymentOrderID];
        
        return (
            _order.customer,
            _order.faaSLevelID,
            _order.faaSDuration,
            _order.highestUnitPrice
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
        emit NewDeploymentOrderEvent(_deploymentOrderID, address(_auction));

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
        emit AuctionEndEvent(_deploymentOrderID, _success, _provider, _unitPrice);
        
        // 匹配失败
        if (_success == false) {
            // 修改匹配失败的订单状态为 Finished
            deploymentOrderStates[_deploymentOrderID] = OrderStates.Finished;
            return (false, "");
        }

        // 匹配成功, 创建新租约
        Lease memory _lease = newLease(_deploymentOrderID, _provider, _unitPrice);
        leases[_deploymentOrderID] = _lease;
        emit NewLeaseEvent(_lease.customer, _lease.provider, _deploymentOrderID);

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
        Lease memory _lease = leases[_deploymentOrderID];

        // 支付区块链维护者费用 TODO
        address _maintainerAddress = address(0);
        require(tokenContract.transfer(_maintainerAddress, _lease.maintainerFee), "");

        // 支付证人费用
        require(tokenContract.transfer(address(wpContract), _lease.witnessFee), "");
        
        // 退回租户预付款多出的部分
        require(tokenContract.transfer(_lease.customer, _lease.customerWithdrawFee), "");  
        
        if (_lease.isViolatedSLA == true) {
            // 如违约，退回租户补偿费
            require(tokenContract.transfer(_lease.customer, _lease.customerWithdrawFee), "");
        } else {
            //  如不违约，支付供应商报酬
            require(tokenContract.transfer(_lease.provider, _lease.providerServiceFee), "");
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
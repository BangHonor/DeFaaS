// SPDX-License-Identifier: MIT

pragma solidity^0.6.0;

import "./FaaSToken.sol";
import "./SimpleAuction.sol";

contract Market {
    
    // 部署订单
    struct DelpoymentOrder {
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
        
        uint providerServiceFee;       // 服务成功，供应商的费用
        uint customerCompensationFee;  // 服务失败，租户取回的费用
        uint witnessFee;               // 证人的费用
        uint maintainerFee;            // 区块链维护者的费用
    }

    // ------------------------------------------------------------------------------------

    FaaSToken tokenContract;
    // WitnessPool WPContract;

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
    uint public numDelpoymentOrders;
    // 部署订单表
    mapping(uint => DelpoymentOrder) public delpoymentOrders;
    // 部署订单的拍卖地址表
    mapping(uint => SimpleAuction) public auctions;

    // 租约数量
    uint public numLeases;
    // 租约表
    mapping(uint => Lease) public leases;

    // ------------------------------------------------------------------------------------

    event newDeploymentOrderEvent(uint _delpoymentOrderID, address _auctionAddress);

    // ------------------------------------------------------------------------------------
    constructor(address tokenContractAddress) public
    {        
        tokenContract = FaaSToken(tokenContractAddress);
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
    
    // 注册供应商资格，支付押金
    function providerLogin() public {
        
        // 信誉值为 0 表示供应商未注册, 大于 0 表示供应商已注册
        require(
            providerReputations[msg.sender] == 0,
            "Market: the address had been registered"
        );

        // 支付押金
        require(
            tokenContract.transferFrom(msg.sender, address(this), providerDeposit) == true,
            "Market: failed to pay a register deposit"
        );

        providerDeposits[msg.sender] = providerDeposit;
        providerReputations[msg.sender] = providerReputationInit;
    }

    // 注销供应商资格，取回押金
    function providerLogout() public {

        // 未注册成为供应商
        require(
            providerReputations[msg.sender] > 0,
            "Market: the address had not been registered"
        );

        // 不合格的供应商不能注销资格并取回押金
        // 即，信誉值在区间 (0, providerReputationQualified) 的供应商被锁定在黑名单中
        require(
            providerReputations[msg.sender] > providerReputationQualified,
            "Market: the provider is not qualified"
        );

        // 信誉值清 0
        providerReputations[msg.sender] = 0;

        // 重入检查
        if (providerDeposits[msg.sender] > 0) {

            // 退还押金
            require(
                tokenContract.transfer(msg.sender, providerDeposits[msg.sender]) == true,
                "Market: failed to refund the provider deposit"
            );

            // 押金清 0
            providerDeposits[msg.sender] = 0;
        }
    }

    // ------------------------------------------------------------------------------------

    // 新建部署订单
    function newDeploymentOrder(
        uint _faaSLevelID,
        uint _highestUnitPrice,
        uint _faaSDuration,
        uint _biddingDuration) 
        public 
        returns (uint, address) 
    {
        uint lockFee = _highestUnitPrice * _faaSDuration;
        require(
            tokenContract.transferFrom(msg.sender, address(this), lockFee) == true,
            "lock fee before creating deployment order"
        );

        uint _delpoymentOrderID = numDelpoymentOrders++;

        delpoymentOrders[_delpoymentOrderID] = DelpoymentOrder({
            customer: msg.sender,
            faaSLevelID: _faaSDuration,
            highestUnitPrice: _highestUnitPrice,
            faaSDuration: _faaSDuration
        });

        SimpleAuction _auction = new SimpleAuction(
            _delpoymentOrderID,
            _highestUnitPrice,
            _biddingDuration
        );
        auctions[_delpoymentOrderID] = _auction;

        emit newDeploymentOrderEvent(_delpoymentOrderID, address(_auction));

        return (_delpoymentOrderID, address(_auction));
    }

    // 查询部署订单
    function getDeploymentOrder(uint _delpoymentOrderID) 
        public
        view
        returns (address, uint, uint, uint)
    {
        DelpoymentOrder memory order = delpoymentOrders[_delpoymentOrderID];
        
        return (
            order.customer,
            order.faaSLevelID,
            order.faaSDuration,
            order.highestUnitPrice
        );
    }

    // ------------------------------------------------------------------------------------

    // 新建租约
    function newLease(
        uint _delpoymentOrderID,
        address _provider,
        uint _unitPrice)
        public
        returns (uint) 
    {
        require(
            msg.sender == address(auctions[_delpoymentOrderID]),
            "Only specific auction smart contract can new lease."
        );

        DelpoymentOrder memory order = delpoymentOrders[_delpoymentOrderID];

        uint _leaseID = numLeases++;
        uint _serviceFee = order.faaSDuration * _unitPrice;

        // TODO
        uint _providerServiceFee = _serviceFee;
        uint _customerCompensationFee = 0;
        uint _witnessFee = 0;
        uint _maintainerFee = 0;

        leases[_leaseID] = Lease({
            customer: order.customer,
            provider: _provider,
            faaSLevelID: order.faaSLevelID,
            faaSDuration: order.faaSDuration,
            unitPrice: _unitPrice,
            providerServiceFee: _providerServiceFee,
            customerCompensationFee: _customerCompensationFee,
            witnessFee: _witnessFee,
            maintainerFee: _maintainerFee
        });

        return (_leaseID);
    }
}
pragma solidity^0.6.0;


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

    // 供应商信誉表
    // 供应商的信誉值低于某个阈值，不能参与竞价
    mapping (address => uint) providerReputation;

    // 部署订单表
    uint numDelpoymentOrders;
    mapping(uint => DelpoymentOrder) delpoymentOrders;

    // 部署订单的拍卖地址表
    mapping(uint => address) auctions;

    // 租约表
    uint numLeases;
    mapping(uint => Lease) leases;
    

    event newDeploymentOrderEvent(uint _delpoymentOrderID, address _auctionAddress);

    constructor() public
    {        
        // TODO
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
        // TODO
        // uint _highestFee = _highestUnitPrice * _faaSDuration;
        // bool success = FaaSToken.transferFrom(msg.sender, address(this), _highestFee);
        // require(success == true, "transferFrom before publishing deployment order"); 

        uint _delpoymentOrderID = numDelpoymentOrders++;

        delpoymentOrders[_delpoymentOrderID] = DelpoymentOrder({
            customer: msg.sender,
            faaSLevelID: _faaSDuration,
            highestUnitPrice: _highestUnitPrice,
            faaSDuration: _faaSDuration,
            biddingEndTime: now + _biddingDuration,
            isBiddingEnd: false
        });

        address _auctionAddress = new SimpleAuction(
            // TODO
        );
        auctions[_delpoymentOrderID] = _auctionAddress;

        emit newDeploymentOrderEvent(_delpoymentOrderID. _auctionAddress);

        return (_delpoymentOrderID, _auctionAddress);
    }

    // 查询部署订单
    function getDeploymentOrder(uint _delpoymentOrderID) 
        public
        view
        returns (address, uint, uint, uint)
    {
        DelpoymentOrder order = delpoymentOrders[_delpoymentOrderID];
        
        return (
            order.customer,
            order.faaSLevelID,
            ordre.faaSDuration,
            order.highestUnitPrice
        );
    }

    // 新建租约
    function newLease(
        uint _delpoymentOrderID,
        address _provider,
        uint _unitPrice,)
        public
        returns (uint) 
    {
        require(
            msg.sender == auctions[_delpoymentOrderID],
            "Only specific auction smart contract can new lease."
        );

        DelpoymentOrder order = delpoymentOrders[_delpoymentOrderID];

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
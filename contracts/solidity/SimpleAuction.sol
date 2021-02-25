// SPDX-License-Identifier: MIT

pragma solidity^0.6.0;

contract SimpleAuction {
    
    address marketContractAddress;

    // 只有 Matket 合约才能调用
    // 因为 Matket 合约预先做了供应商资格检查
    modifier marketCall() {
        require(
            msg.sender == marketContractAddress,
            "SimpleAuction: only market conctract can call this funciton"
        );
        _;
    }

    // 两个状态：接受竞价 和 竞价结束
    enum States {
        AcceptingBids,
        StoppedBids
    }

    States public state;

    uint public delpoymentOrderID;
    uint public highestUnitPrice;
    
    uint public createTime;
    uint public biddingDuration;

    uint lowestUnitPrice;
    address provider;

    constructor(
        address _marketContractAddress,
        uint _delpoymentOrderID, 
        uint _highestUnitPrice,
        uint _biddingDuration) 
        public 
    {
        marketContractAddress = _marketContractAddress;

        state = States.AcceptingBids;
        
        delpoymentOrderID = _delpoymentOrderID;
        highestUnitPrice = _highestUnitPrice;

        createTime = now;
        biddingDuration = _biddingDuration;

        // init value
        lowestUnitPrice = _highestUnitPrice;
        provider = address(0);
    }

    // 状态检查：函数只能在对应状态调用
    modifier atState(States _state) {
        require(
            state == _state,
            "SimpleAuction: function cannot be called at this state"
        );
        _;
    }

    // 有限状态机的状态转移
    function nextState() internal {
        if (state == States.AcceptingBids) {
            state = States.StoppedBids;
        } 
        else if (state == States.StoppedBids) {
            state = States.StoppedBids;
        }
    }

    // 根据时间的状态转换
    modifier timedTransitions() {
        if (state == States.AcceptingBids && now > createTime + biddingDuration) {
                nextState();
        }
        _;
    }

    // 这个修饰器在函数执行结束之后，使合约进入下一个状态。
    modifier transitionNext()
    {
        _;
        nextState();
    }

    // 竞争最低出价
    function bid(address _provider, uint _unitPrice) 
        public
        marketCall
        timedTransitions
        atState(States.AcceptingBids) 
    {
        require (
            _unitPrice <= highestUnitPrice,
            "SimpleAuction: the unit-price is higher than highest the customer accepted"
        );

        require(
            _unitPrice < lowestUnitPrice,
            "SimpleAuction: there already is a lower uint-price."
        );

        lowestUnitPrice = _unitPrice;
        provider = _provider;
    }

    // 结束竞价:调用后合约状态转移到 Finished
    // 返回：（是否匹配成功，供应商，中标出价）
    function auctionEnd() 
        public
        timedTransitions
        atState(States.StoppedBids)
        returns (bool, address, uint)
    {
        bool _success = ( provider != address(0) );
        return (_success, provider, lowestUnitPrice);
    }
}
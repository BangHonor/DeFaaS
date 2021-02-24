// SPDX-License-Identifier: MIT

pragma solidity^0.6.0;

contract SimpleAuction {

    // 两个状态：接受竞价 和 竞价结束
    enum States {
        AcceptingBids,
        StoppedBids,
        Finished
    }

    States public state;

    uint public delpoymentOrderID;
    uint public highestUnitPrice;
    
    uint public createTime;
    uint public biddingDuration;

    uint lowestUnitPrice;
    address provider;

    bool success;

    constructor(
        uint _delpoymentOrderID, 
        uint _highestUnitPrice,
        uint _biddingDuration) 
        public 
    {
        
        state = States.AcceptingBids;
        
        delpoymentOrderID = _delpoymentOrderID;
        highestUnitPrice = _highestUnitPrice;

        createTime = now;
        biddingDuration = _biddingDuration;

        // init value
        lowestUnitPrice = _highestUnitPrice;
        provider = address(0);

        success = false;
    }

    // 状态检查：函数只能在对应状态调用
    modifier atState(States _state) {

        require(
            state == _state,
            "Function cannot be called at this state."
        );
        _;
    }

    // 有限状态机的状态转移
    function nextState() internal {

        if (state == States.AcceptingBids) {
            state = States.StoppedBids;
        }
        if (state == States.StoppedBids) {
            state = States.Finished;
        }
        if (state == States.Finished) {
            state = States.Finished;
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

    // 新的最低出价事件
    event lowestUnitPriceEvent(uint _delpoymentOrderID, uint _lowestUnitPrice, address _provider);
    // 竞价结束事件
    event auctionEndEvent(bool success, uint _delpoymentOrderID, address _provider);

    // 竞争最低出价
    function bid(uint _unitPrice) 
        public
        timedTransitions
        atState(States.AcceptingBids) 
    {

        require (
            _unitPrice <= highestUnitPrice,
            "The unit-price is higher than highest the customer accepted."
        );

        require(
            _unitPrice < lowestUnitPrice,
            "There already is a lower uint-price."
        );

        lowestUnitPrice = _unitPrice;
        provider = msg.sender;

        emit lowestUnitPriceEvent(delpoymentOrderID, lowestUnitPrice, provider);
    }

    // 结束竞价:调用后合约状态转移到 Finished
    function auctionEnd() 
        public
        timedTransitions
        atState(States.StoppedBids)
        transitionNext  
    {
        if (provider != address(0)) {
            success = true;
        }
        emit auctionEndEvent(success, delpoymentOrderID, provider);
    }


    function getResult() 
        public
        view
        atState(States.Finished)
        returns (bool, uint, uint, address)
    {
        return (success, delpoymentOrderID, lowestUnitPrice, provider);
    }
}
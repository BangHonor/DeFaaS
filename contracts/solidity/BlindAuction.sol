pragma solidity^0.6.0;


contract BlindAuction {
    

    // 盲拍合约的状态
    enum States {
        AcceptingBlindedBids,
        RevealingBids,
        StoppedReveal,
        Finished
    }

    States public state;

    uint createTime;
    uint biddingDuration;
    uint revealDuration;

    // 有限状态机的状态转移
    function nextState() internal {

        if (state = States.AcceptingBlindedBids) {
            state = States.RevealingBids;
        }
        if (state == States.RevealingBids) {
            state = States.StoppedReveal;
        }
        if (state == States.StoppedReveal) {
            state = States.Finished;
        }
        if (state == States.Finished) {
            state = States.Finished;
        }
    }


    // 状态检查：函数只能在对应状态调用
    modifier atState(States _state) {

        require(
            state == _state,
            "Function cannot be called at this state."
        );
        _;
    }


    // 根据时间的状态转换
    modifier timedTransitions() {

        if (state == States.AcceptingBlindedBids && now > createTime + biddingDuration) {
            nextState();
        }
        if (state = States.RevealingBids && now > createTime + biddingDuration + revealDuration) {
            nextState();
        }
        _;
    }

    // 这个修饰器在函数执行结束之后，使合约进入下一个状态。
    modifier transitionNext()
    {
        _;
        nextStage();
    }
    
    modifier timedTransitions() {
        if (stage == Stages.AcceptingBlindedBids &&
                    now >= creationTime + 10 days)
            nextStage();
        if (stage == Stages.RevealBids &&
                now >= creationTime + 12 days)
            nextStage();
        // 由交易触发的其他阶段转换
        _;
    }

    constructor()
        public
    {
        state = States.AcceptingBlindedBids;
    }

    // 秘密出价
    function bid()
        public
        timedTransitions
        atState(States.AcceptingBlindedBids)
    {
        // TODO
    }

    // 披露秘密出价
    function reveal()
        public
        timedTransitions
        atState(States.RevealingBids)
    {
        // TODO
    }

    // 结束竞拍
    function auctionEnd()
        public
        timedTransitions
        atState(States.StoppedReveal)
        transitionNext
    {
        // TODO
    }

    // 获取竞拍结果
    function getResult() 
        public
        atState(States.Finished)
    {
        // TODO
    }
}
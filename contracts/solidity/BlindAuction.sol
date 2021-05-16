// SPDX-License-Identifier: MIT

// 用于取代简单竞价（SimpleAuction）合约的秘密竞价（BlindAuction）合约：未完成

pragma solidity>=0.6.0;

contract BlindAuction {
    

    // 盲拍合约的状态
    enum States {
        AcceptingBlindedBids,
        RevealingBids,
        StoppedReveal,
        Finished
    }

    States public state;

    uint creationTime;
    uint biddingDuration;
    uint revealDuration;

    // 有限状态机的状态转移
    function nextState() internal {

        if (state == States.AcceptingBlindedBids) {
            state = States.RevealingBids;
        }
        else if (state == States.RevealingBids) {
            state = States.StoppedReveal;
        }
        else if (state == States.StoppedReveal) {
            state = States.Finished;
        }
        else if (state == States.Finished) {
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

        if (state == States.AcceptingBlindedBids && now > creationTime + biddingDuration) {
            nextState();
        }
        else if (state == States.RevealingBids && now > creationTime + biddingDuration + revealDuration) {
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
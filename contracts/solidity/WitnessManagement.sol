// SPDX-License-Identifier: MIT

pragma solidity>=0.6.0;

import "./FaaSTokenPay.sol";

contract WitnessManagement is FaaSTokenPay {

    enum WStates { Offline, Online, Confirming, Working }

    function isAtWState(address _witness, WStates _state) public view returns (bool) {
        return (witnessPool[_witness].state == _state);
    }

    modifier atWState(address _witness, WStates _state) {
        require(
            isAtWState(_witness, _state) == true,
            "WitenssPool: function cannot be called at this state"
        );
        _;
    }

    // ------------------------------------------------------------------------------------------------
        
    struct Witness {
        WStates state;      // 当前状态
        bool registered;    // 是否注册
        uint index;         // 在 witnessAddrs 数组中的索引
        uint reputation;    // 信誉值，初始为 100，当为 0 时 Witness 被封锁
        uint depoist;       // 交付的押金
        uint reward;        // 证人服务所获费用
    }

    // ------------------------------------------------------------------------------------------------

    // 证人押金标准
    uint public stdWitnessDepoist = 0; 
    // 证人信誉初始值
    uint public witnessReputationInit = 100;
    // 证人信誉合格值
    uint public witnessReputationQualified = 100;

    // 证人表
    mapping(address => Witness) internal witnessPool;    
    address [] internal witnessAddrs;

    // 证人 Online 数量
    uint internal numOnlineWitness;

    
    // ------------------------------------------------------------------------------------------------

    constructor(address _tokenContractAddress) 
        FaaSTokenPay(_tokenContractAddress)
    {
        numOnlineWitness           = 0;
    }

    // ------------------------------------------------------------------------------------------------

    // 证人是否已注册
    function isWitnessRegistered(address _witness) public view returns (bool) {
        return witnessPool[_witness].registered;
    }
    
    // 检查已注册
    modifier witenssRegisterd() {
        require(
            isWitnessRegistered(msg.sender) == true,
            "WitnessPool: the address had not been registered"
        );
        _;
    }

    // 检查未注册
    modifier witenssUnRegistered() {
        require(
            isWitnessRegistered(msg.sender) == false,
            "WitnessPool: the address had been registered"
        );
        _;
    }

    // 证人是否合格
    function isWitnessQualified(address _witness) public view returns (bool) {
        return (witnessPool[_witness].reputation >= witnessReputationQualified);
    }

    // 证人合格
    modifier witnessQualified() {
        require(
            isWitnessQualified(msg.sender) == true,
            "WitnessPool: the witness is unqualified"
        );
        _;
    }

    // ------------------------------------------------------------------------------------------------

    // 证人 API
    // 注册为证人
    function witnessLogin() 
        public 
        witenssUnRegistered
    {
        // 交押金
        require(
            tokenContract.transferFrom(msg.sender, address(this), stdWitnessDepoist),
            "WitnessPool: failed to lock witness depoist"
        );

        witnessAddrs.push(msg.sender);
        witnessPool[msg.sender] = Witness({
            state: WStates.Offline,
            registered: true,
            index: witnessAddrs.length - 1,
            reputation:  witnessReputationInit,
            depoist: stdWitnessDepoist,
            reward: 0
        });
    }
    
    // 证人 API
    // 注销证人资格
    function witenessLogout() 
        public
        witenssRegisterd
        witnessQualified
    {
        // 检查
        uint _depoist = witnessPool[msg.sender].depoist;
        // 生效
        witnessPool[msg.sender].state = WStates.Offline;
        witnessPool[msg.sender].registered = false;
        witnessPool[msg.sender].depoist = 0;
        // 交互：退还押金
        require(
            tokenContract.transfer(msg.sender, _depoist),
            "WitnessPool: failed to refund the witness deposit"
        );
    }

    // 证人 API
    // 转换到可以被抽选的 Online 状态
    function wintessTurnOn()
        public
        witenssRegisterd
        witnessQualified
        atWState(msg.sender, WStates.Offline)
    {   
        witnessPool[msg.sender].state = WStates.Online;
        numOnlineWitness++;
    }
    
    // 证人 API
    // 转换到不被抽选的 Offline 状态
    function witnessTurnOff()
        public
        witenssRegisterd
        atWState(msg.sender, WStates.Online)
    {
        witnessPool[msg.sender].state = WStates.Offline;
        numOnlineWitness--;
    }


    // 证人 API
    // 查询奖励
    function witnessQueryReward() 
        public
        view
        witenssRegisterd
        returns (uint)
    {
        return witnessPool[msg.sender].reward;
    }

    // 证人 API
    // 取出奖励
    function witnessDrawReward() 
        public
        witenssUnRegistered
        witnessQualified
    {
        require(
            witnessPool[msg.sender].reward > 0, 
            "WitnessPool: the reward of witness is zero"
        );

        // 检查
        uint _reward = witnessPool[msg.sender].reward;
        // 生效
        witnessPool[msg.sender].reward = 0;
        // 交互
        require(
            tokenContract.transfer(msg.sender, _reward),
            "WitnessPool: failed to draw the witness reward"
        );
    }
}

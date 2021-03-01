// SPDX-License-Identifier: MIT

pragma solidity^0.6.0;

import "./FaaSTokenPay.sol";
import "./SLA.sol";
import "./Owned.sol";


contract WitnessPool is Owned, FaaSTokenPay {

    // enum WStates { Offline, Online, Candidate, Busy }
    enum WStates { Offline, Online, Working }

    // ------------------------------------------------------------------------------------------------
        
    struct Witness {
        WStates state;      // 当前状态
        bool registered;    // 是否注册
        uint index;         // 在 witnessAddrs 数组中的索引
        uint reputation;    // 信誉值，初始为 100，当为 0 时 Witness 被封锁
        uint depoist;       // 交付的押金
    }

    // SLA 合约信息
    struct SLAInfo{
        bool valid;
        uint curBlockNum;
        uint8 blkNeed;     // how many blocks needed for sortition
    }

    // ------------------------------------------------------------------------------------------------

    // 证人表
    mapping(address => Witness) private witnessPool;    
    address [] private witnessAddrs;

    // 证人 Online 数量
    uint private onlineCounter;
    // 证人押金标准
    uint public stdWitnessDepoist;
    // 证人信誉初始值
    uint public witnessReputationInit;

    // SLA 合约表
    mapping(address => SLAInfo) SLAContractPool;
    
    // ------------------------------------------------------------------------------------------------

    constructor(address _tokenContractAddress)
        public
        FaaSTokenPay(_tokenContractAddress)
    {
        onlineCounter         = 0;
        stdWitnessDepoist     = 100;  // 100 token
        witnessReputationInit = 100;
    }

    // ------------------------------------------------------------------------------------------------

    event SLAContractGenEvent(address indexed _who, uint _time, address _contractAddr);
    event WitnessSelectedEvent(address indexed _who, uint _index, address _forWhom);

    // ------------------------------------------------------------------------------------------------

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
        return (witnessPool[_witness].reputation > 0);
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

    // 是否是有效的 SLA 合约
    function isValidSLAContract(address _sla) public view returns (bool) {
        return SLAContractPool[_sla].valid;
    }

    // 有效的 SLA 合约
    modifier validSLAContract() {
        require(isValidSLAContract(msg.sender) == true, "WitnessPool: invalid SLA contract");
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
            depoist: stdWitnessDepoist
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
    function turnOn()
        public
        witenssRegisterd
        witnessQualified
        atWState(msg.sender, WStates.Offline)
    {   
        witnessPool[msg.sender].state = WStates.Online;
        onlineCounter++;
    }
    
    // 证人 API
    // 转换到不被抽选的 Offline 状态
    function turnOff()
        public
        witenssRegisterd
        atWState(msg.sender, WStates.Online)
    {
        witnessPool[msg.sender].state = WStates.Offline;
        onlineCounter--;
    }

    // // 证人 API
    // // 拒绝被选中为某个 SLA 合约的证人委员会成员
    // function reject()
    //     public
    //     witenssRegisterd
    //     atWState(msg.sender, WStates.Candidate)
    // {   
    //     require( 
    //         now < witnessPool[msg.sender].confirmDeadline,
    //         "WitnessPool:can not reject after the confirm deadline"
    //     );

    //     witnessPool[msg.sender].state = WStates.Online;
    //     onlineCounter++;
    // }
    
    // // 证人 API
    // // 撤销 Candidate 状态，在确认截止日期后，将其自身状态更改为 Online
    // // 由于没有及时确认，需要降低证人的声誉值
    // function reverse()
    //     public
    //     witenssRegisterd
    //     atWState(msg.sender, WStates.Candidate)
    // {
    //     require(
    //         now > witnessPool[msg.sender].confirmDeadline,
    //         "WitnessPool: can not reverse before the confirm deadline"
    //     );

    //     witnessPool[msg.sender].reputation -= 10;
        
    //     if (isWitnessQualified(msg.sender) == false) {
    //         witnessPool[msg.sender].state = WStates.Offline;   // 强制下线
    //     } else{
    //         witnessPool[msg.sender].state = WStates.Online;
    //         onlineCounter++;
    //     }
    // }
    
    // ------------------------------------------------------------------------------------------------

    // Matket API
    // 产生一个 SLA 合约
    function genSLAContract() 
        public 
        onlyOwner
        returns (address)
    {
        SLA4FaaS newSLAContract = new SLA4FaaS(this, msg.sender, address(0x0));
        address newSLAContractAddress = address(newSLAContract);
        SLAContractPool[newSLAContractAddress].valid = true; 
        
        emit SLAContractGenEvent(msg.sender, now, newSLAContractAddress);

        return newSLAContractAddress;
    }
    
    /**
     * Contract Interface::
     * This is for SLA contract to submit a committee sortition request.
     * _blkNeed: This is a number to specify how many blocks needed in the future for the committee sortition. 
     *            Its range should be 2~255. The recommended value is 12.  
     * */
    function request(uint8 _blkNeed)
        public 
        validSLAContract
        returns
        (bool success)
    {
        //record current block number
        SLAContractPool[msg.sender].curBlockNum = block.number;
        SLAContractPool[msg.sender].blkNeed = _blkNeed;
        return true;
    }
    
    /**
     * Contract Interface::
     * Request for a sortition of _N witnesses. The _provider and _customer must not be selected.
     * */
    function sortition(uint _N, address _provider, address _customer)
        public
        validSLAContract
        returns
        (bool success)
    {
        // make sure the request is invoked before this interface
        require(SLAContractPool[msg.sender].curBlockNum != 0);
        // there should be more than 10 times of _N online witnesses
        require(onlineCounter >= _N+2);   //this is debug mode
        // require(onlineCounter > 10*_N);
        
        // currently, the hash value can only be accessed within 255 depth. In this case, invoke 'request' again
        require( block.number < SLAContractPool[msg.sender].curBlockNum + 255);
        // there should be more than extra 2*blkNeed blocks generated  
        require( block.number > SLAContractPool[msg.sender].curBlockNum + 2*SLAContractPool[msg.sender].blkNeed );

        uint seed = 0;
        for(uint bi = 0 ; bi < SLAContractPool[msg.sender].blkNeed; bi++)
        {
            seed += (uint)( blockhash( SLAContractPool[msg.sender].curBlockNum + bi + 1 ) );
        }
            
        
        uint wcounter = 0;
        while(wcounter < _N){
            address sAddr = witnessAddrs[seed % witnessAddrs.length];
            
            if(witnessPool[sAddr].state == WStates.Online && 
                witnessPool[sAddr].reputation > 0 && 
                sAddr != _provider && 
                sAddr != _customer)
            {
                // witnessPool[sAddr].state = WStates.Candidate;
                witnessPool[sAddr].confirmDeadline = now + 5 minutes;   // 5 minutes for confirmation
                witnessPool[sAddr].SLAContract = msg.sender;
                emit WitnessSelectedEvent(sAddr, witnessPool[sAddr].index, msg.sender);
                onlineCounter--;
                wcounter++;
            }
            
            seed = (uint)(keccak256(abi.encodePacked(bytes32(seed))));
        }
        
        // make this interface cannot be invoked twice without 'request'
        SLAContractPool[msg.sender].curBlockNum = 0;
        return true;
    }
    
    /**
     * Contract Interface::
     * Candidate witness calls the SLA contract and confirm the sortition. 
     * */
    function confirm(address _candidate)
        public
        validSLAContract
        returns 
        (bool)
    {
        require(isWitnessRegistered(_candidate) == true, "");

        //have not reached the confirmation deadline
        require( now < witnessPool[_candidate].confirmDeadline );
        
        //only able to confirm in candidate state
        require(witnessPool[_candidate].state == WStates.Candidate);
        
        //only the SLA contract can select it.
        require(witnessPool[_candidate].SLAContract == msg.sender);
        
        witnessPool[_candidate].state = WStates.Busy;
        
        return true;
    }
    
    /**
     * Contract Interface::
     * SLA contract ends and witness calls this from the contract to release the Busy witness.
     * If the reputation smaller than 0, the witness will be turned off.
     * */
    function release(address _witness)
        public
        validSLAContract
    {
        require(isWitnessRegistered(_witness) == true, "");

        //only able to release in Busy state
        require(witnessPool[_witness].state == WStates.Busy);
        
        //only the SLA contract can operate on it.
        require(address( witnessPool[_witness].SLAContract ) == msg.sender);
        
        if(witnessPool[_witness].reputation <= 0){
            witnessPool[_witness].state = WStates.Offline;
        }else{
            witnessPool[_witness].state = WStates.Online;
            onlineCounter++;
        }
        
    }
    
    /**
     * Contract Interface::
     * Decrease the reputation value. 
     * */
    function reputationDecrease(address _witness, int8 _value)
        public
        validSLAContract
    {
        require(isWitnessRegistered(_witness) == true, "");

        //only able to release in Busy state
        require( _value > 0 );
        
        //only the SLA contract can operate on it.
        require(address( witnessPool[_witness].SLAContract ) == msg.sender);
        
        witnessPool[_witness].reputation -= _value;
        
    }
    
    

}



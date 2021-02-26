// SPDX-License-Identifier: MIT

pragma solidity^0.6.0;

import "./FaaSTokenPay.sol";
import "./SLA.sol";
import "./Owned.sol";


contract WitnessPool is Owned, FaaSTokenPay {
    
    // witness Online 数量
    uint public onlineCounter = 0;
    
    enum WStates { Offline, Online, Candidate, Busy }
    
    struct Witness {
        bool registered;    // 是否注册
        uint index;         // 在 witnessAddrs 数组中的索引
        
        WStates state;    // 当前状态
        
        address SLAContract;   // 当前服务的 SLA 合约（被选中到服务该 SLA 合约）
        uint confirmDeadline;  // Candidata 状态下的最晚确认时间
        int8 reputation;       // 信誉值，初始为 100，当为 0 时 Witness 被封锁
    }

    mapping(address => Witness) witnessPool;    
    address [] public witnessAddrs;    // the address pool of witnesses

    // ------------------------------------------------------------------------------------------------

    // SLA 合约的信息
    struct SortitionInfo{
        bool valid;
        uint curBlockNum;
        uint8 blkNeed;   // how many blocks needed for sortition
    }
    mapping(address => SortitionInfo) SLAContractPool;   // record the requester's initial block number. The sortition will be based on the hash value after this block.
    
    
    // ------------------------------------------------------------------------------------------------

    address private marketContractAddress;

    function setMarketContractAddress(address _marketContractAddress) public onlyOwner {
        marketContractAddress = _marketContractAddress;
    }

    constructor(address _tokenContractAddress)
        public
        FaaSTokenPay(_tokenContractAddress)
    {
        marketContractAddress = address(0);
    }



    // ------------------------------------------------------------------------------------------------

    // record the provider _who generates a SLA contract of address _contractAddr at time _time
    event SLAContractGenEvent(address indexed _who, uint _time, address _contractAddr);
    
    event WitnessSelectedEvent(address indexed _who, uint _index, address _forWhom);

    // ------------------------------------------------------------------------------------------------

    // 证人是否已注册
    function isWitnessRegistered(address _witness) internal view returns (bool) {
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

    // ------------------------------------------------------------------------------------------------

    // 证人是否处于指定状态
    function isAtWitnessOrder(address _witness, WStates _state) internal view returns (bool) {
        return (witnessPool[_witness].state == _state);
    }

    // 检查证人的状态
    modifier atWitnessState(WStates _state) {
        require(
            isAtWitnessOrder(msg.sender, _state) == true,
            "WitnessPool: function cannot be called at this state"
        );
        _;
    }
    
    // ------------------------------------------------------------------------------------------------


    // 是否是有效的 SLA 合约
    function isValidSLAContract(address _sla) internal view returns (bool) {
        return SLAContractPool[_sla].valid;
    }

    // 检查有效的 SLA 合约
    modifier validSLAContract() {
        require(isValidSLAContract(msg.sender) == true, "");
        _;
    }

    // ------------------------------------------------------------------------------------------------


    // 产生一个 SLA 合约，仅可由 Market 合约调用
    function genSLAContract() 
        public 
        returns
        (address)
    {
        // require(msg.sender == marketContractAddress, "msg sender is not Market Contract");

        CloudSLA newSLAContract = new CloudSLA(this, msg.sender, address(0x0));
        address newSLAContractAddress = address(newSLAContract);
        SLAContractPool[newSLAContractAddress].valid = true; 
        emit SLAContractGenEvent(msg.sender, now, newSLAContractAddress);
        return newSLAContractAddress;
    }
    
    /**
     * Normal User Interface::
     * Check whether a SLAContract address is valid
     * */
    function validateSLA(address _SLAContract) 
        public
        view
        returns
        (bool)
    {
        if(SLAContractPool[_SLAContract].valid)
            return true;
        else
            return false;
    }
    
    /**
     * Normal User Interface::
     * This is for the normal user to register as a witness into the pool
     * */
    function register() 
        public 
        witenssUnRegistered
    {
        // witnessPool[msg.sender].index = witnessAddrs.push(msg.sender) - 1;
        witnessAddrs.push(msg.sender);
        witnessPool[msg.sender].index = witnessAddrs.length - 1;

        witnessPool[msg.sender].state = WStates.Offline;
        witnessPool[msg.sender].reputation = 100; 
        witnessPool[msg.sender].registered = true;
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
        //make sure the request is invoked before this interface
        require(SLAContractPool[msg.sender].curBlockNum != 0);
        // there should be more than 10 times of _N online witnesses
        require(onlineCounter >= _N+2);   //this is debug mode
        //require(onlineCounter > 10*_N);
        
        //currently, the hash value can only be accessed within 255 depth. In this case, invoke 'request' again
        require( block.number < SLAContractPool[msg.sender].curBlockNum + 255);
        // there should be more than extra 2*blkNeed blocks generated  
        require( block.number > SLAContractPool[msg.sender].curBlockNum + 2*SLAContractPool[msg.sender].blkNeed );
        uint seed = 0;
        for(uint bi = 0 ; bi<SLAContractPool[msg.sender].blkNeed ; bi++)
        {
            // seed += (uint)(block.blockhash( SLAContractPool[msg.sender].curBlockNum + bi + 1 ));
            seed += (uint)(blockhash( SLAContractPool[msg.sender].curBlockNum + bi + 1 ));
        }
            
        
        uint wcounter = 0;
        while(wcounter < _N){
            address sAddr = witnessAddrs[seed % witnessAddrs.length];
            
            if(witnessPool[sAddr].state == WStates.Online && 
                witnessPool[sAddr].reputation > 0 && 
                sAddr != _provider && 
                sAddr != _customer)
            {
                witnessPool[sAddr].state = WStates.Candidate;
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
    
    
    
    /**
     * Witness Interface::
     * Reject the sortition for candidate. Because the SLA contract is not valid.
     * */
    function reject()
        public
        witenssRegisterd
    {
        //only reject in candidate state
        require(witnessPool[msg.sender].state == WStates.Candidate);
        
        //have not reached the rejection deadline
        require( now < witnessPool[msg.sender].confirmDeadline );
        
        witnessPool[msg.sender].state = WStates.Online;
        onlineCounter++;
    }
    
    /**
     * Witness Interface::
     * Reverse its own state to Online after the confirmation deadline. But need to reduece the reputation. 
     * */
    function reverse()
        public
        witenssRegisterd
    {
        //must exceed the confirmation deadline
        require( now > witnessPool[msg.sender].confirmDeadline );
        
        //able to turn only in candidate state
        require(witnessPool[msg.sender].state == WStates.Candidate);
        
        witnessPool[msg.sender].reputation -= 10;
        
        if(witnessPool[msg.sender].reputation <= 0){
            witnessPool[msg.sender].state = WStates.Offline;
        }else{
            witnessPool[msg.sender].state = WStates.Online;
            onlineCounter++;
        }
    }
    
    /**
     * Witness Interface::
     * Turn online to wait for sortition. The witness with reputation samller than 0 cannot be turned on.
     * */
    function turnOn()
        public
        witenssRegisterd
    {
        //must be in the state of offline
        require(witnessPool[msg.sender].state == WStates.Offline);
        
        //its reputation must be bigger than 0
        require( witnessPool[msg.sender].reputation > 0 );
        
        witnessPool[msg.sender].state = WStates.Online;
        onlineCounter++;
    }
    
    /**
     * Witness Interface::
     * Turn offline to avoid sortition.
     * */
    function turnOff()
        public
        witenssRegisterd
    {
        
        //must be in the state of online
        require(witnessPool[msg.sender].state == WStates.Online);
        
        witnessPool[msg.sender].state = WStates.Offline;
        onlineCounter--;
    }
    
    
    /**
     * Witness Interface::
     * For witness itself to check the state of itself and the reputation.
     * If it is selected, following two return values show its confirmation deadline and the address of the SLA contract, who sortited it. 
     * */
    function checkWStates(address _witness)
        public
        view
        returns
        (WitnessPool.WStates, int8, uint, address)
    {
        return (
            witnessPool[_witness].state, 
            witnessPool[_witness].reputation, 
            witnessPool[_witness].confirmDeadline, 
            address( witnessPool[_witness].SLAContract ));
    }
    
}



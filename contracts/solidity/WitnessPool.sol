// SPDX-License-Identifier: MIT

pragma solidity^0.6.0;

import "./SLA.sol";

contract WitnessPool {
    
    uint public onlineCounter = 0;
    
    enum WState { Offline, Online, Candidate, Busy }
    
    struct Witness {
        bool registered;    //true: this witness has registered.
        uint index;         //the index of the witness in the address pool, if it is registered
        
        WState state;    //the state of the witness
        
        address SLAContract;    //the contract address of 
        uint confirmDeadline;   //Must confirm the sortition in the state of Candidate. Otherwise, reputation -10.
        int8 reputation;       //the reputation of the witness, the initial value is 100. If it is 0, than it is blocked.
    }

    mapping(address => Witness) witnessPool;    
    address [] public witnessAddrs;    //the address pool of witnesses

    
    struct SortitionInfo{
        bool valid;
        uint curBlockNum;
        uint8 blkNeed;   //how many blocks needed for sortition
    }
    mapping(address => SortitionInfo) SLAContractPool;   //record the requester's initial block number. The sortition will be based on the hash value after this block.
    
    
    //record the provider _who generates a SLA contract of address _contractAddr at time _time
    event SLAContractGen(address indexed _who, uint _time, address _contractAddr);
    
    event WitnessSelected(address indexed _who, uint _index, address _forWhom);
    
    //check whether the register has already registered
    modifier checkRegister(address _register){
        require(!witnessPool[_register].registered);
        _;
    }
    
    //check whether it is a registered witness
    modifier checkWitness(address _witness){
        require(witnessPool[_witness].registered);
        _;
    }
    
    //check whether it is a valid SLA contract
    modifier checkSLAContract(address _sla){
        require(SLAContractPool[_sla].valid);
        _;
    }
    
    /**
     * Provider Interface::
     * This is for the provider to generate a SLA contract
     * */
    function genSLAContract() 
        public 
        returns
        (address)
    {
        CloudSLA newSLAContract = new CloudSLA(this, msg.sender, address(0x0));
        address newSLAContractAddress = address(newSLAContract);
        SLAContractPool[newSLAContractAddress].valid = true; 
        emit SLAContractGen(msg.sender, now, newSLAContractAddress);
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
        checkRegister(msg.sender) 
    {
        // witnessPool[msg.sender].index = witnessAddrs.push(msg.sender) - 1;
        witnessAddrs.push(msg.sender);
        witnessPool[msg.sender].index = witnessAddrs.length - 1;

        witnessPool[msg.sender].state = WState.Offline;
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
        checkSLAContract(msg.sender)
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
        checkSLAContract(msg.sender)
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
            
            if(witnessPool[sAddr].state == WState.Online && witnessPool[sAddr].reputation > 0
               && sAddr != _provider && sAddr != _customer)
            {
                witnessPool[sAddr].state = WState.Candidate;
                witnessPool[sAddr].confirmDeadline = now + 5 minutes;   // 5 minutes for confirmation
                witnessPool[sAddr].SLAContract = msg.sender;
                emit WitnessSelected(sAddr, witnessPool[sAddr].index, msg.sender);
                onlineCounter--;
                wcounter++;
            }
            
            // seed = (uint)(keccak256(uint(seed)));
            seed = (uint)(keccak256(abi.encodePacked(bytes32(seed))));
            
        }
        
        //make this interface cannot be invoked twice without 'request'
        SLAContractPool[msg.sender].curBlockNum = 0;
        return true;
    }
    
    /**
     * Contract Interface::
     * Candidate witness calls the SLA contract and confirm the sortition. 
     * */
    function confirm(address _candidate)
        public
        checkWitness(_candidate)
        checkSLAContract(msg.sender)
        returns 
        (bool)
    {
        //have not reached the confirmation deadline
        require( now < witnessPool[_candidate].confirmDeadline );
        
        //only able to confirm in candidate state
        require(witnessPool[_candidate].state == WState.Candidate);
        
        //only the SLA contract can select it.
        require(witnessPool[_candidate].SLAContract == msg.sender);
        
        witnessPool[_candidate].state = WState.Busy;
        
        return true;
    }
    
    /**
     * Contract Interface::
     * SLA contract ends and witness calls this from the contract to release the Busy witness.
     * If the reputation smaller than 0, the witness will be turned off.
     * */
    function release(address _witness)
        public
        checkWitness(_witness)
        checkSLAContract(msg.sender)
    {
        //only able to release in Busy state
        require(witnessPool[_witness].state == WState.Busy);
        
        //only the SLA contract can operate on it.
        require(witnessPool[_witness].SLAContract == msg.sender);
        
        if(witnessPool[_witness].reputation <= 0){
            witnessPool[_witness].state = WState.Offline;
        }else{
            witnessPool[_witness].state = WState.Online;
            onlineCounter++;
        }
        
    }
    
    /**
     * Contract Interface::
     * Decrease the reputation value. 
     * */
    function reputationDecrease(address _witness, int8 _value)
        public
        checkWitness(_witness)
        checkSLAContract(msg.sender)
    {
        //only able to release in Busy state
        require( _value > 0 );
        
        //only the SLA contract can operate on it.
        require(witnessPool[_witness].SLAContract == msg.sender);
        
        witnessPool[_witness].reputation -= _value;
        
    }
    
    
    
    /**
     * Witness Interface::
     * Reject the sortition for candidate. Because the SLA contract is not valid.
     * */
    function reject()
        public
        checkWitness(msg.sender)
    {
        //only reject in candidate state
        require(witnessPool[msg.sender].state == WState.Candidate);
        
        //have not reached the rejection deadline
        require( now < witnessPool[msg.sender].confirmDeadline );
        
        witnessPool[msg.sender].state = WState.Online;
        onlineCounter++;
    }
    
    /**
     * Witness Interface::
     * Reverse its own state to Online after the confirmation deadline. But need to reduece the reputation. 
     * */
    function reverse()
        public
        checkWitness(msg.sender)
    {
        //must exceed the confirmation deadline
        require( now > witnessPool[msg.sender].confirmDeadline );
        
        //able to turn only in candidate state
        require(witnessPool[msg.sender].state == WState.Candidate);
        
        witnessPool[msg.sender].reputation -= 10;
        
        if(witnessPool[msg.sender].reputation <= 0){
            witnessPool[msg.sender].state = WState.Offline;
        }else{
            witnessPool[msg.sender].state = WState.Online;
            onlineCounter++;
        }
    }
    
    /**
     * Witness Interface::
     * Turn online to wait for sortition. The witness with reputation samller than 0 cannot be turned on.
     * */
    function turnOn()
        public
        checkWitness(msg.sender)
    {
        //must be in the state of offline
        require(witnessPool[msg.sender].state == WState.Offline);
        
        //its reputation must be bigger than 0
        require( witnessPool[msg.sender].reputation > 0 );
        
        witnessPool[msg.sender].state = WState.Online;
        onlineCounter++;
    }
    
    /**
     * Witness Interface::
     * Turn offline to avoid sortition.
     * */
    function turnOff()
        public
        checkWitness(msg.sender)
    {
        
        //must be in the state of online
        require(witnessPool[msg.sender].state == WState.Online);
        
        witnessPool[msg.sender].state = WState.Offline;
        onlineCounter--;
    }
    
    
    /**
     * Witness Interface::
     * For witness itself to check the state of itself and the reputation.
     * If it is selected, following two return values show its confirmation deadline and the address of the SLA contract, who sortited it. 
     * */
    function checkWState(address _witness)
        public
        view
        returns
        (WitnessPool.WState, int8, uint, address)
    {
        return (witnessPool[_witness].state, witnessPool[_witness].reputation, witnessPool[_witness].confirmDeadline, witnessPool[_witness].SLAContract);
    }
    
}



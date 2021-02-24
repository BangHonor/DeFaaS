// SPDX-License-Identifier: MIT

pragma solidity^0.6.0;

import "./WitnessPool.sol";

contract CloudSLA {
    
    enum State { Fresh, Init, Active, Violated, Completed }
    State public SLAState;
    
    WitnessPool public wp;
    
    string public cloudServiceDetail = "";
    
    uint8 public BlkNeeded = 2;
    
    uint public CompensationFee = 500 finney; //0.5 ether
    uint public ServiceFee = 1 ether;
    uint public ServiceDuration = 10 minutes;  
    uint public ServiceEnd = 0;
    
    uint public WF4NoViolation = 10 finney;  //the fee for the witness if there is no violation
    uint WF4Violation = 10*WF4NoViolation;   //the fee for the witness in case there is violation
    uint VoteFee = WF4NoViolation;   //this is the fee for witness to report its violation
    
    uint public WitnessNumber = 3;   //N
    uint public ConfirmNumRequired = 2;   //M: This is a number to indicate how many witnesses needed to confirm the violation
    
    uint SharedFee = (WitnessNumber * WF4Violation)/2;  //this is the maximum shared fee to pay the witnesses
    uint ReportTimeWin = 2 minutes;   //the time window for waiting all the witnesses to report a violation event 
    uint ReportTimeBegin = 0;
    uint ConfirmRepCount = 0;
    
    uint AcceptTimeWin = 2 minutes;   //the time window for waiting the customer to accept this SLA, otherwise the state of SLA is transferred to Completed
    uint AcceptTimeEnd = 0;

    address public Customer;
    uint CustomerBalance = 0;
    uint CPrepayment = ServiceFee + SharedFee;
    
    address public Provider;
    uint ProviderBalance = 0;
    uint PPrepayment = SharedFee;
    
    //this is the balance to reward the witnesses from the committee
    uint SharedBalance = 0;
    
    // this is the witness committee
    address [] public witnessCommittee;
    

    struct WitnessAccount {
        bool selected;   //wheterh it is a member witness committee
        bool violated;   //whether it has reported that the service agreement is violated 
        uint balance;    //the account balance of this witness
    }
    mapping(address => WitnessAccount) witnesses;
    
    //this is to log event that _who modified the SLA state to _newstate at time stamp _time
    event SLAStateModified(address indexed _who, uint _time, State _newstate);
    
    //this is to log event that _witness report a violation at time stamp _time for a SLA monitoring round of _roundID
    event SLAViolationRep(address indexed _witness, uint _time, uint _roundID);
    
    
    constructor(WitnessPool _witnessPool, address _provider, address _customer) public
    {
        Provider = _provider;
        Customer = _customer;
        wp = _witnessPool;
    }
    
    //following functinos are used for setting parameters instead of default ones
    
    
    modifier checkState(State _state){
        require(SLAState == _state);
        _;
    }
    
    modifier checkProvider() {
        require(msg.sender == Provider);
        _;
    }
    
    modifier checkCustomer() {
        require(msg.sender == Customer);
        _;
    }
    
    modifier checkMoney(uint _money) {
        require(msg.value == _money);
        _;
    }
    
    //check whether the sender is a legal witness member in the committee 
    modifier checkWitness() {
        
        require(witnesses[msg.sender].selected);
        _;
    }
    
    modifier checkTimeIn(uint _endTime) {
        require(now < _endTime);
        _;
    }
    
    modifier checkTimeOut(uint _endTime) {
        require(now > _endTime);
        _;
    }
    
    // to ensure all the customers and witnesses has withdrawn money back
    modifier checkAllBalance(){
        require(CustomerBalance == 0);
        
        for(uint i = 0 ; i < witnessCommittee.length ; i++)
            require(witnesses[witnessCommittee[i]].balance == 0);
        
        _;
    }
    
    function setBlkNeeded(uint8 _blkNeed)
        public 
        checkState(State.Fresh) 
        checkProvider
    {
        require(_blkNeed > 1);
        BlkNeeded = _blkNeed;
    }
    
    //the unit is Szabo = 0.001 finney
    function setCompensationFee(uint _compensationFee)
        public 
        checkState(State.Fresh) 
        checkProvider
    {
        require(_compensationFee > 0);
        uint oneUnit = 1 szabo;
        CompensationFee = _compensationFee*oneUnit;
    }
    
    function setServiceFee(uint _serviceFee)
        public 
        checkState(State.Fresh) 
        checkProvider
    {
        require(_serviceFee > 0);
        uint oneUnit = 1 szabo;
        ServiceFee = _serviceFee*oneUnit;
    }
    
    function setWitnessFee(uint _witnessFee)
        public 
        checkState(State.Fresh) 
        checkProvider
    {
        require(_witnessFee > 0);
        uint oneUnit = 1 szabo;
        WF4NoViolation = _witnessFee*oneUnit;
        VoteFee = WF4NoViolation;
    }
    
    //the unit is minutes
    function setServiceDuration(uint _serviceDuration)
        public 
        checkState(State.Fresh) 
        checkProvider
    {
        require(_serviceDuration > 0);
        uint oneUnit = 1 minutes;
        ServiceDuration = _serviceDuration*oneUnit;
    }
    
    //Set the witness committee number, which is the 'N'
    function setWitnessCommNum(uint _witnessCommNum)
        public 
        checkState(State.Fresh) 
        checkProvider
    {
        require(_witnessCommNum > 2);
        require(_witnessCommNum > witnessCommittee.length);
        WitnessNumber = _witnessCommNum;
    }
    
    //Set the 'M' out of 'N' to confirm the violation
    function setConfirmNum(uint _confirmNum)
        public 
        checkState(State.Fresh) 
        checkProvider
    {
        // N/2 < M < N 
        require(_confirmNum > (WitnessNumber/2));
        require(_confirmNum < WitnessNumber);
        
        ConfirmNumRequired = _confirmNum;
    }
    
    //Set the customer address
    function setCustomer(address _customer)
        public 
        checkState(State.Fresh) 
        checkProvider
    {
        Customer = _customer;
    }
    
     // this is for Cloud provider to publish its service detail
    function publishService(string memory _serviceDetail) 
        public 
        checkState(State.Fresh) 
        checkProvider
    {
        cloudServiceDetail = _serviceDetail;
    }
    
    // this is for Cloud provider to set up this SLA and wait for Customer to accept
    function setupSLA() 
        public 
        payable 
        checkState(State.Fresh) 
        checkProvider
        checkMoney(PPrepayment)
    {
        require(WitnessNumber == witnessCommittee.length);
        
        ProviderBalance += msg.value;
        SLAState = State.Init;
        AcceptTimeEnd = now + AcceptTimeWin;
        emit SLAStateModified(msg.sender, now, State.Init);
    }
    
    function cancleSLA()
        public
        checkState(State.Init)
        checkProvider
        checkTimeOut(AcceptTimeEnd)
    {
        if(ProviderBalance > 0){
            msg.sender.transfer(ProviderBalance);
            ProviderBalance = 0;
        }
        
        SLAState = State.Fresh;
        
    }
    
    // this is for customer to put its prepaid fee and accept the SLA
    function acceptSLA() 
        public 
        payable 
        checkState(State.Init) 
        checkCustomer
        checkTimeIn(AcceptTimeEnd)
        checkMoney(CPrepayment)
    {
        require(WitnessNumber == witnessCommittee.length);
        
        CustomerBalance += msg.value;
        SLAState = State.Active;
        emit SLAStateModified(msg.sender, now, State.Active);
        ServiceEnd = now + ServiceDuration;
        
        //transfer ServiceFee from customer to provider 
        ProviderBalance += ServiceFee;
        CustomerBalance -= ServiceFee;
        
        //setup the SharedBalance
        ProviderBalance -= SharedFee;
        CustomerBalance -= SharedFee;
        SharedBalance += SharedFee*2;
    }
    
    /**
     * Customer Interface:
     * Reset the witnesses' state, who have reported.
     * */
    function resetWitness() 
        public 
        checkState(State.Active) 
        checkCustomer
        checkTimeIn(ServiceEnd)
    {
        //some witness has reported the violation
        require(ReportTimeBegin != 0);
        
        //some witness reported, but the violation is not confirmed 
        require(now > ReportTimeBegin + ReportTimeWin);
        
        
        for(uint i = 0 ; i < witnessCommittee.length ; i++){
            if(witnesses[witnessCommittee[i]].violated == true){
                witnesses[witnessCommittee[i]].violated = false;
                SharedBalance += witnesses[witnessCommittee[i]].balance;    //penalty
                witnesses[witnessCommittee[i]].balance = 0;
                wp.reputationDecrease(witnessCommittee[i], 1);  //the reputation value of this witness decreases by 1
            }
        }
        
        ConfirmRepCount = 0;
        ReportTimeBegin = 0;
        
    }
    
    
    function reportViolation()
        public
        payable
        checkTimeIn(ServiceEnd)
        checkWitness 
        checkMoney(VoteFee)
    {
        uint equalOp = 0;   //nonsense operation to make every one using the same gas 
        
        if(ReportTimeBegin == 0)
            ReportTimeBegin = now;
        else
            equalOp = now; 
            
        //only valid within the confirmation time window
        require(now < ReportTimeBegin + ReportTimeWin);
        
        require( SLAState == State.Violated || SLAState == State.Active );
        
        //one witness cannot vote twice 
        require(!witnesses[msg.sender].violated);
        
        witnesses[msg.sender].violated = true;
        witnesses[msg.sender].balance += VoteFee;
        
        ConfirmRepCount++;
        
        //the witness who reports in the last order pay more gas as penalty
        if( ConfirmRepCount >= ConfirmNumRequired ){
            SLAState = State.Violated;
            emit SLAStateModified(msg.sender, now, State.Violated);
        }
        
        emit SLAViolationRep(msg.sender, now, ServiceEnd);
    }
    
    // the customer end the violated SLA and withdraw its compensation
    function customerEndVSLAandWithdraw()
        public
        checkState(State.Violated) 
        checkCustomer
    {
        //end the Service 
        ServiceEnd = now;
        //end the violation reports
        if(now < ReportTimeBegin + ReportTimeWin)
            ReportTimeBegin = now - ReportTimeWin;
         
        for(uint i = 0 ; i < witnessCommittee.length ; i++){
            if(witnesses[witnessCommittee[i]].violated == true){
                witnesses[witnessCommittee[i]].balance += WF4Violation;  //reward the witness who reported this violation
                SharedBalance -= WF4Violation;
            }else{
                wp.reputationDecrease(witnessCommittee[i], 1);  //the reputation value of this witness decreases by 1
                //witnesses[witnessCommittee[i]].balance += WF4NoViolation;  
                //SharedBalance -= WF4NoViolation;
            }
        }
        
        //compensate the customer for service violation
        CustomerBalance += CompensationFee;
        ProviderBalance -= CompensationFee;
        
        // customer and provider divide the remaining shared balance
        if(SharedBalance > 0){
            CustomerBalance += (SharedBalance/2);
            ProviderBalance += (SharedBalance/2);
        }
        SharedBalance = 0;
        
        
        SLAState = State.Completed;
        emit SLAStateModified(msg.sender, now, State.Completed);
        
        if(CustomerBalance > 0){
            msg.sender.transfer(CustomerBalance);
            CustomerBalance = 0;
        }
        
    }
    
    function customerWithdraw()
        public
        checkState(State.Completed)
        checkTimeOut(ServiceEnd)
        checkCustomer
    {
        require(CustomerBalance > 0);
        
        msg.sender.transfer(CustomerBalance);
            
        CustomerBalance = 0;
    }
    
    function providerWithdraw()
        public
        checkState(State.Completed)
        checkTimeOut(ServiceEnd)
        checkProvider
    {
        require(ProviderBalance > 0);
        
        msg.sender.transfer(ProviderBalance);
        
        ProviderBalance = 0;
    }
    
    
    // this means there is no violation during this service. This function needs provider to invoke to end and gain its benefit
    function providerEndNSLAandWithdraw()
        public
        checkState(State.Active)
        checkTimeOut(ServiceEnd)
        checkProvider
    {
        for(uint i = 0 ; i < witnessCommittee.length ; i++){
            if(witnesses[witnessCommittee[i]].violated == true){
                SharedBalance += witnesses[witnessCommittee[i]].balance;   //penalty for the reported witness, might be cheating
                witnesses[witnessCommittee[i]].balance = 0;
                wp.reputationDecrease(witnessCommittee[i], 1);  //the reputation value of this witness decreases by 1
            }else{
                witnesses[witnessCommittee[i]].balance += WF4NoViolation;       // reward the witness
                SharedBalance -= WF4NoViolation;
            }
            
        }
        
        // customer and provider divide the remaining shared balance
        if(SharedBalance > 0){
            CustomerBalance += (SharedBalance/2);
            ProviderBalance += (SharedBalance/2);
        }
        SharedBalance = 0;
        
        SLAState = State.Completed;
        emit SLAStateModified(msg.sender, now, State.Completed);
        
        if(ProviderBalance > 0){
            msg.sender.transfer(ProviderBalance);
            ProviderBalance = 0;
        }
            
        
    }
    
    function witnessWithdraw()
        public
        checkState(State.Completed)
        checkTimeOut(ServiceEnd)
        checkWitness
    {
        require(witnesses[msg.sender].balance > 0);
            
        msg.sender.transfer(witnesses[msg.sender].balance);
        
        witnesses[msg.sender].balance = 0;
        
        
    }
    
    //this only restart the SLA lifecycle, not including the selecting the witness committee. This is to continuously deliver the servce. 
    function restartSLA()
        public
        payable
        checkState(State.Completed)
        checkTimeOut(ServiceEnd)
        checkProvider
        checkAllBalance
        checkMoney(PPrepayment)
    {
        require(WitnessNumber == witnessCommittee.length);
        
        // reset all the related values
        ConfirmRepCount = 0;
        ReportTimeBegin = 0;
        
        //reset the witnesses' state only
        for(uint i = 0 ; i < witnessCommittee.length ; i++){
            if(witnesses[witnessCommittee[i]].violated == true)
                witnesses[witnessCommittee[i]].violated = false;
        }
        
        
        ProviderBalance = msg.value;
        SLAState = State.Init;
        AcceptTimeEnd = now + AcceptTimeWin;
        emit SLAStateModified(msg.sender, now, State.Init);
    }
    
    //this is to flush all the witnesses in the committee. Go back to Fresh state.
    function resetSLA()
        public
        checkState(State.Completed)
        checkTimeOut(ServiceEnd)
        checkProvider
        checkAllBalance
    {
        // in case there are some unexpected errors happen, provider can withdraw all the money back anyway
        if(address(this).balance > 0)
            msg.sender.transfer(address(this).balance);
        
        // reset all the related values
        ConfirmRepCount = 0;
        ReportTimeBegin = 0;
        
        //reset the witness committee
        for(uint i = 0 ; i < witnessCommittee.length ; i++){
            wp.release(witnessCommittee[i]);
            delete witnesses[witnessCommittee[i]];
        }
        
        delete witnessCommittee;
        
        SLAState = State.Fresh;
        emit SLAStateModified(msg.sender, now, State.Fresh);
    }
    
    
    // this is only for debug in case there is some money stuck in the contract
    /*
    function only4DebugWithdraw()
        public
        checkProvider
    {
        if(address(this).balance > 0)
            msg.sender.transfer(address(this).balance);
    }
    
    
    function only4DebugChangeState(State _newstate)
        public
        checkProvider
    {
        SLAState = _newstate;
    }
    */
    
    
    function requestSotition()
        public
        checkProvider
        returns
        (bool success)
    {
        require(wp.request(BlkNeeded));
        return true;
    }
    
    function sortitionFromWP(uint _N)
        public
        checkProvider
        returns
        (bool success)
    {
        
        require(WitnessNumber > witnessCommittee.length);
        
        require(WitnessNumber - witnessCommittee.length >= _N);
        
        require(Customer != address(0x0));
        
        require(wp.sortition(_N, Provider, Customer));
        return true;
    }
    
    function getCommitteeCount()
        public
        view
        returns
        (uint)
    {
        return witnessCommittee.length; 
    }
    
    //the candidate witness confirm itself
    function witnessConfirm()
        public
        returns
        (bool)
    {
        //have not registered in the witness committee
        require(!witnesses[msg.sender].selected);
        
        //The candidate witness can neither be the provider nor the customer
        require(msg.sender != Provider);
        require(msg.sender != Customer);
        
        //confirm with the witness pool
        require(wp.confirm(msg.sender));
        witnessCommittee.push(msg.sender);
        witnesses[msg.sender].selected = true;
        
        return true;
    }
    
    //the witness has the right to leave the SLA contract in following scenarios
    //1. As long as not in the state of 'Active' or 'Violated'
    //2. If it is the state of 'Init', the time should be out of the 'AcceptTimeEnd'
    function witnessRelease()
        public
        checkWitness
    {
        //not in the 'Active', 'Violated' or 'Completed' state
        require(SLAState != State.Active);
        require(SLAState != State.Violated);
        
        require( (SLAState == State.Init && now > AcceptTimeEnd) 
                 || SLAState == State.Completed );
        
        
        uint index = witnessCommittee.length;
        for(uint i = 0 ; i<witnessCommittee.length ; i++){
            if(witnessCommittee[i] == msg.sender)
                index = i;
        }
        require(index != witnessCommittee.length);
        //move the last one in the list to replace the deleted one
        witnessCommittee[index] = witnessCommittee[witnessCommittee.length - 1];

        // [修改之处]
        // witnessCommittee.length--;
            
        delete witnesses[msg.sender];
            
        wp.release(msg.sender);
        
    }
    
}
// SPDX-License-Identifier: MIT

pragma solidity^0.6.0;

import "./Owned.sol";
import "./FaaSTokenPay.sol";
import "./WitnessManagement.sol";
import "./SLA.sol";


contract WitnessPool is Owned, FaaSTokenPay, WitnessManagement {

    enum SLAStates { Confirming, Monitoring, Finished }

    function isAtSLAState(uint _slaID, SLAStates _state) public view returns (bool) {
        return (SLAPool[_slaID].state == _state);
    }

    modifier atSLAtate(uint _slaID, SLAStates _state) {
        require(
            isAtSLAtate(_slaID _state) == true,
            "WitenssPool: function cannot be called at this state"
        );
        _;
    }


    // ------------------------------------------------------------------------------------------------

    // 证人委员信息
    struct MemberInfo {
        bool isSelected;          // 是否被选择为证人委员
        bool isReportViolation;   // 是否报告了 SLA 违约 
    }

    // SLA 执行信息
    struct SLAInfo {
        // 状态
        SLAStates state;     // SLA 状态
        bool isViolated;     // 是否发生违约
        // 监视内容
        string funcToMonitor;   // 监视函数
        // 监视开始时间
        // 监视结束时间
        // 证人委员会的抽选参数
        uint curBlockNum;
        uint8 blkNeed;                   // how many blocks needed for sortition
        // 证人委员会
        uint numWitness;                             // 证人委员会人数
        uint numJudgeViolationRequired;              // 判定违约需要的人数
        address[] committee;                         // 证人委员会
        mapping(address => MemberInfo) memberInfos;  // 证人委员信息        
    }

    // ------------------------------------------------------------------------------------------------

    // SLA 表: slaID => SLA，使用上 slaID 就是 deploymentOrderID
    mapping(uint => SLAInfo) SLAPool;

    // 一次性囚徒困境博弈的 payoff
    uint public rewardViolationReport    = 10;   //  10 token
    uint public rewardViolationSilence   =  0;   //   0 token
    uint public rewardNoviolationReport  =  0;   //   0 token
    uint public rewardNoviolationSilence =  1;   //   1 token

    // 非诚实证人降低的信誉值
    uint public reputationDishonestReduced = 1;
    
    // ------------------------------------------------------------------------------------------------

    constructor(address _tokenContractAddress)
        public
        WitnessManagement(_tokenContractAddress)
    {
    }

    // ------------------------------------------------------------------------------------------------

    event WitnessSelectedEvent(address indexed _witness, uint _index);

    // ------------------------------------------------------------------------------------------------

    function isWitnessCommitteeMember(uint _slaID, address _witness) public view returns (bool) {
        return (SLAPool[_slaID].memberInfos[_witness].isSelected == true);
    }

    modifier witnessCommitteeMember(uint _slaID) {
        require(
            isWitnessCommitteeMember(_slaID, msg.sender),
            "WitnessPool: the witness is not a committee member "
        );
        _;
    }

    // ------------------------------------------------------------------------------------------------


    // Market API
    // 新建一个 SLA 执行信息
    function newSLA(uint _slaID) 
        public 
        onlyOwner
    {
        // 抽选
        uint _curBlockNum = block.number; 
        uint _blkNeed     = 2;  // 参考值

        // SLA 执行信息初始化
        SLAInfo storage _sla = SLAPool[_slaID];
        _sla.state = SLAStates.Monitoring;
        _sla.isViolated = false;
        _sla.funcToMonitor = "";
        _sla.curBlockNum = _curBlockNum;
        _sla.blkNeed = _blkNeed;
        _sla.numWitness = ;
        _sla.numJudgeViolationRequired = ;
        _sla.committee = ;

        // 证人委员信息初始化
        for(uint i = 0; i < _sla.committee.length; i++) {
            address _witness = _sla.committee[i];
            _sla.memberInfos[_witness].isSelected = true;
            _sla.memberInfos[_witness].isReportViolation = false;
        }
    }

    // 证人 API
    // 为指定的 SLA 报告违约
    function reportViolation(uint _slaID)
        public
        witenssRegisterd
        witnessQualified
        witnessCommitteeMember(_slaID)
        atSLAState(_slaID, SLAStates.Monitoring)
        // 时间检查 TODO
    {
        // 报告违约
        SLAPool[_slaID].memberInfos[msg.sender].isReportViolation = true;
    }

    // ------------------------------------------------------------------------------------------------

    // 判断是否违约，并产生证人奖励
    function _judgeViolation(uint _slaID)
        private
        atSLAState(_slaID, SLAStates.Monitoring)
        // 时间检查 TODO
    {
        SLA storage _sla = SLAPool[_slaID];
        
        // 计数
        uint numReportViolation = 0;
        for(uint i = 0; i < _sla.committee.length; i++) {
            address _witness = _sla.committee[i];
            if (_sla.memberInfos[_witness].isReportViolation == true) {
                numReportViolation++;
            }
        }

        // 记录监视结果
        _sla.isViolated = (numReportViolation > _sla.numJudgeViolationRequired);

        // 证人奖励结算
        for(uint i = 0; i < _sla.committee.length; i++) {
            address _witness = _sla.committee[i];

            if(_sla.isViolated == true) {
                // 违约了
                if (_sla.memberInfos[_witness].isReportViolation == true) {
                    // 报告违约
                    witnessPool[_witness].reward += rewardViolationReport;  
                } else{
                    // 未报告违约
                    witnessPool[_witness].reward += rewardViolationSilence;
                    witnessPool[_witness].reputation -= reputationDishonestReduced;  // 减少信誉值
                }
            } else {
                // 未违约
                if (_sla.memberInfos[_witness].isReportViolation == true) {
                    // 报告违约
                    witnessPool[_witness].reward += rewardNoviolationReport;
                    witnessPool[_witness].reputation -= reputationDishonestReduced;  // 减少信誉值
                } else{
                    // 未报告违约
                    witnessPool[_witness].reward += rewardNoviolationSilence;  
                }
            }
        }
    }

    // Matket API
    function isViolatedSLA(uint _slaID)
        public
        // 时间检查
        // 时间转换
        atSLAState(_slaID, SLAStates.Finished)
        returns (bool)
    {
        return SLAPool[_slaID].isViolated;
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



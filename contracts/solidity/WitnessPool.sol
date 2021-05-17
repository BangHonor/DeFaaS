// SPDX-License-Identifier: MIT

pragma solidity>=0.6.0;

import "./Owned.sol";
import "./FaaSTokenPay.sol";
import "./WitnessManagement.sol";
import "./SafeMath.sol";

contract WitnessPool is Owned, FaaSTokenPay, WitnessManagement {

    using SafeMath for uint;   // 注意到 uint 默认是 uint256

    enum SLAStates { Monitoring, Finished }

    function isAtSLAState(uint _slaID, SLAStates _state) public view returns (bool) {
        return (SLAPool[_slaID].state == _state);
    }

    modifier atSLAState(uint _slaID, SLAStates _state) {
        require(
            isAtSLAState(_slaID, _state) == true,
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
        bool isValid;        // 是否是有效的 SLA，用于拒绝无效的访问
        bool isViolated;     // 是否发生违约
        // 监视内容
        string funcPath;      // 监视函数
        uint monitoringBeginTime;  // 监视开始时间
        uint monitoringDuration;   // 监视时长
        // 证人委员会
        uint numReportRequired;                      // 判定违约需要的报告违约证人人数
        address[] committee;                         // 证人委员会，数组长度为证人人数
        mapping(address => MemberInfo) memberInfos;  // 证人委员信息        
    }

    // ------------------------------------------------------------------------------------------------

    // SLA 表: slaID => SLA，在使用时 slaID 就是 deploymentOrderID
    mapping(uint => SLAInfo) SLAPool;

    // 一次性的证人博弈 payoff
    uint public rewardViolationReport    = 10;   // +10
    uint public fineViolationSilence     =  0;   // -0
    uint public fineNoviolationReport    =  1;   // -1
    uint public rewardNoviolationSilence =  1;   // +1

    // 非诚实证人降低的信誉值
    uint public reputationDishonestReduced = 1;

    // SLA 的证人数量标准
    uint public stdNumWitness = 3;    
    // 违约判定的证人数量标准             
    uint public stdNumReportRequired = 2;  
    
    // 抽选需要的区块数量标准
    uint public stdblockNeed = 2;
    
    // ------------------------------------------------------------------------------------------------

    constructor(address _tokenContractAddress)
        WitnessManagement(_tokenContractAddress)
    {}

    // ------------------------------------------------------------------------------------------------

    // 证人被选中通知事件（被选中证人，证人服务的 SLA 的 ID，监视开始时间，监视时长，证人要监视的函数）
    event WitnessSelectedEvent(address indexed _witness, uint _slaID,  uint _monitoringBeginTime, uint _monitoringDuration, string _funcPath);

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

    // -----------------------------------------------------------------------------------------------

    function isValidSLA(uint _slaID) private view returns (bool) {
        return (SLAPool[_slaID].isValid == true);
    }

    modifier validSLA(uint _slaID) {
        require(
            isValidSLA(_slaID),
            "WitnessPool: invalid sla"
        );
        _;
    }

    // ------------------------------------------------------------------------------------------------

    // 在监视时间中
    modifier inMonitoringTime(uint _slaID) {
        require(
            block.timestamp < SLAPool[_slaID].monitoringBeginTime + SLAPool[_slaID].monitoringDuration, 
            "WitnessPool: monitoring time exceeded"
        );
        _;
    }

    // 监视结束检查
    modifier monitoringEndTrigger(uint _slaID) {
        SLAInfo storage _sla = SLAPool[_slaID];
        if (_sla.state == SLAStates.Monitoring && block.timestamp > _sla.monitoringBeginTime + _sla.monitoringDuration) {
            // 监视结束
            _judgeViolation(_slaID);
            _releaseWitnessCommittee(_slaID);
            _sla.state = SLAStates.Finished;  // 改变 SLA 状态
        }
        _;
    }


    // ------------------------------------------------------------------------------------------------

    // Market API
    // 新建一个 SLA 执行信息
    function newSLA(
        uint _curBlockNum,
        uint _slaID, 
        string memory _funcPath,
        uint _monitoringDuration) 
        public 
        // onlyOwner
    {
        // 抽选证人委员会，使用标准的参数
        address[] memory _committee = _sortitionWitnessCommittee(stdNumWitness, _curBlockNum, stdblockNeed);

        // SLA 执行信息初始化
        SLAInfo storage _sla = SLAPool[_slaID];

        _sla.state      = SLAStates.Monitoring;
        _sla.isValid    = true;
        _sla.isViolated = false;
        
        _sla.funcPath       = _funcPath;
        _sla.monitoringBeginTime = block.timestamp + 1 minutes;
        _sla.monitoringDuration  = _monitoringDuration;

        _sla.numReportRequired = stdNumReportRequired;
        _sla.committee         = _committee;

        // 证人委员信息初始化
        for(uint i = 0; i < _sla.committee.length; i++) {
            address _witness = _sla.committee[i];
            _sla.memberInfos[_witness].isSelected = true;
            _sla.memberInfos[_witness].isReportViolation = false;

            // 发出通知
            emit WitnessSelectedEvent(_witness, _slaID, _sla.monitoringBeginTime, _sla.monitoringDuration, _funcPath);
        }
    }

    // 证人 API
    // 为指定的 SLA 报告违约
    function reportViolation(uint _slaID)
        public
        witenssRegisterd
        witnessQualified
        validSLA(_slaID)
        witnessCommitteeMember(_slaID)
        atSLAState(_slaID, SLAStates.Monitoring)
        inMonitoringTime(_slaID)
    {
        // 报告违约
        SLAPool[_slaID].memberInfos[msg.sender].isReportViolation = true;
    }

    // Matket API
    function isViolatedSLA(uint _slaID)
        public
        validSLA(_slaID)
        monitoringEndTrigger(_slaID)
        atSLAState(_slaID, SLAStates.Finished)
        returns (bool)
    {
        // isViolated 已经在 monitoringEndTrigger 中算出
        return SLAPool[_slaID].isViolated;
    }
    

    // ------------------------------------------------------------------------------------------------

    // 判断是否违约，并产生证人奖励
    function _judgeViolation(uint _slaID)
        private
        atSLAState(_slaID, SLAStates.Monitoring)
    {
        SLAInfo storage _sla = SLAPool[_slaID];
        
        // 计数
        uint numReportViolation = 0;
        for(uint i = 0; i < _sla.committee.length; i++) {
            address _witness = _sla.committee[i];
            if (_sla.memberInfos[_witness].isReportViolation == true) {
                numReportViolation++;
            }
        }

        // 记录监视结果
        _sla.isViolated = (numReportViolation > _sla.numReportRequired);

        // 证人奖励结算
        for(uint i = 0; i < _sla.committee.length; i++) {
            address _witness = _sla.committee[i];

            if(_sla.isViolated == true) {
                // 违约了
                if (_sla.memberInfos[_witness].isReportViolation == true) {
                    // 报告违约
                    witnessPool[_witness].reward += rewardViolationReport;  // 获得奖励
                } else{
                    // 未报告违约
                    witnessPool[_witness].depoist    -= fineViolationSilence;        // 扣罚押金         
                    witnessPool[_witness].reputation -= reputationDishonestReduced;  // 减少信誉值
                }
            } else {
                // 未违约
                if (_sla.memberInfos[_witness].isReportViolation == true) {
                    // 报告违约
                    witnessPool[_witness].reward     -= fineNoviolationReport;       // 扣罚押金
                    witnessPool[_witness].reputation -= reputationDishonestReduced;  // 减少信誉值
                } else{
                    // 未报告违约
                    witnessPool[_witness].reward += rewardNoviolationSilence;  // 获得奖励
                }
            }
        }
    }

    // 改变证人委员状态为可以被抽选的 Online 状态
    function _releaseWitnessCommittee(uint _slaID) private {
        for(uint i = 0; i < SLAPool[_slaID].committee.length; i++) {
            address _witness = SLAPool[_slaID].committee[i];
            witnessPool[_witness].state = WStates.Online;
            numOnlineWitness++;
        }
    }

    // 抽选证人委员会
    function _sortitionWitnessCommittee(
        uint _numWitness, uint _curBlockNum, uint _blockNeed) 
        private 
        returns (address[] memory)
    {
        address[] memory _committee;

        // require(
        //     numOnlineWitness > 10 * _numWitness,
        //     "WintnessPool: not enough online witness"
        // );

        // 除数 1024 是一个缩小系数
        uint _pivotBlockNum = block.number - ( block.number - _curBlockNum ) / 1024 - _blockNeed;

        // https://solidity-cn.readthedocs.io/zh/develop/units-and-global-variables.html
        // solidity 文档：基于可扩展因素，区块哈希不是对所有区块都有效。你仅仅可以访问最近 256 个区块的哈希，其余的哈希均为零。
        // 如果 ( block.number - _curBlockNum ) > 1024（缩小系数）* 256（EVM blockhash 深度限制）= 2^18, 那么该 require 会失败
        require(
            block.number < _pivotBlockNum + 255,
            "WitnessPool: blockhash can only be accessed within 255 depth"
        );

        // 抽选算法生成种子要求 _pivotBlockNum 后面 _blockNeed 数量的区块已经产生
        // 事实上，该 require 总是成功的
        require(
            block.number > _pivotBlockNum + _blockNeed,
            "WintnessPool: no more blockNeed blocks generated"
        );

        // 种子随机性： 未来区块的哈希（各方参与行为）
        uint _seed = 0;
        for(uint bi = 0 ; bi < _blockNeed; bi++)
        {
            _seed += (uint)( blockhash( _pivotBlockNum + bi + 1 ) );
        }

        uint _counter = 0;
        while(_counter < _numWitness){

            // 由种子选择一个“随机”证人
            address sAddr = witnessAddrs[_seed % witnessAddrs.length];
            
            if ( isAtWState(sAddr, WStates.Online) && isWitnessQualified(sAddr) )
            {
                witnessPool[sAddr].state = WStates.Working;   // 改变选中的证人状态为 Working
                numOnlineWitness--;

                _counter++;
            }
            
            _seed = (uint)(keccak256(abi.encodePacked(bytes32(_seed))));
        }

        return _committee;
    }

}



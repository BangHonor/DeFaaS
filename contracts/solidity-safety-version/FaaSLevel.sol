// SPDX-License-Identifier: MIT

pragma solidity>=0.6.0;


import "./Owned.sol";

contract FaaSLevel is Owned {

    // 计算资源
    struct ComputingSource {
        uint core;  // CPU 核心数
        uint mem;   // 函数执行内存，以 MB 为单位
        
        // 尚不支持
        // uint packageSize;          // 部署文件包的大小，以 MB 为单位
        // uint gpu;                  // GPU 计算资源，有待量化
        // uint upstreamBandwidth;    // 上行带宽，以 KBps 为单位
        // uint downstreamBandwidth;  // 下行带宽，以 KBps 为单位
    }

    // ------------------------------------------------------------------------------------

    // FaaS 规格编号
    uint public numFaaSLevel;

    modifier validFaaSLevelID(uint _faasLevelID) {
        require(
            _faasLevelID < numFaaSLevel, 
            "FaaSLevel: invalid faaslevel ID"
        );
        _;
    }

    // ------------------------------------------------------------------------------------

    // FaaS 规格表: faasLevelID => ComputingSource
    // 只允许 append，不允许修改已有的 level
    mapping(uint => ComputingSource) private faaSLevels;

    // ------------------------------------------------------------------------------------

    event addFaaSLevelEvent(
        uint indexed _index, 
        uint _core, 
        uint _mem);
        
    // ------------------------------------------------------------------------------------

    constructor() 
    {
        numFaaSLevel = 0;

        // 规格表
        addFaaSLevel(1, 512);
        addFaaSLevel(1, 1024);
        addFaaSLevel(1, 2048);
        addFaaSLevel(2, 1024);
        addFaaSLevel(4, 2048);
    }

    // ------------------------------------------------------------------------------------

    // 添加 FaaS 规格
    function addFaaSLevel(uint _core, uint _mem) 
        public
        onlyOwner
    {
        uint _faasLevelID = numFaaSLevel++;

        faaSLevels[_faasLevelID] = ComputingSource({
            core: _core,
            mem: _mem
        });
        
        emit addFaaSLevelEvent(_faasLevelID, _core, _mem);
    }

    // ------------------------------------------------------------------------------------

    // 查询 FaaS 规格编号对应的计算机资源
    // 参数（ 规格编号 levelID ）
    // 返回（ levelID 是否有效，core, mem ）
    function getFaaSLevel(uint _faasLevelID) 
        public 
        view
        validFaaSLevelID(_faasLevelID)
        returns (uint, uint) 
    {
        ComputingSource storage cs = faaSLevels[_faasLevelID];
        return (cs.core, cs.mem);
    }
    
}
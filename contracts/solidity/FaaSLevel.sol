// SPDX-License-Identifier: MIT

pragma solidity^0.6.0;


import "./Owned.sol";

contract FaaSLevel is Owned {

    // 计算资源
    struct ComputingSourceLevel {
        uint core;  // CPU 核心数
        uint mem;   // 函数执行内存，以 MB 为单位
        
        // 尚不支持
        // uint packageSize;          // 部署文件包的大小，以 MB 为单位
        // uint gpu;                  // GPU 计算资源，难以分割
        // uint upstreamBandwidth;    // 上行带宽，以 KBps 为单位
        // uint downstreamBandwidth;  // 下行带宽，以 KBps 为单位
    }

    // FaaS 规格表
    uint private numFaaSLevels;
    mapping(uint => ComputingSourceLevel) private faaSLevels;

    event addFaaSLevelEvent(uint indexed index, uint core, uint mem);

    constructor() 
        public 
    {
        numFaaSLevels = 0;

        // 规格表
        addFaaSLevel(1, 512);
        addFaaSLevel(1, 1024);
        addFaaSLevel(1, 2048);
        addFaaSLevel(2, 1024);
        addFaaSLevel(4, 2048);
    }

    // 查询 FaaS 规格的数量
    // 参数 空
    // 返回（ 规格的数量 numFaaSLevels ）
    function getFaaSLevelNumber()
        public
        view
        returns (uint)
    {
        return numFaaSLevels;
    }

    // 查询 FaaS 规格编号对应的计算机资源
    // 参数（ 规格编号 levelID ）
    // 返回（ levelID 是否有效，core, mem ）
    function getFaaSLevel(uint _levelID) 
        public 
        view
        returns (bool, uint, uint) 
    {
        if (_levelID >= numFaaSLevels) {
            // invalid levelID
            return (false, 0, 0);
        }

        ComputingSourceLevel memory cs = faaSLevels[_levelID];
        return (true, cs.core, cs.mem);
    }

    // 添加 FaaS 规格
    // 参数（ core, mem ）
    // 返回（ 规格编号 _levelID ）
    function addFaaSLevel(uint _core, uint _mem) 
        public
        onlyOwner
        returns (uint) 
    {

        uint _levelID = numFaaSLevels++;

        faaSLevels[_levelID] = ComputingSourceLevel({
            core: _core,
            mem: _mem
        });
        
        emit addFaaSLevelEvent(_levelID, _core, _mem);

        return (_levelID);
    }
    
}
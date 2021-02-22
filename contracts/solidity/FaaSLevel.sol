pragma solidity>=0.4.25 <0.6.11;


contract FaaSLevel {

    struct ComputingSourceLevel {
        uint core;  // CPU 核心数
        uint mem;   // 内存大小，以 MB 为单位
    }

    // FaaS 规格表
    uint numLevels;
    mapping(uint => ComputingSourceLevel) levels;

    // 管理者
    address public owner;

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this.");
        _;
    }

    event newFaaSLevelEvent(uint index, uint core, uint mem);

    constructor(address _owner) public {
        owner = _owner;
        numLevels = 0;

        // 规格表
        newFaaSLevel(1, 512);
        newFaaSLevel(1, 1024);
        newFaaSLevel(1, 2048);
        newFaaSLevel(2, 1024);
        newFaaSLevel(4, 2048);
    }

    // 查询 FaaS 规格的数量
    // 参数 空
    // 返回（ 规格的数量 numLevels ）
    function getFaaSLevelNumber()
        public
        view
        returns (uint)
    {
        return numLevels;
    }

    // 查询 FaaS 规格编号对应的计算机资源
    // 参数（ 规格编号 levelID ）
    // 返回（ levelID 是否有效，core, mem ）
    function getFaaSLevel(uint _levelID) 
        public 
        view
        returns (bool, uint, uint) 
    {
        if (_levelID >= numLevels) {
            // invalid levelID
            return (false, 0, 0);
        }

        ComputingSourceLevel memory cs = levels[_levelID];
        return (true, cs.core, cs.mem);
    }

    // 新建 FaaS 规格
    // 参数（ core, mem ）
    // 返回（ 规格编号 _levelID ）
    function newFaaSLevel(uint _core, uint _mem) 
        public
        onlyOwner
        returns (uint) 
    {

        uint _levelID = numLevels++;

        levels[_levelID] = ComputingSourceLevel({
            core: _core,
            mem: _mem
        });
        
        emit newFaaSLevelEvent(_levelID, _core, _mem);

        return (_levelID);
    }
    
}
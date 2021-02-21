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

    constructor(address _owner) {

        owner = _owner;
        numLevels = 0;

        // init
        newFaaSLevel(1, 512);
        newFaaSLevel(1, 1024);
        newFaaSLevel(1, 2048);
        newFaaSLevel(2, 1024);
        newFaaSLevel(4, 2048);
    }

    function getFaaSLevelNumber()
        public
        view
        returns (uint)
    {
        return numLevels;
    }

    // 查询 FaaS 规格编号对应的计算机资源
    function getFaaSLevel(uint _levelID) 
        public 
        view
        returns (uint, uint) 
    {
        require(
            _levelID < numLevels,
            "levelID is NOT less than numLevels, index is out of range."
        );

        ComputingSourceLevel cs = levels[_levelID];
        return (cs.core, cs.mem);
    }

    // 新建 FaaS 规格
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
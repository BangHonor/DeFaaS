// SPDX-License-Identifier: MIT

pragma solidity^0.6.0;

contract ProviderPool {

    // 标准的供应商押金
    uint public stdProviderDeposit;
    // 供应商押金表
    mapping(address => uint) private providerDeposits;
    // 供应商信誉初始值，是新注册的供应商的初始信誉值
    uint public providerReputationInit;  
    // 供应商信誉合格值，低于合格值为不合格
    uint public providerReputationQualified;  
    // 供应商信誉值表
    // 信誉值为 0 表示未注册
    mapping (address => uint) private providerReputations;

    constructor() public {
        // 供应商参数初始化 TODO
        stdProviderDeposit = 100; // 100 token
        providerReputationInit = 5;
        providerReputationQualified = 5;
    }

    // ------------------------------------------------------------------------------------

    // 查询是否已注册为供应商
    function isProviderRegistered(address _provider) public view returns (bool) {
        // 信誉值为 0 表示供应商未注册, 大于 0 表示供应商已注册
        return (providerReputations[_provider] > 0);
    }

    // 查询供应商是否合格
    function isProviderQualified(address _provider) public view returns (bool) {
        // 低于合格值为不合格
        // 即，信誉值在区间 (0, providerReputationQualified) 的供应商为不合格的
        // 不合格的供应商既不能注销供应商资格取回押金，也不能参与竞价
        return (providerReputations[_provider] >= providerReputationQualified);
    }

    // 已注册
    modifier providerRegistered() {
        require(
            isProviderRegistered(msg.sender) == true,
            "Market: the address had not been registered"
        );
        _;
    }

    // 未注册
    modifier providerUnregistered() {
        require(
            isProviderRegistered(msg.sender) == false,
            "Market: the address had been registered"
        );
        _;
    }

    // 合格
    modifier providerQualified() {
        require(
            isProviderQualified(msg.sender),
            "Market: the provider is not qualified"
        );
        _;
    }

    // 查询指定供应商的押金
    function getProviderDeposit(address _provider) 
        public
        view 
        providerRegistered
        returns (uint) 
    {
        return providerDeposits[_provider];
    } 

    // 查询指定供应商的信誉值
    function getProviderReputation(address _provider) 
        public 
        view
        providerRegistered 
        returns (uint) 
    {
        return providerReputations[_provider];
    }

    // 注册供应商资格
    function _providerLogin() 
        internal
        providerUnregistered 
    {
        providerDeposits[msg.sender] = stdProviderDeposit;
        providerReputations[msg.sender] = providerReputationInit;
    }

    // 注销供应商资格
    function _providerLogout() 
        internal 
        providerRegistered
        providerQualified
    {
        // 信誉值清零
        providerReputations[msg.sender] = 0;
        // 押金清零
        providerDeposits[msg.sender] = 0;
    }

}

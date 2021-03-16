// SPDX-License-Identifier: MIT

pragma solidity>=0.6.0;

import "./FaaSTokenPay.sol";

contract ProviderManagement is FaaSTokenPay {

    // 信誉值为 0 表示未注册
    struct Provider {
        uint reputation;  // 信誉值 
        uint deposit;    // 押金
    }

    // 供应商押金标准
    uint private stdProviderDeposit;
    // 供应商信誉初始值，是新注册的供应商的初始信誉值
    uint private providerReputationInit;  
    // 供应商信誉合格值，低于合格值为不合格
    uint private providerReputationQualified;
    // 供应商表
    mapping(address => Provider) providers;

    // ------------------------------------------------------------------------------------

    constructor(address _tokenContractAddress) 
        FaaSTokenPay(_tokenContractAddress)
    {
        // 供应商参数初始化 TODO
        stdProviderDeposit = 100; // 100 token
        providerReputationInit = 5;
        providerReputationQualified = 5;
    }

    // ------------------------------------------------------------------------------------

    // 查询是否已注册为供应商
    function isProviderRegistered(address _provider) public view returns (bool) {
        // 信誉值为 0 表示供应商未注册, 大于 0 表示供应商已注册
        return (providers[_provider].reputation > 0);
    }

    // 查询供应商是否合格
    function isProviderQualified(address _provider) public view returns (bool) {
        // 低于合格值为不合格
        // 即，信誉值在区间 (0, providerReputationQualified) 的供应商为不合格的
        // 不合格的供应商既不能注销供应商资格取回押金，也不能参与竞价
        return (providers[_provider].reputation >= providerReputationQualified);
    }

    // 已注册
    modifier providerRegistered(address _provider) {
        require(
            isProviderRegistered(_provider) == true,
            "Market: the address had not been registered"
        );
        _;
    }

    // 未注册
    modifier providerUnregistered(address _provider) {
        require(
            isProviderRegistered(_provider) == false,
            "Market: the address had been registered"
        );
        _;
    }

    // 合格
    modifier providerQualified(address _provider) {
        require(
            isProviderQualified(_provider),
            "Market: the provider is not qualified"
        );
        _;
    }

    // ------------------------------------------------------------------------------------

    function getStdProviderDeposit() public view returns (uint) {
        return stdProviderDeposit;
    }

    // 查询指定供应商的押金
    function getProviderDeposit(address _provider) 
        public
        view 
        providerRegistered(_provider)
        returns (uint) 
    {
        return providers[_provider].deposit;
    } 

    // 查询指定供应商的信誉值
    function getProviderReputation(address _provider) 
        public 
        view
        providerRegistered(_provider)
        returns (uint) 
    {
        return providers[_provider].reputation;
    }

    // ------------------------------------------------------------------------------------

    // 供应商 API
    // 注册供应商资格，支付押金
    function providerLogin() 
        public
        providerUnregistered(msg.sender)
    {
        // 注册
        providers[msg.sender].deposit    = stdProviderDeposit;
        providers[msg.sender].reputation = providerReputationInit;

        // 支付押金，供应商应预先进行授权转账操作
        require(
            tokenContract.transferFrom(msg.sender, address(this), providers[msg.sender].deposit) == true,
            "Market: failed to pay a register deposit"
        );    
    }

    // 供应商 API
    // 注销供应商资格，取回押金
    function providerLogout() 
        public
        providerRegistered(msg.sender)
        providerQualified(msg.sender)
    {
        // 检查：记录押金
        uint _depoist = providers[msg.sender].deposit;
        // 生效：注销
        providers[msg.sender].deposit    = 0;  // 信誉值清零
        providers[msg.sender].reputation = 0;  // 押金清零
        // 交互：退还押金
        require(
            tokenContract.transfer(msg.sender, _depoist) == true,
            "Market: failed to refund the provider deposit"
        );
    }

}

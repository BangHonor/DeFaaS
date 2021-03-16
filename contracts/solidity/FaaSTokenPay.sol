// SPDX-License-Identifier: MIT

import "./FaaSToken.sol";

pragma solidity>=0.6.0;

contract FaaSTokenPay {

    FaaSToken public tokenContract;

    constructor(address _tokenContractAddress) public {
        // 合约地址初始化
        tokenContract = FaaSToken(_tokenContractAddress);
    }
}
// SPDX-License-Identifier: MIT

pragma solidity>=0.6.0;

contract Owned {

    // 所有者
    address public owner;

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this.");
        _;
    }

    constructor() {
        owner = msg.sender;
    }
}
// SPDX-License-Identifier: MIT

pragma solidity^0.6.0;

contract HelloWorld {
    string value;

    constructor() public {
        value = "Hello, World!";
    }

    function get() public view returns (string memory) {
        return value;
    }

    function set(string memory v) public {
        value = v;
    }
}
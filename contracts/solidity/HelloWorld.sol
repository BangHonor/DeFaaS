pragma solidity>=0.4.25 <0.6.11;

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
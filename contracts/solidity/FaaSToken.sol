// SPDX-License-Identifier: MIT

pragma solidity>=0.6.0;

import "./ERC20.sol";
import "./Context.sol";


// 可参考 https://soliditydeveloper.com/mocking-contracts

contract FaaSToken is Context, ERC20 {

    constructor() 
        ERC20("FaaSToken", "FST") 
        public 
    {
        // TODO
    }

    // 铸币给自己，自己地址的 toekn 增加 amount
    // 注意：测试阶段，未添加任何限制
    function mint(uint amount) public virtual {
        _mint(_msgSender(), amount);
    }

    // 铸币给指定地址
    // 注意：测试阶段，未添加任何限制
    function mintTo(address account, uint amount) public virtual {
        _mint(account, amount);
    }

    // 继承自 ERC20 的转帐限制函数
    function _beforeTokenTransfer(address from, address to, uint256 amount) internal virtual override {
        // 测试阶段，还不加什么限制
    }

}
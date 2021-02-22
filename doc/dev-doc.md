# tricks


## ERC20

- https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20.md
- https://blog.51cto.com/13784902/2324024
- （标准实现）https://github.com/OpenZeppelin/openzeppelin-contracts/tree/master/contracts/token/ERC20
- （可用实现）https://github.com/vittominacori/erc20-generator/tree/master/dist

-------------

## 使用代币支付

- https://ethfans.org/ajian1984/articles/816

-------------


## 盲拍合约

- https://solidity-cn.readthedocs.io/zh/develop/solidity-by-example.html#id5
- https://solidity-cn.readthedocs.io/zh/develop/common-patterns.html#id5


-------------


## 客户端调用合约

1. 连接到区块链（client）
2. 构造合约实例（instance），需要地址（address），连接（client）
3. 调用合约函数

-------------


## 如何写出高质量设计说明文档

### 文档编写之『术』

『术』主要就是编写文档时结构上所需要素，这里分为了三类：

- 业务背景：首先编写文档时不要预设他人能理解所有的专业术语，其次要清楚地介绍业务存在的痛点，区块链能否解决这些痛点。
- 技术方案设计：需清晰地说明基础业务需求（如参与者、场景、活动），设计思路概要，详细的合约职责、函数、参数等。
- 使用说明：实际使用场景、上手指南和使用手册等。

### 文档编写之『道』

有了术，就能确保文档内容结构的完整性，文档就有了骨架和血肉；但是，『道』才是文档的灵魂和精髓所在，这里整理了5个关注点：

- 亮点：展示出本方案的独特性（创意/设计/功能/规范/文档）；
- 痛点：说清楚问题和解决方法；
- 重点：代码、注释、文档编写要条理清晰、可读性强，围绕问题和解决方案展开，不要为了炫技而炫技；
- 难点：基于智能合约特性，权衡规范、效率和安全性；
- 要点：以解释作为基本导向，不要预设别人能够理解所有业务和技术术语。

本文主要分享了存证和积分两个典型应用场景合约设计思路和实例代码解析，为大家总结了相关的开发技巧和文档写作技巧。 在智能合约开发过程中，开发者需根据实际业务需求选择适用的技巧与方案。正所谓『兵无常势，水无常形』，没有最优的设计，只有最合适的设计。


-------------


## 坑

- 坑一：solidity 版本 和 solc 版本
- 坑二：FISCO-BCOS 提供的 go-sdk 功能不完整
- 坑三：安全漏洞
    - 智能合约的整数溢出
    - 智能合约的恶意重入
    - 非对称加密
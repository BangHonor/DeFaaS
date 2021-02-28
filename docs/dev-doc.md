# dev-doc

## 开发计划

- 修改 Market
- 修改 WP/SLA
- 联调 Matket 和 WP/SLA
- 设计命令行工具 cli，使用 cobra 编写
- 根据命令行工具 cli 的需要，写客户端

## 经济模型（是否可行）

- 以给 maintainer 奖励的方式来发币
- 以给 witness 奖励的方式来发币

## Provider 客户端设计

- bidder Pool
- accessSecretKey Pool

## ERC20

- https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20.md
- https://blog.51cto.com/13784902/2324024
- （标准实现）https://github.com/OpenZeppelin/openzeppelin-contracts/tree/master/contracts/token/ERC20
- （可用实现）https://github.com/vittominacori/erc20-generator/tree/master/dist

-------------

## 使用代币支付

- https://ethfans.org/ajian1984/articles/816

-------------


## 市场合约

资金转帐流程：
1. 租户调用 token 合约的 approve，授权足够额度给市场合约
2. 调用市场合约的新建部署订单函数
3. 市场合约调用 token 合约的 transferFrom，锁定预付款
4. 租户查看订单匹配是否成功
5. 若成功，租户向供应商发出部署请求
6. 证人完成报告并且服务时间到，订单完成
7. 订单完成后，租户调用市场合约的 withdraw，取回多出的预付款和可能有的违约款
8. 订单完成后，供应商可以从市场合约中提取酬劳

---------

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
    - 引用类型（结构体，数组，映射）需要显式指定数据位置
    - 对于引用类型，状态变量向 storage 局部变量赋值时仅仅传递一个引用
    - 多重继承已最后继承的为主，故从左往右从祖先写到父亲
    - 外部创建合约 SC 的 sender 是 SC 的部署者，内部创建合约 SC 的部署者是 new SC 的合约
- 坑二：FISCO-BCOS 提供的 go-sdk 功能不完整
    - 0.10.3 版本的 go-sdk 不支持合约事件监听。。。
    - java-sdk 令人羡慕
- 坑三：无处不在的安全漏洞
    - 智能合约的整数溢出（合约里的算术运算都是不安全的！！！）
        - SafeMath
    - 智能合约的恶意重入（还没有检查可能的 bug！！！）
        - 在你合约中状态变量进行各种变化后再调用外部函数，这样，你的合约就不会轻易被滥用的重入 (reentrancy) 所影响
        - 多发生再 withdraw 函数
        - 解决方法：“检查-生效-交互”（Checks-Effects-Interactions）模式
    - 非对称加密（该加密的地方还没有加密！！！）
- 坑四：token 的 approveAndCall 是糟糕的转帐方式
- 坑五：配置地狱
    - 区块链节点的IP地址，端口配置
    - 客户端的证书配置
        - ca.crt CA 根证书文件
        - sdk.crt SDK 证书文件
        - sdk.key SDK 私钥文件
    - 公钥账户配置
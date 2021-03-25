# dev-doc

## 开发计划

- 修改 Market `Done`
- 修改 WP/SLA `Done`
- 联调 Matket 和 WP/SLA `Done`
- 改进 WP 合约的 report 为秘密报告
- 单元测试 Market 合约函数
- 单元测试 WP 合约函数
- 设计命令行工具 cli，使用 cobra 编写
- 根据命令行工具 cli 的需要，写客户端
- 链端-客户端集成测试
- 生产环境部署
- 生产环境测试
- 制图
  - Order 状态图
  - 链端-客户端交互图
  - FaaS 用例图
- docker 为函数实体


## 实现细节

### Provider 客户端设计

- bidder Pool
- accessSecretKey Pool

### ERC20

- https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20.md
- https://blog.51cto.com/13784902/2324024
- （标准实现）https://github.com/OpenZeppelin/openzeppelin-contracts/tree/master/contracts/token/ERC20
- （可用实现）https://github.com/vittominacori/erc20-generator/tree/master/dist


### 使用代币支付

- https://ethfans.org/ajian1984/articles/816



### 客户端调用合约

1. 连接到区块链（client）
2. 构造合约实例（instance），需要地址（address），连接（client）
3. 调用合约函数



### 客户端监听合约事件

- https://goethereumbook.org/zh/event-subscribe/

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


-------------


## 坑

### solidity

- solidity 版本
- 引用类型（结构体，数组，映射）需要显式指定数据位置
- 对于引用类型，状态变量向 storage 局部变量赋值时仅仅传递一个引用
- 多重继承已最后继承的为主，故从左往右从祖先写到父亲
- 外部创建合约 SC 的 sender 是 SC 的部署者，内部创建合约 SC 的部署者是 new SC 的合约
- 在智能合约中，结构体struct无法直接作为返回值直接返回，如果需要返回结构体的数据，必须逐一返回，并在接收处完成数据的拼装。


### 安全漏洞

- 整数溢出
    - SafeMath
- 恶意重入
    - 在你合约中状态变量进行各种变化后再调用外部函数，这样，你的合约就不会轻易被滥用的重入 (reentrancy) 所影响
    - 多发生在 withdraw 函数
    - 解决方法：“检查-生效-交互”（Checks-Effects-Interactions）模式
- 非对称加密



## eth


### gas

nonce: 记录发起交易的账户已执行交易总数。Nonce的值随着每个新交易的执行不断增加，这能让网络了解执行交易需要遵循的顺序，并且作为交易的重放保护。

gasPrice:该交易每单位gas的价格，Gas价格目前以Gwei为单位（即10^9wei），其范围是大于0.1Gwei，可进行灵活设置

gasLimit:该交易支付的最高gas上限。该上限能确保在出现交易执行问题（比如陷入无限循环）之时，交易账户不会耗尽所有资金。一旦交易执行完毕，剩余所有gas会返还至交易账户。

to：该交易被送往的地址（调用的合约地址或转账对方的账户地址）。
value：交易发送的以太币总量。

data: 若该交易是以太币交易，则data为空；若是部署合约，则data为合约的bytecode；若是合约调用，则需要从合约ABI中获取函数签名，并取函数签名hash值前4字节与所有参数的编码方式值进行拼接而成，具体参见文章https://github.com/linjie-1/guigulive-operation/wiki/Ethereum%E7%9A%84%E5%90%88%E7%BA%A6ABI%E6%8B%93%E5%B1%95


### wait for pedding

相关讨论

https://medium.com/kinblog/go-ethereum-pending-wait-for-it-balance-936a9fb1c6e2


truffle 的实现

https://github.com/trufflesuite/truffle/blob/500fdcc35202839ac785f7b0fe3de623cbcac074/packages/deployer/src/deployment.js#L79

## TODO

### 生产环境中的安全配置

开发环境下，使用的连接是不安全。

在生产环境，需要做一些安全工作：

- 配置 SSL

    - 把 http 配置为 https

    - 把 ws 配置为 wss


    解决方案：反向代理 https://ethereum.stackexchange.com/questions/26026/how-to-ssl-ethereum-geth-node


- 访问限制

    - http.corsdomain 的跨域限制

    - ws.origins 的访问限制

    解决方案：仍然通过代理服务器来处理，不在 geth 节点这一层解决。



-------------





git config --local http.proxy 127.0.0.1:8889 && git push origin master

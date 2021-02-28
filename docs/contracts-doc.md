# contracts-doc

## 1. 合约

| 合约名        | 描述                                                         |
| ------------- | ------------------------------------------------------------ |
| BlindAuction  | 秘密拍卖合约（盲拍）                                         |
| Context       | 工具类，提供安全的上下文获取                                 |
| ERC20         | IERC20 的实现，是符合 ERC20 标准的 token 实现                |
| FaaSLevel     | FaaS 规格表，一个规格编号对应一组服务参数                    |
| FaaSToken     | 继承自 ERC20 的 token 合约                                   |
| FaaSTokePay   | 抽象合约，使用 FaaSToken 支付的合约                          |
| IERC20        | ERC20 标准的接口定义                                         |
| Market        | 市场合约                                                     |
| Owned         | 抽象合约，所有者合约                                         |
| ProviderPool  | 管理 Provider 的合约                                         |
| SafeMath      | 工具类，库合约，提供 uint256 的安全运算                      |
| SimpleAuction | 简单拍卖合约（公开拍卖）                                     |
| SLA           | 执行 SLA  的合约                                             |
| WitnessPool   | 证人池合约，有三个功能：管理证人，生成 SLA 合约，为 SLA 合约抽取证人 |
|               |                                                              |







## 2. Market 合约

### 2.1 管理部署订单的状态

部署订单的状态有：

- Bidding
- Confirming
- Deploying
- Fulfilling
- Settled
- Finished



### 2.2 Function 的部署信息

#### 2.2.1 HTTP Handler

- Function 的形式是 HTTP Handler 函数，调用一个 Function 方法是访问它的 URL。

  ```
  http://[endAddr]/[funcPath]
  ```

-  组成一个 URL 需要两个数据，一个是 `endAddr`，一个是 `funcPath`

  - `endAddr` 是部署了 Function 的端的地址。

    `endAddr` 可以是 IP地址加上端口地址，如 `127.0.0.1:6666`，也可以是一个域名 。

  - `funcPath` 是一个路径

    如 `endAddr` 为一个域名`www.xxx.com`， `funcPath` 为 `add`，调用一次 Function，即发送一个 HTTP Request 到 `http://www.xxx.com/add`



#### 2.2.2 发布部署信息

部署订单需要 Customer 和 Provider 双方提供部署信息。

- `Comfirming` 阶段
  - Customer 需要发布以下信息：
    -  `funcPath`，通常会是一个哈希值
  - Provider 提供以下信息：
    - `endAddr`
    - `deploymentServerAddr` 接受部署请求的服务器 `deploymentServer` 的地址
    - `accessSecretKey` 访问密钥，一个经过 Customer 公钥加密的哈希值




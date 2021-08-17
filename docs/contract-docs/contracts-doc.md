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




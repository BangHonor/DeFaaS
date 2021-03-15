https://geth.ethereum.org/docs/interface/private-network
https://cloud.tencent.com/developer/article/1559634
https://blog.csdn.net/u010921136/article/details/107222259

├── config: 链配置，新区块的出块规则依赖链配置。
├────── chainId: 即区块链网络 id，用于连接其他节点，不同 id 网络的节点无法相互连接。以太坊公网 id 为 1。
├── alloc: 即分配，创世初始账户的资产配置，直接将账户资产写入区块中。
├── coinbase: 即「生成交易」中的 from。
├── difficulty: 挖矿难度系数，与出块速度成负相关。
├── extraData: 额外数据。
├── gasLimit: 十六进制，燃料上限。
├── nonce: 随机数。
├── parentHash: 母区块的哈希值。
└── timestamp： UTC 时间戳。










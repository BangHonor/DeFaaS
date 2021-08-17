# cmd-doc


## 一键 编译智能合约

把合约放进 `defaas/contracts/solidity`，在 `defaas/` 目录下执行以下命令，可以编译合约，并生成 Go 语言合约。

```bash
go run dev-cmd/gen/main.go -name=合约名
```

如，
```bash
go run dev-cmd/gen/main.go -name=HelloWorld
```

生成的 Go 语言合约以 pkg 的形式保存在 `defaas/contracts/go` 目录下，可以在 Go 代码中调用。




-----------

## 一键 部署合约上链

TODO 首先，需要填写配置。

```bash
go run dev-cmd/deploy/main.go
```

--------


## 启动开发模式的 eth private network

注意：
- **开发模式** （dev mode）的 eth private network 会在关闭后清空区块数据，其次启动将从零开始。
- 开发模式生成的账户密码为空，即 ""

```bash
go run dev-cmd/private-chain-dev/main.go
```


## 一键部署 private network

构建创世区块

```bash
go run dev-cmd/private-chain/main.go -action=build
```


```
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
```


启动一个节点，形成仅有一个节点的网络

```bash
go run dev-cmd/private-chain/main.go -action=run
```


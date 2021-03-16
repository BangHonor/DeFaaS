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


启动一个节点，形成仅有一个节点的网络

```bash
go run dev-cmd/private-chain/main.go -action=run
```


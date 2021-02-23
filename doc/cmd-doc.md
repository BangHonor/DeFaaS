# cmd-doc


## 一键 编译智能合约

把合约放进 `defaas/contracts/solidity`，在 `defaas/` 目录下执行以下命令，可以编译合约，并生成 Go 语言合约。

```bash
go run cmd/gen/main.go -name=合约名
```

如，
```bash
go run cmd/gen/main.go -name=HelloWorld
```

生成的 Go 语言合约以 pkg 的形式保存在 `defaas/contracts/go` 目录下，可以在 Go 代码中调用。

一键编译的方式是整合了下面的多个步骤。


-----------

## 一键 部署合约上链

根据 config 配置完成集成部署。

```bash
go run cmd/deployContracts/main.go
```

--------
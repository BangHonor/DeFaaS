# tools-doc

## 生成私钥与地址

文档 https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/manual/account.html

在 `defaas/` 目录下执行以下命令，可以生成私钥与地址。

```bash
bash ./tools/get_account.sh
```

生成的私钥文件保存在 `defaas/accounts` 目录下。

---------------

## 编译

- 一键编译

把合约放进 `defaas/contracts/solidity`，在 `defaas/` 目录下执行以下命令，可以编译合约，并生成 Go 语言合约。

```bash
go run cmd/gen/main.go -name=ShaTest
```

生成的 Go 语言合约以 pkg 的形式保存在 `defaas/contracts/go` 目录下，可以在 Go 代码中调用。

一键编译的方式是整合了下面的多个步骤。

- 多个步骤

文档 https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/sdk/go_sdk/contractExamples.html

`solc`

```bash
bash ./tools/download_solc.sh -v 0.4.25
bash ./tools/download_solc.sh -v 0.5.2
bash ./tools/download_solc.sh -v 0.6.10
```

`abigen`

由 `solidity` 文件编译出 `.bin` 和 `.abi` 文件

```bash
./tools/solc-0.6.10 --bin --abi -o ./contracts/bin-abi/ ./contracts/solidity/HelloWorld.sol
```

由 `.bin` 和 `.abi` 文件生成 `.go` 代码

```bash
mkdir -p ./contracts/go/helloworld 
```

```bash
./tools/abigen --bin ./contracts/bin-abi/HelloWorld.bin --abi ./contracts/bin-abi/HelloWorld.abi --pkg helloworld --type HelloWorld --out ./contracts/go/helloworld/HelloWorld.go
```

---------------

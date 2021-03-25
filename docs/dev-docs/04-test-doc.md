# test-doc


## 1. 单元测试

### 1.1 合约函数测试

#### 1.1.1 功能性测试


#### 1.1.2 安全性测试

- 检查整数溢出


### 1.2 客户端测试

#### 1.2.1 功能性测试


#### 1.2.2 安全性测试

- 检查负数输入

https://pkg.go.dev/math/big#Int.Sign


------------------------------

## 2. 集成测试

### 2.1 模拟环境

go-ethereum 提供了一种模拟区块链的方法。

可以使用这种方法来模拟集成测试。

https://duckduckgo.com/?sites=geth.ethereum.org&kz=-1&q=Blockchain+simulator&ia=web

```go

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
)

func main() {

    // setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc)

    // commit all pending transactions
	blockchain.Commit()
}

```

### 2.2 开发环境

可以在运行一个开发模式的单节点 eth 私有网络来进行测试，开发模式会自动对提交 pedding 事务的挖矿。



## 2.3 私链环境


可以运行一个 eth 私有网络，可视为一个封闭的生产环境。

在生产环境中，运行 eth 网络需要有挖矿的节点，对机器的性能要求较高。


## 2.4 公链测试网络


接入公链的测试网络来测试，门槛较高，需要费用。


------------------------------






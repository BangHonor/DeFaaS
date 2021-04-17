package main

// 自动化部署合约

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-cmd/cmd"

	"defaas/contracts/go/faastoken"
	"defaas/contracts/go/market"
	"defaas/core/config"

	devutils "defaas/dev-cmd/utils"
)

func init() {
	// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetFlags(log.Lmicroseconds | log.Llongfile)
}

const (
	ethClientRawURL          = "ws://127.0.0.1:8546"
	workDir                  = "/home/dds/kitchen/defaas"
	defaasConfigFilePath     = workDir + "/" + "defaas-config.toml"
	deployerKeyStoreFileName = "UTC--2021-04-16T17-39-18.307832917Z--aea14ff60c1584b7b8d78847f7c0a2cf87350f44"
	deployerKeyStorePassword = "123456"
	deployerKeyStoreFilePath = workDir + "/" + "private-chain/data-0/keystore" + "/" + deployerKeyStoreFileName
)

// go run dev-cmd/deploy/main.go
func main() {
	genContracts()
	deployContracts()
}

// ----------------------------------------------------------------------

func genContracts() {

	contractNames := []string{
		"FaaSToken",
		"Market",
		"WitnessPool",
	}

	for _, name := range contractNames {

		devutils.RunCmd(cmd.NewCmd(
			"go",
			"run",
			"dev-cmd/gen/main.go",
			"-name="+name))

		fmt.Println()
	}
}

func GetAuthFromKeyStore(keyStoreFilePath, password string, client *ethclient.Client) (*bind.TransactOpts, error) {

	// 获取 chainID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	// 加载私钥
	keyjson, err := ioutil.ReadFile(keyStoreFilePath)
	if err != nil {
		return nil, err
	}

	// 构造一个授权事务
	auth, err := bind.NewTransactorWithChainID(bytes.NewBuffer(keyjson), password, chainID)
	if err != nil {
		return nil, err
	}

	// auth.GasPrice = big.NewInt(1000000000 + 1)
	// auth.GasLimit = (1 << 30)

	return auth, nil
}

func deployContracts() {

	// 构造一个连接
	client, err := ethclient.Dial(ethClientRawURL)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := GetAuthFromKeyStore(deployerKeyStoreFilePath, deployerKeyStorePassword, client)
	if err != nil {
		log.Fatal(err)
	}

	faastokenContractAddress, tx, _, err := faastoken.DeployFaaSToken(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[deploy] FaaSToken contract is deployed at", faastokenContractAddress)
	_ = tx
	// fmt.Println("[deploy] Transaction Hash of deployment", tx.Hash())

	marketContractAddress, tx, marketInstance, err := market.DeployMarket(auth, client, faastokenContractAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[deploy] Market contract is deployed at", marketContractAddress)
	_ = tx
	// fmt.Println("[deploy] Transaction Hash of deployment", tx.Hash())

	// wait for mined
	time.Sleep(1 * time.Second)

	witnesspoolContractAddress, err := marketInstance.WpContract(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[deploy] WitnessPool contract is deployed at", witnesspoolContractAddress)

	// 写入配置文件
	if err := config.WriteContractAddress(
		defaasConfigFilePath,
		faastokenContractAddress,
		marketContractAddress,
		witnesspoolContractAddress); err != nil {
		log.Fatal(err)
	}
	fmt.Println("[deploy] contract address is written to file", defaasConfigFilePath)
}

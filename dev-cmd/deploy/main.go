package main

// 自动化部署合约

import (
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-cmd/cmd"

	"defaas/contracts/go/faastoken"
	"defaas/contracts/go/market"
	"defaas/core/config"

	corehelper "defaas/core/helper"
	devutils "defaas/dev-cmd/utils"
)

const (
	// dev config
	ethClientRawURL          = "ws://127.0.0.1:8546"
	deployerKeyStoreFilePath = "./tmp/dev/data/keystore/UTC--2021-03-22T16-13-33.324948003Z--18f3ed9cc0c0727390d38202bd74446c0ded47a0"
	deployerPassword         = ""
	defaasConfigFilePath     = "./defaas-config.toml"
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

func deployContracts() {

	// 构造一个连接
	client, err := ethclient.Dial(ethClientRawURL)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := corehelper.GetAuthFromKeyStore(deployerKeyStoreFilePath, deployerPassword, client)
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

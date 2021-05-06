package main

// 自动化部署合约

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-cmd/cmd"

	"defaas/contracts/go/faastoken"
	"defaas/contracts/go/market"
	defaasconfig "defaas/core/config"

	devutils "defaas/dev-cmd/utils"
)

func init() {
	// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetFlags(log.Lmicroseconds | log.Llongfile)
}

// ----------------------------------------------------------------------

type deployConfig struct {
	ethClientRawURL          string
	defaasConfigFilePath     string
	deployerKeyStoreFilePath string
	deployerKeyStorePassword string
}

// https://geth.ethereum.org/docs/interface/private-network
// go run dev-cmd/deploy/main.go
// geth attach --datadir ./private-chain/data-0
func main() {

	genContracts()

	deployContracts(getDeployConfig())
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

func getAuthFromKeyStore(keyStoreFilePath, password string, client *ethclient.Client) (*bind.TransactOpts, error) {

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

func (dcfg deployConfig) String() string {

	var b strings.Builder

	b.WriteString("deploy contract config\n")
	b.WriteString(fmt.Sprintf("ethClientRawURL [%v]\n", dcfg.ethClientRawURL))
	b.WriteString(fmt.Sprintf("defaasConfigFilePath [%v]\n", dcfg.defaasConfigFilePath))
	b.WriteString(fmt.Sprintf("deployerKeyStoreFilePath [%v]\n", dcfg.deployerKeyStoreFilePath))
	b.WriteString(fmt.Sprintf("deployerKeyStorePassword [%v]\n", dcfg.deployerKeyStorePassword))

	return b.String()
}

func getDeployConfig() *deployConfig {

	dcfg := &deployConfig{}

	dcfg.ethClientRawURL = "ws://127.0.0.1:8546"

	workDir := "/home/kitchen/ktichent-defaas/defaas"
	dcfg.defaasConfigFilePath = path.Join(workDir, "defaas-config.toml")

	deployerKeyStoreDir := path.Join(workDir, "private-chain/data-0/keystore")
	names, err := devutils.ReadDirNames(deployerKeyStoreDir)
	if err != nil {
		log.Fatal(err)
	}

	dcfg.deployerKeyStoreFilePath = path.Join(deployerKeyStoreDir, names[0])
	dcfg.deployerKeyStorePassword = ""

	return dcfg
}

func deployContracts(dcfg *deployConfig) {

	fmt.Println(dcfg)

	// 构造一个连接
	client, err := ethclient.Dial(dcfg.ethClientRawURL)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := getAuthFromKeyStore(dcfg.deployerKeyStoreFilePath, dcfg.deployerKeyStorePassword, client)
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

	fmt.Println("wait for mined ...")
	time.Sleep(10 * time.Second)

	witnesspoolContractAddress, err := marketInstance.WpContract(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[deploy] WitnessPool contract is deployed at", witnesspoolContractAddress)

	// 写入配置文件
	if err := defaasconfig.WriteContractAddress(
		dcfg.defaasConfigFilePath,
		faastokenContractAddress,
		marketContractAddress,
		witnesspoolContractAddress); err != nil {
		log.Fatal(err)
	}
	fmt.Println("[deploy] contract address is written to config file", dcfg.defaasConfigFilePath)
}

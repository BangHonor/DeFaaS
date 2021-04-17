package basic

import (
	"log"
	"math/big"
	"path"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	devutils "defaas/dev-cmd/utils"
)

type testConfig struct {
	testDeFaaSConfigFilePath string
	testKeyStoreFilePath     string
	testKeyStorePassword     string
}

func getTestConfig() *testConfig {

	cfg := &testConfig{}

	workDir := "/home/dds/kitchen/defaas"

	cfg.testDeFaaSConfigFilePath = path.Join(workDir, "defaas-config.toml")
	testKeyStoreDir := path.Join(workDir, "private-chain/data-0/keystore")
	names, err := devutils.ReadDirNames(testKeyStoreDir)
	if err != nil {
		log.Fatal(err)
	}

	cfg.testKeyStoreFilePath = path.Join(testKeyStoreDir, names[0])
	cfg.testKeyStorePassword = ""

	return cfg
}

func getTestBasicClientFromFile() (*BasicClient, error) {

	tcfg := getTestConfig()

	return NewBasicClientWithFile(
		tcfg.testDeFaaSConfigFilePath,
		tcfg.testKeyStoreFilePath,
		tcfg.testKeyStorePassword)
}

func TestNewBasicClient(t *testing.T) {

	client, err := getTestBasicClientFromFile()
	if err != nil {
		t.Fatal(err)
	}

	_ = client
}

func TestComfirmTxByPolling(t *testing.T) {

	assert := assert.New(t)

	client, err := getTestBasicClientFromFile()
	if err != nil {
		t.Fatal(err)
	}

	confirmTxFunc := client.ComfirmTxByPolling

	senderAddress := client.Key.Address
	senderBalanceBefore, err := client.FaaSToken.BalanceOf(senderAddress)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(senderBalanceBefore)
	t.Logf("sender [%v], balance [%v]\n", senderAddress, senderBalanceBefore)

	mintAmount := big.NewInt(1000)
	txMint, err := client.FaaSToken.Mint(mintAmount)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(txMint)

	if err := confirmTxFunc(txMint.Hash(), NumBlockToWaitRecommended); err != nil {
		t.Fatal(err)
	}

	senderBalanceAfter, err := client.FaaSToken.BalanceOf(senderAddress)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(senderBalanceAfter)
	t.Logf("sender [%v], balance [%v]\n", senderAddress, senderBalanceAfter)

	recipientAddress := common.HexToAddress("0x01")
	transferAmount := big.NewInt(1)

	beforeBalance, err := client.FaaSToken.BalanceOf(recipientAddress)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(beforeBalance)

	txTransfer, err := client.FaaSToken.Transfer(recipientAddress, transferAmount)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(txTransfer)

	if err := confirmTxFunc(txTransfer.Hash(), NumBlockToWaitRecommended); err != nil {
		t.Fatal(err)
	}

	afterBalance, err := client.FaaSToken.BalanceOf(recipientAddress)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(afterBalance)

	assert.Equal(0, afterBalance.Cmp(
		big.NewInt(0).Add(
			beforeBalance,
			transferAmount)))
}

func TestComfirmTxBySubscription(t *testing.T) {

	// 如果在用于开发的私有链下测试,
	// 注意私有链不会主动挖矿出块,
	// 因此, 需要在 geth 节点的控制台执行 miner.start() ,
	// 关闭挖矿的则在控制台执行 miner.stop() .

	assert := assert.New(t)

	client, err := getTestBasicClientFromFile()
	if err != nil {
		t.Fatal(err)
	}

	confirmTxFunc := client.ComfirmTxBySubscription

	senderAddress := client.Key.Address
	senderBalanceBefore, err := client.FaaSToken.BalanceOf(senderAddress)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(senderBalanceBefore)
	t.Logf("sender [%v], balance [%v]\n", senderAddress, senderBalanceBefore)

	mintAmount := big.NewInt(1000)
	txMint, err := client.FaaSToken.Mint(mintAmount)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(txMint)

	if err := confirmTxFunc(txMint.Hash(), NumBlockToWaitRecommended); err != nil {
		t.Fatal(err)
	}

	senderBalanceAfter, err := client.FaaSToken.BalanceOf(senderAddress)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(senderBalanceAfter)
	t.Logf("sender [%v], balance [%v]\n", senderAddress, senderBalanceAfter)

	recipientAddress := common.HexToAddress("0x01")
	transferAmount := big.NewInt(1)

	beforeBalance, err := client.FaaSToken.BalanceOf(recipientAddress)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(beforeBalance)

	txTransfer, err := client.FaaSToken.Transfer(recipientAddress, transferAmount)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(txTransfer)

	if err := confirmTxFunc(txTransfer.Hash(), NumBlockToWaitRecommended); err != nil {
		t.Fatal(err)
	}

	afterBalance, err := client.FaaSToken.BalanceOf(recipientAddress)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(afterBalance)

	assert.Equal(0, afterBalance.Cmp(
		big.NewInt(0).Add(
			beforeBalance,
			transferAmount)))
}

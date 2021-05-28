package basic

import (
	"context"
	"defaas/core/helper"
	"defaas/tests/testconfig"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func getTestBasicClientFromFile() *BasicClient {

	tcfg := testconfig.GetTestConfig()

	client, err := NewBasicClientWithFile(
		tcfg.TestDeFaaSConfigFilePath,
		tcfg.TestKeyStoreFilePath,
		tcfg.TestKeyStorePassword)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func TestNewBasicClient(t *testing.T) {

	client := getTestBasicClientFromFile()
	_ = client
}

func TestConfirmTxByPolling(t *testing.T) {

	assert := assert.New(t)

	client := getTestBasicClientFromFile()

	client.FaaSTokenEventSub(context.TODO())

	confirmTxFunc := client.ConfirmTxByPolling

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

	if err := confirmTxFunc(txMint.Hash(), helper.NumBlockToWaitRecommended); err != nil {
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

	if err := confirmTxFunc(txTransfer.Hash(), helper.NumBlockToWaitRecommended); err != nil {
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

func TestConfirmTxBySubscription(t *testing.T) {

	assert := assert.New(t)

	client := getTestBasicClientFromFile()

	client.FaaSTokenEventSub(context.TODO())

	confirmTxFunc := client.ConfirmTxBySubscription

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

	if err := confirmTxFunc(txMint.Hash(), helper.NumBlockToWaitRecommended); err != nil {
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

	if err := confirmTxFunc(txTransfer.Hash(), helper.NumBlockToWaitRecommended); err != nil {
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

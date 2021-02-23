package main

import (
	"fmt"
	"log"
	"math/big"

	// "defaas/contracts/go/faastoken"
	"defaas/contracts/go/faastoken"
	"defaas/utils"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
)

func main() {

	// load config
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]

	// client connect...
	client, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}

	// deploy the contract
	address, tx, _instance, err := faastoken.DeployFaaSToken(client.GetTransactOpts(), client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract address: ", address.Hex()) // the address should be saved
	fmt.Println("transaction hash: ", tx.Hash().Hex())
	_ = _instance

	// load the contract
	contractAddress := common.HexToAddress(address.Hex())
	instance, err := faastoken.NewFaaSToken(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	faastokenSession := faastoken.FaaSTokenSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	// ---------------------------

	account, _ := utils.GetAddressFromConfig(config)
	amount := big.NewInt(100) // 100 token

	// before mint
	totalSupplyBeforeMint, err := faastokenSession.TotalSupply()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("> total supply before mint [%v]\n", totalSupplyBeforeMint)

	balanceOfBeforeMint, err := faastokenSession.BalanceOf(account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("> balance of [%s] before mint is [%v]\n", account.Hex(), balanceOfBeforeMint)

	// mint
	tx, receipt, err := faastokenSession.MintTo(account, amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("mint [%v] tokens to [%s]\n", amount, account.Hex())
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

	// after mint
	totalSupplyAfterMint, err := faastokenSession.TotalSupply()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("> total supply after mint [%v]\n", totalSupplyAfterMint)

	balanceOfAfterMint, err := faastokenSession.BalanceOf(account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("balance of [%s] after mint is [%v]\n", account.Hex(), balanceOfAfterMint)

}
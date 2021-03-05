package main

import (
	"fmt"
	"log"
	"math/big"

	// "defaas/contracts/go/faastoken"
	"defaas/contracts/go/faastoken"
	"defaas/dev-cmd/utils"

	sdkClient "github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
)

func main() {

	// load config
	configs, err := conf.ParseConfigFile("client-config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]

	// client connect...
	client, err := sdkClient.Dial(config)
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
	account, _ := utils.GetAddressFromConfig(config)
	amount := big.NewInt(100) // 100 token

	// ---------------------------

	// 铸币测试
	fmt.Printf("\n\n\n")
	{
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
		fmt.Printf("> balance of [%s] after mint is [%v]\n", account.Hex(), balanceOfAfterMint)
	}

	// 授权测试，转帐测试
	fmt.Printf("\n\n\n")
	{
		spender, spenderPrivateKey := utils.GetAddressFromGenerating()
		approveAmount := big.NewInt(20)     // 20 tokens
		transferFromAmount := big.NewInt(5) // 5 tokens

		// approve
		{
			tx, receipt, err := faastokenSession.Approve(spender, approveAmount)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("> [%s] approves [%v] tokens to spender [%s]\n", account.Hex(), amount, spender.Hex())
			fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
			fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())
		}

		// query
		{
			balanceOfSpender, err := faastokenSession.BalanceOf(spender)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("> balance of spender [%s] is [%v]\n", spender.Hex(), balanceOfSpender)

			allowance, err := faastokenSession.Allowance(account, spender)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("> allowance of [%s] to spender [%s] is [%v]\n", account.Hex(), spender.Hex(), allowance)
		}

		// spendr client
		spenderConfig := *config
		spenderConfig.PrivateKey = spenderPrivateKey

		spenderClient, err := sdkClient.Dial(&spenderConfig)
		if err != nil {
			log.Fatal(err)
		}

		spenderInstance, err := faastoken.NewFaaSToken(contractAddress, spenderClient)
		if err != nil {
			log.Fatal(err)
		}

		spenderFaastokenSession := faastoken.FaaSTokenSession{
			Contract:     spenderInstance,
			CallOpts:     *spenderClient.GetCallOpts(),
			TransactOpts: *spenderClient.GetTransactOpts(),
		}

		// transferFrom
		tx, receipt, err := spenderFaastokenSession.TransferFrom(account, spender, transferFromAmount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("> transfer [%v] tokens from [%s] to [%s]\n", transferFromAmount, account.Hex(), spender.Hex())
		fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
		fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

		// query
		{
			balanceOfSpender, err := spenderFaastokenSession.BalanceOf(spender)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("> After transferFrom, balance of spender [%s] is [%v]\n", spender.Hex(), balanceOfSpender)

			allowance, err := spenderFaastokenSession.Allowance(account, spender)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("> After transferFrom, allowance of [%s] to spender [%s] is [%v]\n", account.Hex(), spender.Hex(), allowance)
		}

		// tansfer
		fmt.Printf("\n\n\n")
		{
			transferAmount := big.NewInt(1)
			tx, receipt, err := spenderFaastokenSession.Transfer(account, transferAmount)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("> transfer [%v] tokens from [%s] to [%s]\n", transferAmount, spender.Hex(), account.Hex())
			fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
			fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())
		}

		// query
		{
			balanceOfSpender, err := spenderFaastokenSession.BalanceOf(spender)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("> balance of spender [%s] is [%v]\n", spender.Hex(), balanceOfSpender)

			balanceOfAccount, err := spenderFaastokenSession.BalanceOf(account)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("> balance of account [%s] after mint is [%v]\n", account.Hex(), balanceOfAccount)
		}
	}
}

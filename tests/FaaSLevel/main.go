package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"

	"defaas/contracts/go/faaslevel" // import faaslevel

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
)

// go run tests/FaaSLevel/main.go
func main() {

	// load config
	configs, err := conf.ParseConfigFile("client-config.toml")
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
	// owner, _ := utils.GetAddressFromConfig(config)
	address, tx, _instance, err := faaslevel.DeployFaaSLevel(client.GetTransactOpts(), client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract address: ", address.Hex()) // the address should be saved
	fmt.Println("transaction hash: ", tx.Hash().Hex())
	_ = _instance

	// load the contract
	contractAddress := common.HexToAddress(address.Hex())
	instance, err := faaslevel.NewFaaSLevel(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	faaslevelSession := faaslevel.FaaSLevelSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	// list FaaS levels
	numLevels, err := faaslevelSession.GetFaaSLevelNumber()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("num levels: ", numLevels)

	for i := big.NewInt(0); i.Cmp(numLevels) == -1; i.Add(i, big.NewInt(1)) {
		isValid, core, mem, err := faaslevelSession.GetFaaSLevel(i)
		if err != nil {
			log.Fatal(err)
		}
		if !isValid {
			log.Fatal(errors.New("invalid level id"))
		}
		fmt.Printf("level %v: %v core, %v MB\n", i, core, mem)
	}

	// add FaaS level
	newCore := big.NewInt(4)
	newMem := big.NewInt(4096)
	tx, receipt, err := faaslevelSession.AddFaaSLevel(newCore, newMem)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("add FaaS level %v core, %v MB\n", newCore, newMem)
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

	// -----------------------------------------------

	fmt.Println("After adding FaaS level...")
	numLevels, err = faaslevelSession.GetFaaSLevelNumber()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("num levels: ", numLevels)

	for i := big.NewInt(0); i.Cmp(numLevels) == -1; i.Add(i, big.NewInt(1)) {
		isValid, core, mem, err := faaslevelSession.GetFaaSLevel(i)
		if err != nil {
			log.Fatal(err)
		}
		if !isValid {
			log.Fatal(errors.New("invalid level id"))
		}

		fmt.Printf("level %v: %v core, %v MB\n", i, core, mem)
	}

	// error
	isValid, _, _, err := faaslevelSession.GetFaaSLevel(big.NewInt(10000))
	if err != nil {
		log.Fatal(err)
	}
	if !isValid {
		log.Println("invalid level id")
	}

	// bind.ContractFilterer{}
	// faaslevel.NewFaaSLevelFilterer()
}

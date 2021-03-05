package main

import (
	"defaas/clients/config"
	"fmt"
	"log"
)

// go run tests/defaas-config/main.go
func main() {
	dfc, err := config.ParseConfigFile("./defaas-config.toml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dfc.FaaSTokenContractAddress.Hex())
	fmt.Println(dfc.MarketContractAddress.Hex())
	fmt.Println(dfc.WitnessPoolContractAddress.Hex())
}

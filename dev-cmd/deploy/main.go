package main

// 自动化部署合约

import (
	"defaas/dev-cmd/utils"
	"fmt"

	"github.com/go-cmd/cmd"
)

// ----------------------------------------------------------------------

// go run dev-cmd/gen/main.go -name=FaaSToken
// go run dev-cmd/gen/main.go -name=Market
func genContracts() {

	contractNames := []string{
		"FaaSToken",
		"Market",
	}

	for _, name := range contractNames {

		fmt.Printf("generate contracts [ %s ] \n", name)

		utils.RunCmd(cmd.NewCmd(
			"go",
			"run",
			"cmd/gen/main.go",
			"-name="+name))

		fmt.Println()
	}
}

// ----------------------------------------------------------------------

func deployContracts() {
	// TODO
}

// ----------------------------------------------------------------------

// go run cmd/deployContracts/main.go

func main() {
	genContracts()
	deployContracts()
}

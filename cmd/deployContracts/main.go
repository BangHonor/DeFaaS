package main

import (
	"defaas/utils"
	"fmt"

	"github.com/go-cmd/cmd"
)

// ----------------------------------------------------------------------

// go run cmd/gen/main.go -name=FaaSLevel
// go run cmd/gen/main.go -name=FaaSToken

func genContracts() {

	contractNames := []string{
		"FaaSLevel",
		"FaaSToken",
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
}

// ----------------------------------------------------------------------

// go run cmd/deployContracts/main.go

func main() {
	genContracts()
	deployContracts()
}

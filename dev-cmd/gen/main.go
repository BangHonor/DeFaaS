package main

// 自动化编译合约，生成合约

import (
	"flag"
	"fmt"
	"log"
	"path"
	"strings"

	"defaas/dev-cmd/utils"

	"github.com/go-cmd/cmd"
)

func init() {
	// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetFlags(log.Lmicroseconds | log.Llongfile)
}

const (
	solidityDir = "./contracts/solidity"
	binAbiDir   = "./contracts/bin-abi"
	goDir       = "./contracts/go"
)

// e.g.
// go run dev-cmd/gen/main.go -name=Market

func main() {

	var smartContractName string

	flag.StringVar(&smartContractName, "name", "HelloWorld", "smart contract name")
	flag.Parse()

	binName := path.Join(binAbiDir, smartContractName+".bin")
	abiName := path.Join(binAbiDir, smartContractName+".abi")
	lowerSmartContractName := strings.ToLower(smartContractName)
	goPkgDir := path.Join(goDir, lowerSmartContractName) // Go语言的 pkg 路径只能小写
	goName := path.Join(goPkgDir, smartContractName+".go")

	// Create Cmd, buffered output
	binAbiGenCmd := cmd.NewCmd(
		"./tools/solc-0.6.10",
		"--bin",
		"--abi",
		"-o",
		binAbiDir,
		path.Join(solidityDir, smartContractName+".sol"),
		"--overwrite")

	utils.RunCmd(binAbiGenCmd)

	fmt.Printf("[gen] %s \n", binName)
	fmt.Printf("[gen] %s \n", abiName)

	utils.RunCmd(cmd.NewCmd("mkdir", "-p", goPkgDir))

	goGenCmd := cmd.NewCmd(
		"./tools/abigen",
		"--bin", binName,
		"--abi", abiName,
		"--pkg", lowerSmartContractName,
		"--type", smartContractName,
		"--out", goName)

	utils.RunCmd(goGenCmd)

	fmt.Printf("[gen] %s \n", goName)
}

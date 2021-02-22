package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"path"
	"strings"

	"defaas/utils"

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

func runCmd(_cmd *cmd.Cmd) {

	fmt.Print(_cmd.Name)
	for _, arg := range _cmd.Args {
		fmt.Print(" ", arg)
	}
	fmt.Println()

	// Run and wait for Cmd to return Status
	status := <-_cmd.Start()

	// Print each line of STDOUT from Cmd
	for _, line := range status.Stdout {
		fmt.Println(line)
	}

	if status.Exit != 0 {
		// Print each line of STDERR from Cmd
		for _, line := range status.Stderr {
			fmt.Println(line)
		}

		err := errors.New("runCmd")
		log.Fatal(err)
	}
}

func main() {

	var smartContractName string

	flag.StringVar(&smartContractName, "name", "HelloWorld", "smart contract name")
	flag.Parse()

	binName := path.Join(binAbiDir, smartContractName+".bin")
	abiName := path.Join(binAbiDir, smartContractName+".abi")
	lowerSmartContractName := strings.ToLower(smartContractName)
	goPkgDir := path.Join(goDir, lowerSmartContractName) // Go语言的 pkg 路径只能小写
	goName := path.Join(goPkgDir, smartContractName+".go")

	utils.RunCmd(cmd.NewCmd("rm", "-f", binName))
	utils.RunCmd(cmd.NewCmd("rm", "-f", abiName))

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

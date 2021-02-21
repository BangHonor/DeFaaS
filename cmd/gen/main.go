package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"path"
	"strings"

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
	goPkgDir := path.Join(goDir, lowerSmartContractName)
	goName := path.Join(goPkgDir, smartContractName+".go")

	runCmd(cmd.NewCmd("rm", "-f", binName))
	runCmd(cmd.NewCmd("rm", "-f", abiName))

	// Create Cmd, buffered output
	binAbiGenCmd := cmd.NewCmd(
		"./tools/solc-0.4.25",
		"--bin",
		"--abi",
		"-o",
		binAbiDir,
		path.Join(solidityDir, smartContractName+".sol"))

	runCmd(binAbiGenCmd)

	fmt.Printf("[gen] %s \n", binName)
	fmt.Printf("[gen] %s \n", abiName)

	runCmd(cmd.NewCmd("mkdir", "-p", goPkgDir))

	goGenCmd := cmd.NewCmd(
		"./tools/abigen",
		"--bin", binName,
		"--abi", abiName,
		"--pkg", lowerSmartContractName,
		"--type", smartContractName,
		"--out", goName)

	runCmd(goGenCmd)

	fmt.Printf("[gen] %s \n", goName)
}

package utils

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-cmd/cmd"
)

// RunCmd ...
func RunCmd(_cmd *cmd.Cmd) {

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

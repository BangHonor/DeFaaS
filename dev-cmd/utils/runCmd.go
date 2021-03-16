package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-cmd/cmd"
)

// 运行命令行

// RunCmd ...
func RunCmd(_cmd *cmd.Cmd) {

	fmt.Print("[dev-cmd] ")
	fmt.Print(_cmd.Name)
	for _, arg := range _cmd.Args {
		fmt.Print(" ", arg)
	}
	fmt.Println()

	// listen SIGINT and SIGTERM from terminal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// run the cmd
	statusChan := _cmd.Start()

	select {
	case status := <-statusChan:
		{
			if status.Exit != 0 {

				// Print each line of STDERR from Cmd
				for _, line := range status.Stderr {
					fmt.Println(line)
				}
				log.Fatal(errors.New("RunCmd"))

			} else {

				// Print each line of STDOUT from Cmd
				for _, line := range status.Stdout {
					fmt.Println(line)
				}

			}
		}
	case <-sigs:
		{
			_cmd.Stop()
		}
	}
}

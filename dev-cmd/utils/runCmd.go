package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	statusChan := _cmd.Start()

	// Print last line of stdout every 2s
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker.C {
			status := _cmd.Status()
			n := len(status.Stdout)
			if n > 0 {
				fmt.Println(status.Stdout[n-1])
			}
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case status := <-statusChan:
		{
			if status.Exit != 0 {
				// Print each line of STDERR from Cmd
				for _, line := range status.Stderr {
					fmt.Println(line)
				}

				err := errors.New("runCmd")
				log.Fatal(err)
			} else {
				fmt.Println("[dev-cmd] done")
			}
		}
	case <-sigs:
		{
			_cmd.Stop()
		}
	}
}

package main

// 运行开发 eth 私有网络

import (
	"fmt"
	"log"

	devutils "defaas/dev-cmd/utils"

	"github.com/go-cmd/cmd"
)

func init() {
	// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetFlags(log.Lmicroseconds | log.Llongfile)
}

const (
	devChainDir = "./tmp/dev/data"
	httpRPCAddr = "127.0.0.1"
)

func main() {

	devNodeCmd := cmd.NewCmd(
		"geth",

		"--datadir", devChainDir,

		"--dev",
		"--identity", "dev-node",

		"--http",
		"--http.addr", "127.0.0.1",
		"--http.port", "8545",
		"--http.api", "eth,web3,miner,admin,personal,net",
		"--http.corsdomain", "*",

		"--ws",
		"--ws.addr", "127.0.0.1",
		"--ws.port", "8546",
		"--ws.api", "eth,web3,miner,admin,personal,net",
		"--ws.origins", "*")

	devutils.RunCmd(devNodeCmd)
	fmt.Println("[dev-cmd] run dev done")

}

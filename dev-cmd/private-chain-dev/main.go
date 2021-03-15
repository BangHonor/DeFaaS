package main

// 运行开发 eth 私有网络

import (
	"fmt"
	"log"

	"defaas/dev-cmd/utils"

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

// go run dev-cmd/private-chain-dev/main.go

// dev 模式生成的账户密码为空，即 ""

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
		"--http.corsdomain", "*")

	utils.RunCmd(devNodeCmd)
	fmt.Println("[dev-cmd] run dev done")

}

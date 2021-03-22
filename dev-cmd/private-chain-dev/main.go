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

		// "--mine",
		// "--miner.threads", "1",
		// "--miner.gasprice", "0",
		// "--miner.gastarget", "0",
		// --mine                              Enable mining
		// --miner.threads value               Number of CPU threads to use for mining (default: 0)
		// --miner.notify value                Comma separated HTTP URL list to notify of new work packages
		// --miner.gasprice value              Minimum gas price for mining a transaction (default: 1000000000)
		// --miner.gastarget value             Target gas floor for mined blocks (default: 8000000)
		// --miner.gaslimit value              Target gas ceiling for mined blocks (default: 8000000)
		// --miner.etherbase value             Public address for block mining rewards (default = first account) (default: "0")
		// --miner.extradata value             Block extra data set by the miner (default = client version)
		// --miner.recommit value              Time interval to recreate the block being mined (default: 3s)
		// --miner.noverify                    Disable remote sealing verification

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

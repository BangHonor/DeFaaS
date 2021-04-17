package main

// 自动构建，运行生产环境 eth 私有网络

import (
	"flag"
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
	chainNetworkID    = "666666"
	chainDir          = "./private-chain"
	genesisDataDir    = chainDir + "/" + "data-0"
	genesisConfigFile = chainDir + "/" + "genesis.json"
)

// 开发模式
const (
	devChainDir = "./tmp/dev/data"
)

func main() {

	var actionName string

	flag.StringVar(&actionName, "action", "run", "action")
	flag.Parse()

	if actionName == "build" {

		// 构建创世区块

		devutils.RunCmd(cmd.NewCmd(
			"rm", "-rf", genesisDataDir))

		genesisCmd := cmd.NewCmd(
			"geth",
			"--datadir", genesisDataDir,
			"init", genesisConfigFile)

		devutils.RunCmd(genesisCmd)
		fmt.Println("[dev-cmd] build genesis done")

	} else if actionName == "run" {

		// 为私有网络运行一个区块链节点
		// 如需多个节点连接：见文档 https://geth.ethereum.org/docs/interface/peer-to-peer

		// --nodiscover
		// To start geth without the discovery protocol,
		// you can use the --nodiscover parameter.
		// You only want this is you are running a test node or an experimental test network with fixed nodes.

		// --netrestrict
		// If Internet connectivity is not required or all member nodes connect using well-known IPs,
		// we strongly recommend setting up Geth to restrict peer-to-peer connectivity to an IP subnet.
		// Doing so will further isolate your network and prevents cross-connecting with other blockchain networks
		// in case your nodes are reachable from the Internet.
		// Use the --netrestrict flag to configure a whitelist of IP networks:

		node0Cmd := cmd.NewCmd(
			"geth",

			"--datadir", genesisDataDir,

			"--networkid", chainNetworkID,

			"--port", "30303",

			"--identity", "node-0",

			// "--mine",                 // 启动挖矿

			"--nodiscover",
			"--netrestrict", "10.0.0.0/8", // 局域网

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

		devutils.RunCmd(node0Cmd)
		fmt.Println("[dev-cmd] run node0 done")

	} else if actionName == "dev-run" {

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
		fmt.Println("[dev-cmd] run dev-run done")
	}
}

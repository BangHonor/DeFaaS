package main

// 自动构建 eth 私有网络

import (
	"flag"
	"fmt"
	"log"
	"path"

	"defaas/dev-cmd/utils"

	"github.com/go-cmd/cmd"
)

func init() {
	// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetFlags(log.Lmicroseconds | log.Llongfile)
}

const (
	chainNetworkID = "666666"
	chainDir       = "./private_chain"
	genesisDataDir = "data-0"
	httpRPCAddr    = "127.0.0.1"
)

// go run dev-cmd/private-chain/main.go -action=build
// go run dev-cmd/private-chain/main.go -action=run

func wrap(name string) string {
	return path.Join(chainDir, name)
}

func main() {

	var actionName string

	flag.StringVar(&actionName, "action", "run", "action")
	flag.Parse()

	if actionName == "build" {

		// 构建创世区块

		utils.RunCmd(cmd.NewCmd(
			"rm", "-rf", wrap(genesisDataDir)))

		genesisCmd := cmd.NewCmd(
			"geth",
			"--datadir", wrap(genesisDataDir),
			"init",
			wrap("genesis.json"))

		utils.RunCmd(genesisCmd)
		fmt.Println("[dev-cmd] build genesis done")

	} else if actionName == "run" {

		// 私有测试网络：运行一个区块链阶段
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
			"--datadir", wrap(genesisDataDir),
			"--networkid", chainNetworkID,
			"--port", "30303",
			"--identity", "node-0",
			"--nodiscover",
			"--netrestrict", "192.168.199.0/24",
			"--http",
			"--http.addr", "127.0.0.1",
			"--http.port", "8545",
			"--http.api", "eth,web3,miner,admin,personal,net",
			"--http.corsdomain", "*")

		utils.RunCmd(node0Cmd)
		fmt.Println("[dev-cmd] run node0 done")

	}
}

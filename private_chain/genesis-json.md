https://geth.ethereum.org/docs/interface/private-network

https://cloud.tencent.com/developer/article/1559634
https://blog.csdn.net/u010921136/article/details/107222259

├── config: 链配置，新区块的出块规则依赖链配置。
├────── chainId: 即区块链网络 id，用于连接其他节点，不同 id 网络的节点无法相互连接。以太坊公网 id 为 1。
├── alloc: 即分配，创世初始账户的资产配置，直接将账户资产写入区块中。
├── coinbase: 即「生成交易」中的 from。
├── difficulty: 挖矿难度系数，与出块速度成负相关。
├── extraData: 额外数据。
├── gasLimit: 十六进制，燃料上限。
├── nonce: 随机数。
├── parentHash: 母区块的哈希值。
└── timestamp： UTC 时间戳。



 personal.newAccount("123456") 
 personal.newAccount("000000") 


8df69a21a6bde726cf249b4f792784215a0436e7
5c731e6dc87f8eb96249dd54b926b46cf5cc5304


引导节点

"enr:-Je4QD6mtMMnyYrQF03WYp7uevy4Q9zIiN2UK-q1_I6ACRCHXXmmgO0Q9JHc_JefggKWTqIxOMr9jcamaz5YMe6IM5kGg2V0aMfGhLBf8zqAgmlkgnY0gmlwhKwSkiaJc2VjcDI1NmsxoQLj9tTvzq4V3McYnkRbW5r5IbUEIfilP3nlj6Np-KbRz4N0Y3CCdl-DdWRwgnZf"



初始化创世区块：
	geth --datadir chaindata init genesis.json
启动节点：
	geth --datadir chaindata \
	--port 30303 \
	--networkid 1108 \
	--identity "node1" \
	--rpc \
	--rpcport 9545 \
	--rpccorsdomain "*" \
	--allow-insecure-unlock \
	console
————————————————
版权声明：本文为CSDN博主「路飞的纯白世界」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/u010921136/article/details/107222259


 192.168.199.202/24 


搭建私有网络

- 创世区块

geth --datadir data-0 init genesis.json

- 在创世区块的基础上启动第一个节点

geth --datadir data-0 --networkid 666666 --port 30303


隔离网络：
Use the `--netrestrict` flag to configure a whitelist of IP networks

geth --datadir data-0 --networkid 666666 --port 30303 --netrestrict 192.168.199.0/24

启动 HTTP RPC 服务

API AND CONSOLE OPTIONS:
  --ipcdisable                        Disable the IPC-RPC server
  --ipcpath value                     Filename for IPC socket/pipe within the datadir (explicit paths escape it)
  --http                              Enable the HTTP-RPC server
  --http.addr value                   HTTP-RPC server listening interface (default: "localhost")
  --http.port value                   HTTP-RPC server listening port (default: 8545)
  --http.api value                    API's offered over the HTTP-RPC interface
  --http.corsdomain value             Comma separated list of domains from which to accept cross origin requests (browser enforced)
  --http.vhosts value                 Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts '*' wildcard. (default: "localhost")

geth \
--datadir data-0 --networkid 666666 --port 30303 \
--identity "node-0" \
--nodiscover \
--netrestrict 192.168.199.0/24 \
--http --http.addr "127.0.0.1" --http.port "8545" --http.api "eth,web3,miner,admin,personal,net" --http.corsdomain "*"


- 启动第二个阶段

geth --datadir data-1 --networkid 666666 --port 30305  











func init() {
	// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetFlags(log.Lmicroseconds | log.Llongfile)
}

/*
https://infura.io 为开发人员提供了一种不必运行全节点就可以连接以太坊网络的方法。

https://mainnet.infura.io/v3/1d69d7036ac046af9e31ff2a789d74c0
https://ropsten.infura.io/v3/1d69d7036ac046af9e31ff2a789d74c0
https://kovan.infura.io/v3/1d69d7036ac046af9e31ff2a789d74c0
https://rinkeby.infura.io/v3/1d69d7036ac046af9e31ff2a789d74c0
https://goerli.infura.io/v3/1d69d7036ac046af9e31ff2a789d74c0

*/
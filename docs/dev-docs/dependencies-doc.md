# dependencies doc

## version

### go

1.15.8


### solc

Version: 0.8.1+commit.df193b15.Linux.g++


### go-ethereum

github.com/ethereum/go-ethereum v1.10.1

### geth

Version: 1.9.25-stable
Git Commit: e7872729012a4871397307b12cc3f4772ffcbec6
Git Commit Date: 20201211
Architecture: amd64
Protocol Versions: [65 64 63]
Go Version: go1.16
Operating System: linux

--------------------

## document

### solidity

- 官方文档 https://docs.soliditylang.org/en/latest/

    > 如果之前你使用的是低版本，需要关注 v0.8.0 的破坏性变更 https://docs.soliditylang.org/en/latest/080-breaking-changes.html#solidity-v0-8-0-breaking-changes


### go-ethereum

- 用Go来做以太坊开发 https://goethereumbook.org/zh/

### geth

- 基本实操 https://www.jianshu.com/p/017da2910795

- geth 命令行参数说明 https://geth.ethereum.org/docs/interface/command-line-options

    需要关注的参数有:

    - API 参数

        ```
        --http                              Enable the HTTP-RPC server
        --http.addr value                   HTTP-RPC server listening interface (default: "localhost")
        --http.port value                   HTTP-RPC server listening port (default: 8545)
        --http.api value                    API's offered over the HTTP-RPC interface
        --http.corsdomain value             Comma separated list of domains from which to accept cross origin requests (browser enforced)
        --http.vhosts value                 Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts '*' wildcard. (default: "localhost")

        --ws                                Enable the WS-RPC server
        --ws.addr value                     WS-RPC server listening interface (default: "localhost")
        --ws.port value                     WS-RPC server listening port (default: 8546)
        --ws.api value                      API's offered over the WS-RPC interface
        --ws.origins value                  Origins from which to accept websockets requests
        ```

    - 区块链网络参数

        ```
        NETWORKING OPTIONS:
        --bootnodes value                   Comma separated enode URLs for P2P discovery bootstrap
        --discovery.dns value               Sets DNS discovery entry points (use "" to disable DNS)
        --port value                        Network listening port (default: 30303)
        --maxpeers value                    Maximum number of network peers (network disabled if set to 0) (default: 50)
        --maxpendpeers value                Maximum number of pending connection attempts (defaults used if set to 0) (default: 0)
        --nat value                         NAT port mapping mechanism (any|none|upnp|pmp|extip:<IP>) (default: "any")
        --nodiscover                        Disables the peer discovery mechanism (manual peer addition)
        --v5disc                            Enables the experimental RLPx V5 (Topic Discovery) mechanism
        --netrestrict value                 Restricts network communication to the given IP networks (CIDR masks)
        --nodekey value                     P2P node key file
        --nodekeyhex value                  P2P node key as hex (for testing)
        ```


- 账户管理 https://geth.ethereum.org/docs/interface/managing-your-accounts

- eth 开发模式 https://geth.ethereum.org/docs/getting-started/dev-mode



## docker 


- 官方文档 https://docs.docker.com/

- 实践文档 https://yeasy.gitbook.io/docker_practice/


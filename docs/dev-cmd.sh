ps -a | grep geth
ps -a | grep go
geth attach --datadir ./private-chain/data-0
go run dev-cmd/private-chain/main.go -action=run
go run dev-cmd/deploy/main.go
go run dev-cmd/gen/main.go -name=Market

personal.newAccount()
eth.blockNumber
eth.accounts
miner.start(1)
miner.stop()

git push origin main
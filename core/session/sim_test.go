package session

import (
	"defaas/contracts/go/faastoken"
	"defaas/contracts/go/market"
	"defaas/contracts/go/witnesspool"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	singleSimSuiteOnce sync.Once
	singleSimSuite     *simSuite
)

type simSuite struct {
	auth         *bind.TransactOpts
	blockchain   *backends.SimulatedBackend
	faastoken    *FaaSTokenSession
	market       *MarketSeesion
	witnesspool  *WitnessPoolSession
	ownAddress   common.Address
	otherAddress common.Address
}

// 获得一个模拟套件的单例
func getSingleSimSuite() *simSuite {

	singleSimSuiteOnce.Do(func() {

		auth, blockchain := newSimBlockchain()

		faastokenAddress, _, _, _ := faastoken.DeployFaaSToken(auth, blockchain)
		blockchain.Commit()
		faastoken, _ := NewFaaSTokenSeesion(blockchain, faastokenAddress, auth)

		marketAddress, _, _, _ := market.DeployMarket(auth, blockchain, faastokenAddress)
		blockchain.Commit()
		market, _ := NewMarketSeesion(blockchain, marketAddress, auth)

		witnesspoolAddress, _ := market.WpContract()
		witnesspool, _ := NewWitnessPoolSession(blockchain, witnesspoolAddress, auth)

		singleSimSuite = &simSuite{
			auth:         auth,
			blockchain:   blockchain,
			faastoken:    faastoken,
			market:       market,
			witnesspool:  witnesspool,
			ownAddress:   auth.From,
			otherAddress: common.HexToAddress("0x01"),
		}
	})

	return singleSimSuite
}

// 获得一个模拟区块链的新实例
func newSimBlockchain() (*bind.TransactOpts, *backends.SimulatedBackend) {

	// Generate a new random account and a funded simulator
	key, _ := crypto.GenerateKey()
	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 1<<32)

	return auth, blockchain
}

// 获得一个 FaaSToken 合约新实例
func newSimFaaSToken() (*bind.TransactOpts, *backends.SimulatedBackend, common.Address, *faastoken.FaaSTokenSession) {

	auth, blockchain := newSimBlockchain()

	address, _, _, _ := faastoken.DeployFaaSToken(auth, blockchain)
	blockchain.Commit()
	session, _ := NewFaaSTokenSeesion(blockchain, address, auth)

	return auth, blockchain, address, session
}

// 获得一个 Market 合约新实例
func newSimMarket() (*bind.TransactOpts, *backends.SimulatedBackend, common.Address, *market.MarketSession) {

	auth, blockchain, tokenAddress, _ := newSimFaaSToken()

	address, _, _, _ := market.DeployMarket(auth, blockchain, tokenAddress)
	blockchain.Commit()
	session, _ := NewMarketSeesion(blockchain, address, auth)

	return auth, blockchain, address, session
}

// 获得一个 WitnessPool 合约新实例
func newSimWitnessPool() (*bind.TransactOpts, *backends.SimulatedBackend, common.Address, *witnesspool.WitnessPoolSession) {

	auth, blockchain, _, marketSession := newSimMarket()

	address, _ := marketSession.WpContract()
	session, _ := NewWitnessPoolSession(blockchain, address, auth)

	return auth, blockchain, address, session
}

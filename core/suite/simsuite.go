package suite

import (
	"defaas/contracts/go/faastoken"
	"defaas/contracts/go/market"
	"log"
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
	singleSimSuite     *SimSuite
)

type SimSuite struct {
	Auth         *bind.TransactOpts
	Blockchain   *backends.SimulatedBackend
	FaasToken    *FaaSTokenSession
	Market       *MarketSeesion
	WitnessPool  *WitnessPoolSession
	OwnAddress   common.Address
	OtherAddress common.Address
}

// 获得模拟套件的单例
func SingleSimSuite() *SimSuite {

	singleSimSuiteOnce.Do(func() {

		auth, blockchain := NewSimBlockchain()

		faastokenAddress, _, _, err := faastoken.DeployFaaSToken(auth, blockchain)
		if err != nil {
			log.Fatal(err)
		}
		blockchain.Commit()

		marketAddress, _, _, err := market.DeployMarket(auth, blockchain, faastokenAddress)
		if err != nil {
			log.Fatal(err)
		}
		blockchain.Commit()

		marketSession, err := NewMarketSeesion(blockchain, marketAddress, auth)
		if err != nil {
			log.Fatal(err)
		}
		witnesspoolAddress, _ := marketSession.WpContract()

		faastokenSession, _ := NewFaaSTokenSeesion(blockchain, faastokenAddress, auth)
		witnesspoolSession, _ := NewWitnessPoolSession(blockchain, witnesspoolAddress, auth)

		singleSimSuite = &SimSuite{
			Auth:         auth,
			Blockchain:   blockchain,
			FaasToken:    faastokenSession,
			Market:       marketSession,
			WitnessPool:  witnesspoolSession,
			OwnAddress:   auth.From,
			OtherAddress: common.HexToAddress("0x01"),
		}
	})

	return singleSimSuite
}

// 获得一个模拟区块链的新实例
func NewSimBlockchain() (*bind.TransactOpts, *backends.SimulatedBackend) {

	// Generate a new random account and a funded simulator
	key, _ := crypto.GenerateKey()
	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(100000000000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 1<<32)
	return auth, blockchain
}

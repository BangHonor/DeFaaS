package suite

import (
	"defaas/contracts/go/faastoken"
	"defaas/contracts/go/market"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

type SimSuite struct {
	Suite
	Blockchain *backends.SimulatedBackend
	Owner      *bind.TransactOpts
	Others     []*bind.TransactOpts
}

// 获得一个模拟套件的新实例
func NewSimSuite() *SimSuite {

	ownerAuth, blockchain, othersAuth := NewSimBlockchain()

	faastokenAddress, _, _, err := faastoken.DeployFaaSToken(ownerAuth, blockchain)
	if err != nil {
		log.Fatal(err)
	}
	blockchain.Commit()

	marketAddress, _, _, err := market.DeployMarket(ownerAuth, blockchain, faastokenAddress)
	if err != nil {
		log.Fatal(err)
	}
	blockchain.Commit()

	marketSession, err := NewMarketSeesion(blockchain, marketAddress, ownerAuth)
	if err != nil {
		log.Fatal(err)
	}
	witnesspoolAddress, _ := marketSession.WpContract()

	suite, _ := NewSuite(blockchain, ownerAuth, faastokenAddress, marketAddress, witnesspoolAddress)

	return &SimSuite{
		Blockchain: blockchain,
		Owner:      ownerAuth,
		Others:     othersAuth,
		Suite:      *suite,
	}
}

// 获得一个模拟区块链的新实例
func NewSimBlockchain() (*bind.TransactOpts, *backends.SimulatedBackend, []*bind.TransactOpts) {

	chainID := big.NewInt(1337)

	// Generate a new random account and a funded simulator
	ownerKey, _ := crypto.GenerateKey()
	ownerAuth, _ := bind.NewKeyedTransactorWithChainID(ownerKey, chainID)
	alloc := make(core.GenesisAlloc)
	alloc[ownerAuth.From] = core.GenesisAccount{Balance: big.NewInt(100000000000000000)}

	othersAuth := []*bind.TransactOpts{}
	for i := 0; i < 10; i++ {
		key, _ := crypto.GenerateKey()
		auth, _ := bind.NewKeyedTransactorWithChainID(key, chainID)
		alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(100000000000000000)}

		othersAuth = append(othersAuth, auth)
	}

	blockchain := backends.NewSimulatedBackend(alloc, 1<<32)
	return ownerAuth, blockchain, othersAuth
}

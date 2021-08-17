package suite

import (
	"defaas/contracts/go/faastoken"
	"defaas/contracts/go/market"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core"
)

type SimSuite struct {
	Suite
	Blockchain *backends.SimulatedBackend
	Keys       []*keystore.Key
}

// 获得一个模拟环境的新实例
func NewSim() *SimSuite {

	blockchain, keys := NewSimBlockchain()

	ownerAuth, err := bind.NewKeyedTransactorWithChainID(keys[0].PrivateKey, blockchain.Blockchain().Config().ChainID)
	if err != nil {
		log.Fatal(err)
	}

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
	witnesspoolAddress, err := marketSession.WpContract()
	if err != nil {
		log.Fatal(err)
	}

	suite, err := NewSuite(blockchain, ownerAuth, faastokenAddress, marketAddress, witnesspoolAddress)
	if err != nil {
		log.Fatal(err)
	}

	return &SimSuite{
		Blockchain: blockchain,
		Keys:       keys,
		Suite:      *suite,
	}
}

// 获得一个模拟区块链的新实例
func NewSimBlockchain() (*backends.SimulatedBackend, []*keystore.Key) {

	tempDir, err := ioutil.TempDir("", "ks")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	ks := keystore.NewKeyStore(tempDir, keystore.LightScryptN, keystore.LightScryptP)
	password := ""
	keys := []*keystore.Key{}
	numKey := 10

	for i := 0; i < numKey; i++ {
		a, err := ks.NewAccount(password)
		if err != nil {
			log.Fatal(err)
		}
		j, err := ks.Export(a, password, password)
		if err != nil {
			log.Fatal(err)
		}
		key, err := keystore.DecryptKey(j, password)
		if err != nil {
			log.Fatal(err)
		}

		keys = append(keys, key)
	}

	alloc := make(core.GenesisAlloc)

	for i := 0; i < len(keys); i++ {
		alloc[keys[i].Address] = core.GenesisAccount{Balance: big.NewInt(100000000000000000)}
	}

	blockchain := backends.NewSimulatedBackend(alloc, 1<<32)

	return blockchain, keys
}

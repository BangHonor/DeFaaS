package session

import (
	"defaas/contracts/go/witnesspool"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
)

func getSimWitnessPool() (*bind.TransactOpts, *backends.SimulatedBackend, common.Address, *witnesspool.WitnessPoolSession) {

	auth, blockchain, _, marketSession := getSimMarket()

	address, _ := marketSession.WpContract()
	session, _ := NewWitnessPoolSeesion(blockchain, address, auth)

	return auth, blockchain, address, session
}

func TestStdWitnessDepoist(t *testing.T) {

	auth, blockchain, _, session := getSimWitnessPool()
	_ = auth
	_ = blockchain
	_ = session

	deposit, err := session.StdWitnessDepoist()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("deposit", deposit)
}

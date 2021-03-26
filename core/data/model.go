package data

import (
	"errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type FaaSLevel struct {
	ID   int
	Core *big.Int
	Mem  *big.Int
}

type DeploymentOrderState int

const (
	InitState       DeploymentOrderState = 0
	BiddingState    DeploymentOrderState = 1
	ConfirmingState DeploymentOrderState = 2
	DeployingState  DeploymentOrderState = 3
	FulfillingState DeploymentOrderState = 4
	FinishedState   DeploymentOrderState = 5
)

var (
	ErrWrongState = errors.New("wrong state")
)

type DeploymentOrder struct {
	Customer         common.Address
	Nonce            *big.Int
	FaaSLevelID      *big.Int
	HighestUnitPrice *big.Int
	FaaSDuration     *big.Int
	BiddingDuration  *big.Int
	PublicKey        string
	FulfillSecretKey [32]byte
	FulfillKey       [32]byte // FulfillKey = sha256.Sum256(FulfillSecretKey)
}

type DeploymentInfo struct {
	Provider   common.Address
	UnitPrice  *big.Int
	FuncPath   string
	DeployPath string
	AccessKey  string
}

type Lease struct {
	IsProviderViolated bool
}

type DeploymentItem struct {
	sync.Mutex

	ID    *big.Int
	State DeploymentOrderState
	Order DeploymentOrder
	Info  DeploymentInfo
}

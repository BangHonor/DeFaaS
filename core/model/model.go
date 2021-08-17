package data

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type FaaSLevel struct {
	ID   int
	Core *big.Int
	Mem  *big.Int
}

type DeploymentOrder struct {
	Customer         common.Address
	Nonce            *big.Int
	FaaSLevelID      *big.Int
	HighestUnitPrice *big.Int
	FaaSDuration     *big.Int
	BiddingDuration  *big.Int
}

type DeploymentInfo struct {
	Provider   common.Address
	UnitPrice  *big.Int
	FuncPath   string
	DeployPath string
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

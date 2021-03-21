package data

import (
	"math/big"

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
	PublicKey        string
}

type DeploymentInfo struct {
	Provider   common.Address
	FuncPath   string
	DeployPath string
	AccessKey  string
}

type Lease struct {
	Customer           common.Address
	Provider           common.Address
	IsProviderViolated bool
}

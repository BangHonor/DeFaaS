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
	FaaSLevelID      *big.Int
	HighestUnitPrice *big.Int
	FaaSDuration     *big.Int
	BiddingDuration  *big.Int
}

type DeploymentInfo struct {
	EndAddr          string
	DeployServerAddr string
	AccessSecreKey   string
}

type Lease struct {
	Customer           common.Address
	Provider           common.Address
	IsProviderViolated bool
}

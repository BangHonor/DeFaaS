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

type DeploymentOrderState int

const (
	Bidding    DeploymentOrderState = 1
	Confirming DeploymentOrderState = 2
	Deploying  DeploymentOrderState = 3
	Fulfilling DeploymentOrderState = 4
	Finished   DeploymentOrderState = 5
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
	FuncPath   string
	DeployPath string
	AccessKey  string
}

type Lease struct {
	Customer           common.Address
	Provider           common.Address
	IsProviderViolated bool
}

package data

import (
	"errors"
)

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

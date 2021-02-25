package provider

import (
	"defaas/contracts/go/witnesspool"

	sdkClient "github.com/FISCO-BCOS/go-sdk/client"
)

type Clinet struct {
	_client *sdkClient.Client
	wp      witnesspool.WitnessPoolSession
}

func NewClient() *Clinet {
	client := &Clinet{}
	return client
}

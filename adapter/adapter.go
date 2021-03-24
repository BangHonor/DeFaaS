package adapter

import (
	"defaas/core/data"
	"log"
)

// Adapter is an interface of the off-chain deployment action.
type Adapter interface {
	// DeployTo is used by `CustomerClient`
	DeployTo(item *data.DeploymentItem, adapterData interface{}) error
	// DeployFrom is used by `ProviderClient`
	DeployFrom(item *data.DeploymentItem, adapterData interface{}) error
}

func New(adapterName string) Adapter {

	if adapterName == "docker" {

		return NewDockerAdapter()
	}

	log.Fatal("unknown adpater")
	return nil
}

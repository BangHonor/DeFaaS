package adapter

import (
	"defaas/core/data"
	"log"
)

// data := &XXXAdapterData{}
// adpater.By("XXX").Encode(data)
// adapter.By("XXX").Decode(encoded)
// adapter.By("XXX").Deploy(data)

func By(name string) Adapter {

	if name == "raw-go" {
		return nil
	}

	// if name == "docker" {
	// 	return NewDockerAdapter()
	// }

	// if name == "openfaas" {
	// 	return NewOpenFaaSAdapter()
	// }

	log.Fatal("unkown adapter")

	return nil
}

type Adapter interface {
	Encode() ([]byte, error)
	Decode(encoded []byte) error
	Deploy(item *data.DeploymentItem) error
}

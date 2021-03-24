package adapter

import "defaas/core/data"

// OpenFaaSAdapter is the deploymentt action adapter for OpenFaaS.
// OpenFaaSAdapter implements `Adapter` interface.
type OpenFaaSAdapter struct{}

func (ofa *OpenFaaSAdapter) DeployTo(item *data.DeploymentItem, adapterData interface{}) error {
	// TODO
	return nil
}

func (ofa *OpenFaaSAdapter) DeployFrom(item *data.DeploymentItem, adapterData interface{}) error {
	// TODO
	return nil
}

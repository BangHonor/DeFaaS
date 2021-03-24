package adapter

import "defaas/core/data"

// DockerAdapter is the deploymentt action adapter for Docker.
// DockerAdapter implements `Adapter` interface.
type DockerAdapter struct{}

func NewDockerAdapter() *DockerAdapter {
	ds := &DockerAdapter{}
	return ds
}

func (da *DockerAdapter) DeployTo(item *data.DeploymentItem, adapterData interface{}) error {
	// TODO
	return nil
}

func (da *DockerAdapter) DeployFrom(item *data.DeploymentItem, adapterData interface{}) error {
	// TODO
	return nil
}

package adapter

import "defaas/core/data"

// DockerAdapter is the deploymentt action adapter for Docker.
// DockerAdapter implements `Adapter` interface.
type DockerAdapter struct{}

func NewDockerAdapter() *DockerAdapter {
	ds := &DockerAdapter{}
	return ds
}

type DockerAdapterDeployToData struct {
	ImageURL  string
	ImageHash string
}

type DockerAdapterDeployFromData struct {
	//
}

func (da *DockerAdapter) DeployTo(item *data.DeploymentItem, adapterData interface{}) error {

	data := adapterData.(DockerAdapterDeployToData)
	_ = data

	return nil
}

func (da *DockerAdapter) DeployFrom(item *data.DeploymentItem, adapterData interface{}) error {

	data := adapterData.(DockerAdapterDeployFromData)
	_ = data

	return nil
}

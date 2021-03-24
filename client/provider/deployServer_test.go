package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDeployPath(t *testing.T) {

	assert := assert.New(t)

	adapter := "docker"
	serverAddr := "127.0.0.1:60666"
	serverEntry := "deploy"

	deployPath := "docker|127.0.0.1:60666|deploy"

	assert.Equal(deployPath, GetDeployPath(adapter, serverAddr, serverEntry))

	_adapter, _serverAddr, _serverEntry := ParseDeployPath(GetDeployPath(adapter, serverAddr, serverEntry))
	assert.Equal(adapter, _adapter)
	assert.Equal(serverAddr, _serverAddr)
	assert.Equal(serverEntry, _serverEntry)
}

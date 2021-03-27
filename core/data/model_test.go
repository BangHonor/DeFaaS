package data

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

func TestServerAddrTrim(t *testing.T) {

	assert := assert.New(t)

	serverAddrs := []string{
		"http://127.0.0.1:60666",
		"https://127.0.0.1",
	}

	assert.Equal(serverAddrs[0][7:], ServerAddrTrim(serverAddrs[0]))
	assert.Equal(serverAddrs[1][8:], ServerAddrTrim(serverAddrs[1]))
}

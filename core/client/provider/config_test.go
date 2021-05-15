package provider

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {

	assert := assert.New(t)

	var testConfig1 = []byte(`
	[deploypath]
	adapter = "docker"
	serverAddr = "127.0.0.1:60666"
	serverEntry = "deploy"
	`)

	pc, err := ParseConfig(bytes.NewBuffer(testConfig1))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(pc.Adapter, "docker")
	assert.Equal(pc.ServerAddr, "127.0.0.1:60666")
	assert.Equal(pc.ServerEntry, "deploy")
}

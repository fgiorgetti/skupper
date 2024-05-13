package common

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestLoadClientSecret(t *testing.T) {
	r := FileSystemConfigurationRenderer{
		OutputPath: "/home/fgiorget/.local/share/skupper/sites/fedora-west",
	}
	secret, err := r.loadClientSecret("client-go-west")
	assert.Assert(t, err)
	assert.Assert(t, secret != nil)
	secretJson, _ := json.MarshalIndent(secret, "", "  ")
	fmt.Println(string(secretJson))
}

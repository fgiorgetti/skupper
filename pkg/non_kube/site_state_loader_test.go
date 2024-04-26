package non_kube

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestFileSystemSiteStateLoder(t *testing.T) {
	fsStateLoader := &FileSystemSiteStateLoader{
		Path: "/home/fgiorget/git/skupper-v2/cmd/controller-podman-v2/crs/west",
	}
	siteState, err := fsStateLoader.Load()
	assert.Assert(t, err)
	assert.Assert(t, siteState != nil)
	assert.Assert(t, siteState.Site.Name != "")
	assert.Equal(t, len(siteState.Listeners), 3)
	siteStateJson, _ := json.MarshalIndent(siteState, "", "  ")
	fmt.Println(string(siteStateJson))
}

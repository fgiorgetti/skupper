package common

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestFileSystemSiteStateLoder(t *testing.T) {
	fsStateLoader := &FileSystemSiteStateLoader{
		//Path: "/home/fgiorget/git/skupper-v2/api/types/crds",
		Path: "/home/fgiorget/git/skupper-v2/cmd/bootstrap/crs/west",
	}
	siteState, err := fsStateLoader.Load()
	assert.Assert(t, err)
	assert.Assert(t, siteState != nil)
	assert.Assert(t, siteState.Site.Name != "")
	//assert.Equal(t, len(siteState.Listeners), 2)
	assert.Equal(t, len(siteState.Listeners), 1)
	siteStateJson, _ := json.MarshalIndent(siteState, "", "  ")
	fmt.Println(string(siteStateJson))
}

package netconvert

import (
	"testing"

	"github.com/EricNeid/go-netconvert/internal/util"
)

func TestDecode(t *testing.T) {
	// action
	net, err := Decode("testdata/sample.osm.xml")
	// verify
	util.Ok(t, err)
	util.Assert(t, len(net.Nodes) > 0, "Empty list of nodes returned")
	util.Assert(t, len(net.Ways) > 0, "Empty list of ways returned")

	way := net.Ways[0]
	util.Assert(t, len(way.NodeRefs) > 0, "Way has no node references")
	util.Assert(t, len(way.NodeRefs[0].NodeID) > 0, "Node reference is empty")

	util.Assert(t, len(way.Tags) > 0, "Way has no tags")
	util.Assert(t, len(way.Tags[0].Name) > 0, "Tag has empty name")
	util.Assert(t, len(way.Tags[0].Value) > 0, "Tag has empty value")
}

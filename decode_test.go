package netconvert

import (
	"testing"

	"github.com/EricNeid/go-netconvert/internal/verify"
)

func TestDecode(t *testing.T) {
	// action
	net, err := Decode("testdata/sample.osm.xml")
	// verify
	verify.Ok(t, err)
	verify.Assert(t, len(net.Nodes) > 0, "Empty list of nodes returned")
	verify.Assert(t, len(net.Ways) > 0, "Empty list of ways returned")

	way := net.Ways[0]
	verify.Assert(t, len(way.NodeRefs) > 0, "Way has no node references")
	verify.Assert(t, len(way.NodeRefs[0].NodeID) > 0, "Node reference is empty")

	verify.Assert(t, len(way.Tags) > 0, "Way has no tags")
	verify.Assert(t, len(way.Tags[0].Name) > 0, "Tag has empty name")
	verify.Assert(t, len(way.Tags[0].Value) > 0, "Tag has empty value")
}

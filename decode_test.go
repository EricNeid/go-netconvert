package netconvert

import (
	"testing"

	"github.com/EricNeid/go-netconvert/internal/test"
)

func TestDecode(t *testing.T) {
	// action
	net, err := Decode("testdata/sample.osm.xml")
	// verify
	test.Ok(t, err)
	test.Assert(t, len(net.Nodes) > 0, "Empty list of nodes returned")
	test.Assert(t, len(net.Ways) > 0, "Empty list of ways returned")

	way := net.Ways[0]
	test.Assert(t, len(way.NodeRefs) > 0, "Way has no node references")
	test.Assert(t, len(way.NodeRefs[0].NodeID) > 0, "Node reference is empty")

	test.Assert(t, len(way.Tags) > 0, "Way has no tags")
	test.Assert(t, len(way.Tags[0].Name) > 0, "Tag has empty name")
	test.Assert(t, len(way.Tags[0].Value) > 0, "Tag has empty value")
}

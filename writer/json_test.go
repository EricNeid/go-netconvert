package writer

import (
	"testing"

	"github.com/EricNeid/go-netconvert/internal/verify"
	"github.com/EricNeid/go-netconvert/osm"
)

func TestNodesAsJSON(t *testing.T) {
	// arrange
	testData := []osm.Node{
		{
			Tags: []osm.Tag{
				{
					Name: "foo",
				},
			},
		},
		{
			Tags: []osm.Tag{
				{
					Name: "bar",
				},
			},
		},
	}

	// action
	err := NodesAsJSON(testData, "../testdata/nodes.json")

	// verify
	verify.Ok(t, err)
}

func TestWaysAsJSON(t *testing.T) {
	// arrange
	testData := []osm.Way{
		{
			Tags: []osm.Tag{
				{
					Name: "foo",
				},
			},
		},
		{
			Tags: []osm.Tag{
				{
					Name: "bar",
				},
			},
		},
	}

	// action
	err := WaysAsJSON(testData, "../testdata/ways.json")

	// verify
	verify.Ok(t, err)
}

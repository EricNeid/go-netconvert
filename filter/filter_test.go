package filter

import (
	"testing"

	"github.com/EricNeid/go-netconvert/internal/test"
	"github.com/EricNeid/go-netconvert/osm"
)

func TestNodes(t *testing.T) {
	// arrange
	testData := []osm.Node{
		osm.Node{
			Tags: []osm.Tag{
				osm.Tag{
					Name: "foo",
				},
			},
		},
		osm.Node{
			Tags: []osm.Tag{
				osm.Tag{
					Name: "bar",
				},
			},
		},
	}

	// action
	result := Nodes(testData, func(node osm.Node) bool {
		for _, t := range node.Tags {
			if t.Name == "foo" {
				return true
			}
		}
		return false
	})

	// verify
	test.Equals(t, 1, len(result))
}

func TestWays(t *testing.T) {
	// arrange
	testData := []osm.Way{
		osm.Way{
			Tags: []osm.Tag{
				osm.Tag{
					Name: "foo",
				},
			},
		},
		osm.Way{
			Tags: []osm.Tag{
				osm.Tag{
					Name: "bar",
				},
			},
		},
	}

	// action
	result := Ways(testData, func(node osm.Way) bool {
		for _, t := range node.Tags {
			if t.Name == "foo" {
				return true
			}
		}
		return false
	})

	// verify
	test.Equals(t, 1, len(result))
}

func TestToFilter(t *testing.T) {
	// arrange
	testData := ""

	// action
	result := ToFilter(testData)

	// verify
	test.Equals(t, 0, len(result))
}

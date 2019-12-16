package filter

import (
	"testing"

	"github.com/EricNeid/go-netconvert/internal/verify"
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
	verify.Equals(t, 1, len(result))
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
	verify.Equals(t, 1, len(result))
}

func TestToFilter(t *testing.T) {
	// arrange
	testData := "a=2,b<3,c>4,d"

	// action
	result, err := ToFilter(testData)

	// verify
	verify.Ok(t, err)
	verify.Equals(t, 4, len(result))

	filter0 := result[0]
	verify.Equals(t, "a", filter0.Name)
	verify.Equals(t, "2", filter0.Value)
	verify.Assert(t, EQ == filter0.Operand, "Expected operand to be eq")

	filter1 := result[1]
	verify.Equals(t, "b", filter1.Name)
	verify.Equals(t, "3", filter1.Value)
	verify.Assert(t, LT == filter1.Operand, "Expected operand to be lt")

	filter2 := result[2]
	verify.Equals(t, "c", filter2.Name)
	verify.Equals(t, "4", filter2.Value)
	verify.Assert(t, GT == filter2.Operand, "Expected operand to be gt")

	filter3 := result[3]
	verify.Equals(t, "d", filter3.Name)
	verify.Assert(t, NOP == filter3.Operand, "Expected operand to be nop")
}

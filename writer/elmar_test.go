package writer

import (
	"testing"

	"github.com/EricNeid/go-netconvert/internal/verify"
	"github.com/EricNeid/go-netconvert/osm"
)

func Test_toElmarWays(t *testing.T) {
	// arrnage
	testData := osm.Net{
		Nodes: []osm.Node{
			osm.Node{ID: 1},
			osm.Node{ID: 2},
			osm.Node{ID: 3},
		},
		Ways: []osm.Way{
			osm.Way{
				ID:       42,
				NodeRefs: []osm.NodeRef{osm.NodeRef{NodeID: 1}, osm.NodeRef{NodeID: 2}, osm.NodeRef{NodeID: 3}},
			},
		},
	}
	// action
	result := toElmarWays(&testData)
	// verify
	verify.Equals(t, 1, len(result))
	verify.Equals(t, 2, len(result[0].edges))

	verify.Equals(t, int64(1), result[0].edges[0].from.ID)
	verify.Equals(t, int64(2), result[0].edges[0].to.ID)
	verify.Equals(t, int64(2), result[0].edges[1].from.ID)
	verify.Equals(t, int64(3), result[0].edges[1].to.ID)
}

func TestGetNames_shouldNotContainDuplicates(t *testing.T) {
	// arrange
	testData := []osm.Way{
		osm.Way{
			Tags: []osm.Tag{
				osm.Tag{
					Name:  "name",
					Value: "street",
				},
			},
		},
		osm.Way{
			Tags: []osm.Tag{
				osm.Tag{
					Name:  "name",
					Value: "street",
				},
			},
		},
	}
	// action
	resultNames := getNames(testData)
	// verify
	verify.Equals(t, 1, len(resultNames))
}

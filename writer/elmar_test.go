package writer

import (
	"testing"

	"github.com/EricNeid/go-netconvert/internal/verify"
	"github.com/EricNeid/go-netconvert/osm"
)

func Test_toElmarNames(t *testing.T) {
	// arrange
	testData := []osm.Way{
		{},
	}
	// action
	result := toElmarNames(testData)
	// verify
	verify.Assert(t, result != nil, "Parsed names are nil")
}

func Test_toElmarWays(t *testing.T) {
	// arrnage
	testData := osm.Net{
		Nodes: []osm.Node{
			{ID: 1},
			{ID: 2},
			{ID: 3},
		},
		Ways: []osm.Way{
			{
				ID:       42,
				NodeRefs: []osm.NodeRef{{NodeID: 1}, {NodeID: 2}, {NodeID: 3}},
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

func Test_toElmarLinks(t *testing.T) {
	// arrange
	testData := elmarWay{
		way: osm.Way{
			ID: 1,
		},
		edges: []nodeTupel{
			{
				from: osm.Node{ID: 2},
				to:   osm.Node{ID: 3},
			},
		},
	}
	// action
	result := toElmarLinks(testData)
	// verify
	verify.Equals(t, "1_2_3", result[0].linkID)
	verify.Equals(t, int64(2), result[0].nodeIDFrom)
	verify.Equals(t, int64(3), result[0].nodeIDTo)
}

func Test_writeWaysAsElmarFormat(t *testing.T) {
	// arrange
	testData := []elmarWay{
		{
			way: osm.Way{
				ID: 1,
			},
			edges: []nodeTupel{
				{
					from: osm.Node{ID: 2},
					to:   osm.Node{ID: 3},
				},
			},
		},
	}
	// action
	result := writeWaysAsElmarFormat(testData, "../testdata/test_writeWaysAsElmarFormat.txt")
	// verify
	verify.Ok(t, result)
}

func Test_writeNodesAsElmarFormat(t *testing.T) {
	// arrange
	testData := []elmarNode{
		{
			nodeID: 1,
			coordinates: []floatTupel{
				{x: 1, y: 2},
				{x: 4, y: 3},
			},
		},
	}
	// action
	result := writeNodesAsElmarFormat(testData, "../testdata/test_writeNodessAsElmarFormat.txt")
	// verify
	verify.Ok(t, result)
}

func Test_toElmarNodes(t *testing.T) {
	// arrange
	testData := osm.Net{
		Nodes: []osm.Node{
			{ID: 1, Lat: 1, Lon: 2},
			{ID: 1, Lat: 3, Lon: 4},
			{ID: 2, Lat: 1, Lon: 2},
		},
	}
	// action
	result := toElmarNodes(&testData)
	// verify
	verify.Equals(t, 2, len(result))
	verify.Equals(t, 2, len(result[0].coordinates))
	verify.Equals(t, float32(2), result[0].coordinates[0].x)
	verify.Equals(t, float32(1), result[0].coordinates[0].y)
	verify.Equals(t, float32(4), result[0].coordinates[1].x)
	verify.Equals(t, float32(3), result[0].coordinates[1].y)
	verify.Equals(t, 1, len(result[1].coordinates))
	verify.Equals(t, float32(2), result[1].coordinates[0].x)
	verify.Equals(t, float32(1), result[1].coordinates[0].y)
}

func Test_getNames_shouldNotContainDuplicates(t *testing.T) {
	// arrange
	testData := []osm.Way{
		{
			Tags: []osm.Tag{
				{
					Name:  "name",
					Value: "street",
				},
			},
		},
		{
			Tags: []osm.Tag{
				{
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

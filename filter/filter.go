package filter

import (
	"github.com/EricNeid/go-netconvert/osm"
)

// ConditionNode defines a filter which is applied
// to each node during the filtering process.
type ConditionNode = func(node osm.Node) bool

// ConditionWay defines a filter which is applied
// to each way during the filtering process.
type ConditionWay = func(way osm.Way) bool

// Nodes applies the given condition on all nodes and
// returns slice of nodes, satisfying the condition.
func Nodes(nodes []osm.Node, filter ConditionNode) []osm.Node {
	var result []osm.Node
	for _, n := range nodes {
		if filter(n) {
			result = append(result, n)
		}
	}
	return result
}

// Ways applies the given condition on all ways and
// returns slice of ways, satisfying the condition.
func Ways(ways []osm.Way, filter ConditionWay) []osm.Way {
	var result []osm.Way
	for _, w := range ways {
		if filter(w) {
			result = append(result, w)
		}
	}
	return result
}

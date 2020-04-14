package writer

import "github.com/EricNeid/go-netconvert/osm"
import "github.com/EricNeid/go-netconvert/internal/util"

const delimiter = "\t"

var headerElmarLinks []string = []string{
	"LINK_ID",
	"NODE_ID_FROM",
	"NODE_ID_TO",
	"BETWEEN_NODE_ID",
	"LENGTH",
	"VEHICLE_TYPE",
	"FORM_OF_WAY",
	"BRUNNEL_TYPE",
	"FUNCTIONAL_ROAD_CLASS",
	"SPEED_CATEGORY",
	"NUMBER_OF_LANES",
	"SPEED_LIMIT",
	"SPEED_RESTRICTION",
	"NAME_ID1_REGIONAL",
	"NAME_ID2_LOCAL",
	"HOUSENUMBERS_RIGHT",
	"HOUSENUMBERS_LEFT",
	"ZIP_CODE",
	"AREA_ID",
	"SUBAREA_ID",
	"THROUGH_TRAFFIC",
	"SPECIAL_RESTRICTIONS",
	"EXTENDED_NUMBER_OF_LANES",
	"ISRAMP",
	"CONNECTION",
}
var headerElmarNames []string = []string{
	"NAME_ID",
	"PERMANENT_ID_INFO",
	"NAME",
}

var log = util.Log{Context: "elmar"}

// elmarWay is an intermediate model to convert osm ways and nodes
// to elmar format.
type elmarWay struct {
	way   osm.Way
	edges []nodeTupel
}

type nodeTupel struct {
	from osm.Node
	to   osm.Node
}

type elmarEdge struct {
	// tbd
}

// AsElmarFormat writes the given net to filesystem using
// the elmar format.
func AsElmarFormat(net *osm.Net, baseName string) {
	toElmarWays(net)
}

func toElmarWays(net *osm.Net) []elmarWay {
	// create map of nodes for easier access
	var nodes = make(map[int64]osm.Node)
	for _, n := range net.Nodes {
		nodes[n.ID] = n
	}

	// parse osm data into elmar structs
	var elmarWays []elmarWay
	for _, w := range net.Ways {
		newElmarWay := elmarWay{
			way: w,
		}
		for i := 0; i < len(w.NodeRefs)-1; i++ {
			startRef := w.NodeRefs[i].NodeID
			endRef := w.NodeRefs[i+1].NodeID

			newElmarWay.edges = append(newElmarWay.edges, nodeTupel{
				from: nodes[startRef],
				to:   nodes[endRef],
			})
		}

		elmarWays = append(elmarWays, newElmarWay)
	}

	return elmarWays
}

func getNames(ways []osm.Way) (names []string) {
	uniqueNames := map[string]bool{}
	for _, w := range ways {
		for _, t := range w.Tags {
			if t.IsName() || t.IsRegName() {
				name := t.Value
				_, present := uniqueNames[name]
				if !present {
					uniqueNames[name] = true
					names = append(names, name)
				}
			}
		}
	}
	return names
}

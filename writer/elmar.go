package writer

import (
	"fmt"
	"os"
	"strings"

	"github.com/EricNeid/go-netconvert/internal/util"
	"github.com/EricNeid/go-netconvert/osm"
)

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

var headerElmarNodes []string = []string{
	"NODE_ID",
	"IS_BETWEEN_NODE",
	"amount_of_geocoordinates",
	"x1	y1	[x2 y2  ... xn  yn]",
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

// elmarLink represent a link between 2 nodes in elmar format.
// It can be converted to csv string.
type elmarLink struct {
	linkID     string
	nodeIDFrom int64
	nodeIDTo   int64
	// tbd
}

// toCSVString converts this elmar link to csv.
func (l elmarLink) toCSVString() string {
	return fmt.Sprintf("%s%s%d%s%d\n",
		l.linkID,
		delimiter,
		l.nodeIDFrom,
		delimiter,
		l.nodeIDTo,
	)
}

// elmarNode represent a single node in elmar format.
// It can be converted to csv string.
type elmarNode struct {
	nodeID      int64
	coordinates []floatTupel
}

type floatTupel struct {
	x float32
	y float32
}

// toCSVString converts the given elmar node to csv.
func (n elmarNode) toCSVString() string {
	csv := ""
	// write static information
	csv += fmt.Sprintf("%d%s%d%s%d",
		n.nodeID,
		delimiter,
		-1,
		delimiter,
		len(n.coordinates),
	)
	// write all available coordinates
	for _, c := range n.coordinates {
		csv += fmt.Sprintf("%s%f%s%f",
			delimiter,
			c.x,
			delimiter,
			c.y,
		)
	}
	// new line
	csv += delimiter + "\n"
	return csv
}

// elmarName represents a single name in elmar format. This name is either
// a local name or regional name. The nameID is unique, regardles of type.
type elmarName struct {
	nameID      int64
	name        string
	isLocalName bool
}

// AsElmarFormat writes the given net to filesystem using
// the elmar format.
func AsElmarFormat(net *osm.Net, baseName string) {
	ways := toElmarWays(net)
	nodes := toElmarNodes(net)

	writeWaysAsElmarFormat(ways, baseName+"_links.txt")
	writeNodesAsElmarFormat(nodes, baseName+"_nodes.txt")
}

func toElmarNodes(net *osm.Net) []elmarNode {
	nodes := make(map[int64][]osm.Node)
	for _, n := range net.Nodes {
		nodes[n.ID] = append(nodes[n.ID], n)
	}

	var elmarNodes []elmarNode
	for nodeID, nodes := range nodes {
		newElmarNode := elmarNode{
			nodeID: nodeID,
		}
		for _, n := range nodes {
			newTupel := floatTupel{
				x: n.Lon,
				y: n.Lat,
			}
			newElmarNode.coordinates = append(newElmarNode.coordinates, newTupel)
		}
		elmarNodes = append(elmarNodes, newElmarNode)
	}

	return elmarNodes
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

func toElmarNames() map[string]elmarName {

	// tbd
	return nil
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

func writeWaysAsElmarFormat(ways []elmarWay, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(strings.Join(headerElmarLinks, delimiter))
	f.WriteString("\n")
	for _, w := range ways {
		links := toElmarLinks(w)
		for _, l := range links {
			f.WriteString(l.toCSVString())
		}
	}
	return nil
}

func writeNodesAsElmarFormat(nodes []elmarNode, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(strings.Join(headerElmarNodes, delimiter))
	f.WriteString("\n")
	for _, n := range nodes {
		f.WriteString(n.toCSVString())
	}
	return nil
}

func toElmarLinks(way elmarWay) []elmarLink {
	var edges []elmarLink
	for _, e := range way.edges {
		newEdge := elmarLink{
			linkID:     fmt.Sprintf("%d_%d_%d", way.way.ID, e.from.ID, e.to.ID),
			nodeIDFrom: e.from.ID,
			nodeIDTo:   e.to.ID,
		}
		edges = append(edges, newEdge)
	}
	return edges
}

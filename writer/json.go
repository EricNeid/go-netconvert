package writer

import (
	"encoding/json"
	"io/ioutil"

	"github.com/EricNeid/go-netconvert/osm"
)

// AsJSON writes the given network to filesystem by using the
// json format
func AsJSON(net osm.Net, baseName string) error {
	err := NodesAsJSON(net.Nodes, baseName+".nodes.json")
	if err != nil {
		return err
	}
	return WaysAsJSON(net.Ways, baseName+".ways.json")
}

// NodesAsJSON writes given list of nodes in
// json format to given file.
func NodesAsJSON(nodes []osm.Node, file string) error {
	json, err := json.MarshalIndent(nodes, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, json, 0644)
}

// WaysAsJSON writes given list of ways in
// json format to given file.
func WaysAsJSON(ways []osm.Way, file string) error {
	json, err := json.MarshalIndent(ways, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, json, 0644)
}

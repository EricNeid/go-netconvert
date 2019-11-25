package writer

import (
	"encoding/json"
	"io/ioutil"

	"github.com/EricNeid/go-netconvert/osm"
)

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

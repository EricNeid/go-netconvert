package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/EricNeid/go-netconvert"
	"github.com/EricNeid/go-netconvert/filter"
	"github.com/EricNeid/go-netconvert/internal/util"
	"github.com/EricNeid/go-netconvert/osm"
	"github.com/EricNeid/go-netconvert/writer"
)

func main() {

	// read user input
	args, err := parseArgs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	xmlFile := args.xmlFile
	filterTags := args.filterTags

	// parse file
	net, err := parseFile(xmlFile)
	if err != nil {
		util.Error("main", err)
		return
	}
	fmt.Printf("Finished parsing xml\n")
	fmt.Printf("  Number of nodes: %d\n", len(net.Nodes))
	fmt.Printf("  Number of ways:  %d\n", len(net.Ways))

	// process parsed data
	net = filterNet(net, filterTags)
	fmt.Printf("Finished parsing xml\n")
	fmt.Printf("  Number of nodes: %d\n", len(net.Nodes))
	fmt.Printf("  Number of ways:  %d\n", len(net.Ways))

	// use name of input file as base name
	baseName, err := fileName(xmlFile)
	if err != nil {
		baseName = "out"
	}
	writeResult(net, baseName)
}

func parseFile(xmlFile string) (*osm.Net, error) {
	fmt.Printf("Reading input\n")
	defer util.TimeTrack(time.Now(), "Parsing")
	return netconvert.Decode(xmlFile)
}

func filterNet(net *osm.Net, tagNames string) *osm.Net {
	fmt.Printf("Filtering tags\n")
	defer util.TimeTrack(time.Now(), "Filtering")

	nodes := filter.Nodes(net.Nodes, func(n osm.Node) bool {
		for _, t := range n.Tags {
			if strings.Contains(tagNames, t.Name) {
				return true
			}
		}
		return false
	})
	ways := filter.Ways(net.Ways, func(n osm.Way) bool {
		for _, t := range n.Tags {
			if strings.Contains(tagNames, t.Name) {
				return true
			}
		}
		return false
	})

	return &osm.Net{
		Nodes: nodes,
		Ways:  ways,
	}
}

func writeResult(net *osm.Net, baseName string) {
	fmt.Printf("Writing output\n")
	defer util.TimeTrack(time.Now(), "Writing")
	writer.NodesAsJSON(net.Nodes, baseName+".nodes.json")
	writer.WaysAsJSON(net.Ways, baseName+".ways.json")
}

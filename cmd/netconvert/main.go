package main

import (
	"fmt"
	"strconv"
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
	fmt.Printf("  Number of nodes: %d\n", len(net.Nodes))
	fmt.Printf("  Number of ways:  %d\n", len(net.Ways))

	// process parsed data
	net, err = filterNet(net, filterTags)
	if err != nil {
		util.Error("main", err)
		return
	}
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

func filterNet(net *osm.Net, filterTags string) (*osm.Net, error) {
	fmt.Printf("Filtering tags\n")
	defer util.TimeTrack(time.Now(), "Filtering")

	if filterTags == "" {
		return net, nil
	}
	filterList, err := filter.ToFilter(filterTags)
	if err != nil {
		return net, err
	}

	nodes := filter.Nodes(net.Nodes, func(n osm.Node) bool {
		for _, t := range n.Tags {
			if matchesFilter(filterList, t) {
				return true
			}
		}
		return false
	})
	ways := filter.Ways(net.Ways, func(n osm.Way) bool {
		for _, t := range n.Tags {
			if matchesFilter(filterList, t) {
				return true
			}
		}
		return false
	})

	return &osm.Net{
		Nodes: nodes,
		Ways:  ways,
	}, nil
}

func matchesFilter(filterList []filter.Filter, tag osm.Tag) bool {
	for _, f := range filterList {
		if f.Name == tag.Name {
			switch f.Operand {
			case filter.NOP:
				return true
			case filter.EQ:
				return f.Value == tag.Value
			case filter.LT:
				valueTag, err := strconv.Atoi(tag.Value)
				if err != nil {
					return f.Value == tag.Value
				}
				return valueTag < f.ValueInt
			case filter.GT:
				valueTag, err := strconv.Atoi(tag.Value)
				if err != nil {
					return f.Value == tag.Value
				}
				return valueTag > f.ValueInt
			}
		}
	}
	return false
}

func writeResult(net *osm.Net, baseName string) {
	fmt.Printf("Writing output\n")
	defer util.TimeTrack(time.Now(), "Writing")
	writer.NodesAsJSON(net.Nodes, baseName+".nodes.json")
	writer.WaysAsJSON(net.Ways, baseName+".ways.json")
}

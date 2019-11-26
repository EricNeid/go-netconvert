package main

import (
	"fmt"
	"time"

	"github.com/EricNeid/go-netconvert"
	"github.com/EricNeid/go-netconvert/internal/util"
	"github.com/EricNeid/go-netconvert/osm"
	"github.com/EricNeid/go-netconvert/writer"
)

func main() {

	arg, err := parseArgs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	xmlFile := arg.xmlFile

	net, err := parseFile(xmlFile)
	if err != nil {
		util.Error("main", err)
		return
	}

	fmt.Printf("Finished parsing xml\n")
	fmt.Printf("Number of nodes: %d\n", len(net.Nodes))
	fmt.Printf("Number of ways:  %d\n", len(net.Ways))

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

func writeResult(net *osm.Net, baseName string) {
	fmt.Printf("Writing output\n")
	defer util.TimeTrack(time.Now(), "Writing")
	writer.NodesAsJSON(net.Nodes, baseName+".nodes.json")
	writer.WaysAsJSON(net.Ways, baseName+".ways.json")
}

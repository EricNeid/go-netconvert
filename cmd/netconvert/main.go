package main

import (
	"fmt"
	"os"
	"time"

	"github.com/EricNeid/go-netconvert"
	"github.com/EricNeid/go-netconvert/internal/util"
	"github.com/EricNeid/go-netconvert/osm"
	"github.com/EricNeid/go-netconvert/writer"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("Usage: netconvert <pathToXml>")
		return
	}

	xmlFile := arg[0]

	net, err := parseFile(xmlFile)
	if err != nil {
		util.Error("main", err)
		return
	}

	fmt.Printf("Finished parsing xml\n")
	fmt.Printf("Number of nodes: %d\n", len(net.Nodes))
	fmt.Printf("Number of ways:  %d\n", len(net.Ways))

	fmt.Printf("Writing output\n")
	writer.NodesAsJSON(net.Nodes, "out-nodes.json")
	writer.WaysAsJSON(net.Ways, "out-ways.json")
}

func parseFile(xmlFile string) (*osm.Net, error) {
	fmt.Printf("Reading input\n")
	defer util.TimeTrack(time.Now(), "Parsing")
	return netconvert.Decode(xmlFile)
}

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/EricNeid/go-netconvert"
	"github.com/EricNeid/go-netconvert/internal/util"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("Usage: netconvert <pathToXml>")
		return
	}

	defer util.TimeTrack(time.Now(), "Parsing")
	xmlFile := arg[0]
	net, err := netconvert.Decode(xmlFile)
	if err != nil {
		util.Error("main", err)
		return
	}

	fmt.Printf("Finished parsing xml\n")
	fmt.Printf("Number of nodes: %d\n", len(net.Nodes))
	fmt.Printf("Number of ways:  %d\n", len(net.Ways))
}

# go-netconvert

A simple parser for osm.xml files. Reads nodes and ways into golang structs. Provides util functions
to filter and to write back parsed network.

A command line tool to use the functionalities is also provided: [netconvert](cmd/netconvert/)

## Go get

```bash
go get github.com/EricNeid/go-netconvert
```

## Usage

```go
import "github.com/EricNeid/go-netconvert"

net, err := netconvert.Decode(xmlFile)
if err != nil {
  // handle error
  return
}

fmt.Printf("Finished parsing xml\n")
fmt.Printf("Number of nodes: %d\n", len(net.Nodes))
fmt.Printf("Number of ways:  %d\n", len(net.Ways))
```

package netconvert

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/EricNeid/go-netconvert/internal/util"
	"github.com/EricNeid/go-netconvert/osm"
)

var log = util.Log{Context: "decode"}

// Decode parses given xmlFile (which should be osm.xml) and returns
// the parsed net (containing edges and nodes).
func Decode(xmlFile string) (*osm.Net, error) {
	var nodes []osm.Node
	var ways []osm.Way

	f, err := os.Open(xmlFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "node" {
				n, err := DecodeNode(decoder, &se)
				if err != nil {
					log.E("decode", err)
				} else {
					nodes = append(nodes, n)
				}
			} else if se.Name.Local == "way" {
				w, err := DecodeWay(decoder, &se)
				if err != nil {
					log.E("decode", err)
				} else {
					ways = append(ways, w)
				}
			}
		}
	}
	return &osm.Net{
		Nodes: nodes,
		Ways:  ways,
	}, nil
}

// DecodeNode parsed given start element into instance of Node
func DecodeNode(decoder *xml.Decoder, se *xml.StartElement) (osm.Node, error) {
	var n osm.Node
	err := decoder.DecodeElement(&n, se)
	return n, err
}

// DecodeWay parses given start element into instance of Way
func DecodeWay(decoder *xml.Decoder, se *xml.StartElement) (osm.Way, error) {
	var w osm.Way
	err := decoder.DecodeElement(&w, se)
	return w, err
}

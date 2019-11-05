package osm

import "encoding/xml"

// Node is a single osm node
type Node struct {
	ID   string  `xml:"id,attr"`
	Lat  float32 `xml:"lat,attr"`
	Lon  float32 `xml:"lon,attr"`
	Tags []Tag   `xml:"tag"`
}

// DecodeNode parsed given start element into instance of Node
func DecodeNode(decoder *xml.Decoder, se *xml.StartElement) (Node, error) {
	var n Node
	err := decoder.DecodeElement(&n, se)
	return n, err
}

package osm

// Node is a single osm node
type Node struct {
	ID   string  `xml:"id,attr"`
	Lat  float32 `xml:"lat,attr"`
	Lon  float32 `xml:"lon,attr"`
	Tags []Tag   `xml:"tag"`
}

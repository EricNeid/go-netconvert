package osm

// Tag represents a single osm tag. It is used by
// both way and edges.
type Tag struct {
	Name  string `xml:"k,attr"`
	Value string `xml:"v,attr"`
}

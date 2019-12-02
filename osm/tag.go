package osm

// Tag represents a single osm tag. It is used by
// both way and edges.
type Tag struct {
	Name  string `xml:"k,attr"`
	Value string `xml:"v,attr"`
}

// IsName returns true if the given tag describes a name.
func (t Tag) IsName() bool {
	return t.Name == "name"
}

// IsRegName returns true if the given tag describes a regional name.
func (t Tag) IsRegName() bool {
	return t.Name == "reg_name"
}

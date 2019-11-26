package osm

// Way is a single osm way
//
// Sample of pared xml:
//
// 	<way id="4040481">
// 		<nd ref="21380156"/>
// 		<nd ref="243068174"/>
// 		<tag k="highway" v="motorway"/>
// 		<tag k="int_ref" v="E 26"/>
// 		<tag k="lanes" v="2"/>
// 		<tag k="lit" v="no"/>
// 		<tag k="maxspeed" v="none"/>
// 		<tag k="oneway" v="yes"/>
// 		<tag k="ref" v="A 24"/>
// 		<tag k="source:lit" v="http://www.autobahn-bilder.de"/>
// 		<tag k="surface" v="asphalt"/>
// 	</way>
type Way struct {
	ID       string    `xml:"id,attr"`
	NodeRefs []NodeRef `xml:"nd"`
	Tags     []Tag     `xml:"tag"`
}

// NodeRef is a single node reference in osm way
type NodeRef struct {
	NodeID string `xml:"ref,attr"`
}

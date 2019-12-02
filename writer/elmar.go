package writer

import "github.com/EricNeid/go-netconvert/osm"

import "os"

import "strings"

import "fmt"

const delimiter = "\t"

var elmarWaysHeader []string = []string{
	"LINK_ID",
	"NODE_ID_FROM",
	"NODE_ID_TO",
	"BETWEEN_NODE_ID",
	"LENGTH",
	"VEHICLE_TYPE",
	"FORM_OF_WAY",
	"BRUNNEL_TYPE",
	"FUNCTIONAL_ROAD_CLASS",
	"SPEED_CATEGORY",
	"NUMBER_OF_LANES",
	"SPEED_LIMIT",
	"SPEED_RESTRICTION",
	"NAME_ID1_REGIONAL",
	"NAME_ID2_LOCAL",
	"HOUSENUMBERS_RIGHT",
	"HOUSENUMBERS_LEFT",
	"ZIP_CODE",
	"AREA_ID",
	"SUBAREA_ID",
	"THROUGH_TRAFFIC",
	"SPECIAL_RESTRICTIONS",
	"EXTENDED_NUMBER_OF_LANES",
	"ISRAMP",
	"CONNECTION",
}
var elmarNameHeader []string = []string{
	"NAME_ID",
	"PERMANENT_ID_INFO",
	"NAME",
}

type tagName struct {
	name    string
	regName string
}

// AsElmarFormat writes the given net to filesystem using
// the elmar format.
func AsElmarFormat(net *osm.Net, baseName string) {

	names := getNames(net.Ways)
	writeNamesAsElmarFormat(names, baseName+"names.csv")
}

func getNames(ways []osm.Way) []tagName {
	var names []tagName
	count := 0
	for _, w := range ways {
		for _, t := range w.Tags {
			var name string
			var regName string
			if t.IsName() {
				name = t.Value
			} else if t.IsRegName() {
				regName = t.Value
			}

			if name != "" || regName != "" {
				names = append(names, tagName{
					name:    name,
					regName: regName,
				})
				count = count + 1
			}
		}
	}

	return names
}

func writeNamesAsElmarFormat(names []tagName, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(strings.Join(elmarNameHeader, delimiter))
	f.WriteString("\n")
	for i, n := range names {
		name := n.name
		if n.regName != "" {
			name = n.regName
		}
		f.WriteString(fmt.Sprintf("%d%s-1%s%s\n", i, delimiter, delimiter, name))
	}
	return nil
}

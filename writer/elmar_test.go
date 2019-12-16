package writer

import (
	"testing"

	"github.com/EricNeid/go-netconvert/internal/verify"
	"github.com/EricNeid/go-netconvert/osm"
)

func TestGetNames_shouldNotContainDuplicates(t *testing.T) {
	// arrange
	testData := []osm.Way{
		osm.Way{
			Tags: []osm.Tag{
				osm.Tag{
					Name:  "name",
					Value: "street",
				},
			},
		},
		osm.Way{
			Tags: []osm.Tag{
				osm.Tag{
					Name:  "name",
					Value: "street",
				},
			},
		},
	}
	// action
	resultNames := getNames(testData)
	// verify
	verify.Equals(t, 1, len(resultNames))
}

func TestWriteNamesAsElmarFormat(t *testing.T) {
	// arrange
	testData := []string{"name1", "name2"}
	// action
	err := writeNamesAsElmarFormat(testData, "../testdata/elmar.names.csv")
	// verify
	verify.Ok(t, err)
}

func TestWriteLinksAsElmarFormat(t *testing.T) {
	// arrange
	testData := []link{
		link{
			id:         1,
			nodeIDFrom: 2,
			nodeIDTo:   3,
		},
	}
	// action
	err := writeLinksAsElmarFormat(testData, "../testdata/elmar.links.csv")
	// verify
	verify.Ok(t, err)
}

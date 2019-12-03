package writer

import (
	"testing"

	"github.com/EricNeid/go-netconvert/internal/test"
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
	test.Equals(t, 1, len(resultNames))
}

func TestWriteNamesAsElmarFormat(t *testing.T) {
	// arrange
	testData := []string{"name1", "name2"}
	// action
	err := writeNamesAsElmarFormat(testData, "../testdata/elmar.names.csv")
	// verify
	test.Ok(t, err)
}

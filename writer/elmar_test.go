package writer

import (
	"testing"

	"github.com/EricNeid/go-netconvert/internal/test"
)

func TestWriteNamesAsElmarFormat(t *testing.T) {
	// arrange
	testData := []tagName{
		tagName{"", "regNam1"},
		tagName{"name2", ""},
	}
	// action
	err := writeNamesAsElmarFormat(testData, "../testdata/elmar.names.csv")
	// verify
	test.Ok(t, err)
}

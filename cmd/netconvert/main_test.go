package main

import (
	"testing"

	"github.com/EricNeid/go-netconvert/internal/verify"
)

func TestParseFile(t *testing.T) {
	// action
	result, err := parseFile("../../testdata/sample.osm.xml")

	// verify
	verify.Ok(t, err)
	verify.Assert(t, result != nil, "Parsing net failed")
}

func TestTilterNet(t *testing.T) {
	// arrange
	net, _ := parseFile("../../testdata/sample.osm.xml")

	// action
	result, err := filterNet(net, "maxheight<15")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, 0, len(result.Nodes))
	verify.Equals(t, 4, len(result.Ways))

	// action
	result, err = filterNet(net, "maxheight")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, 8, len(result.Nodes))
	verify.Equals(t, 15, len(result.Ways))

	// action
	result, err = filterNet(net, "obstacle=bridge")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, 7, len(result.Nodes))
	verify.Equals(t, 0, len(result.Ways))
}

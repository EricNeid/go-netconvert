package main

import "testing"
import "github.com/EricNeid/go-netconvert/internal/verify"

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
	verify.Equals(t, 7, len(result.Ways))

	// action
	result, err = filterNet(net, "maxheight")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, 55, len(result.Nodes))
	verify.Equals(t, 22, len(result.Ways))

	// action
	result, err = filterNet(net, "obstacle=bridge")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, 53, len(result.Nodes))
	verify.Equals(t, 0, len(result.Ways))
}

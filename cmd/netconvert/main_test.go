package main

import "testing"

import "github.com/EricNeid/go-netconvert/internal/test"

func TestParseFile(t *testing.T) {
	// action
	result, err := parseFile("../../testdata/sample.osm.xml")

	// verify
	test.Ok(t, err)
	test.Assert(t, result != nil, "Parsing net failed")
}

func TestTilterNet(t *testing.T) {
	// arrange
	net, _ := parseFile("../../testdata/sample.osm.xml")

	// action
	result, err := filterNet(net, "maxheight<15")
	// verify
	test.Ok(t, err)
	test.Equals(t, 0, len(result.Nodes))
	test.Equals(t, 7, len(result.Ways))

	// action
	result, err = filterNet(net, "maxheight")
	// verify
	test.Ok(t, err)
	test.Equals(t, 55, len(result.Nodes))
	test.Equals(t, 22, len(result.Ways))

	// action
	result, err = filterNet(net, "obstacle=bridge")
	// verify
	test.Ok(t, err)
	test.Equals(t, 53, len(result.Nodes))
	test.Equals(t, 0, len(result.Ways))
}

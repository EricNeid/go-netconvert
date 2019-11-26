# netconvert

A command line tool to process and filter osm.xml.

## Usage

There are some command line arguments to filter input.

The following command reads testdata/sample.oms.xml file, drops
all nodes and ways that do not contain tag named a,b or c.
The result is saved to testdata/sample.osm.xml.nodes.json and
testdata/sample.osm.xml.ways.json.

```bash
netconvert --filter-tags=a,b,c testdata/sample.osm.xml
```

```bash
netconvert --filter-tags=maxheigh>30,maxweight<20 testdata/sample.osm.xml
```

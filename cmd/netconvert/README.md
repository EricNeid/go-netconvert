# netconvert

A command line tool to process and filter osm.xml.

## Usage

Simple command to convert osm.xml to json:

```bash
netconvert testdata/sample.osm.xml
```

You can also filter the xml for specific tags. The following statement would
only write nodes and ways wich have a tag with name of either a or b:

```bash
netconvert --filter-tags=a,b testdata/sample.osm.xml
```

It is also possible to filter with simple conditions. The following
command would return only nodes and ways which a maxheigh of 30 or which are bridges:

```bash
netconvert --filter-tags=maxheigh<30,obstacle=bridge testdata/sample.osm.xml
```

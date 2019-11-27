package main

import "flag"
import "fmt"
import "os"
import "errors"

type args struct {
	filterTags string
	xmlFile    string
}

func parseArgs() (args, error) {
	flag.Usage = func() {
		fmt.Printf("Usage: %s <pathToXml>\n", os.Args[0])
		flag.PrintDefaults()
	}

	var args args
	flag.StringVar(&args.filterTags, "filter-tags", "", "filter result to contain only ways and nodes with the specified tags")

	flag.Parse()

	if len(flag.Args()) == 0 {
		return args, errors.New("no xml file given, use -h to show help")
	}
	args.xmlFile = flag.Arg(0)

	return args, nil
}

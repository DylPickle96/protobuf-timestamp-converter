package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	timeString = flag.String("time", "", "This flag is equal to the time you are trying to convert to a Protobuf Timestamp, required")
)

func main() {
	flag.Parse()
	usage()

	if *timeString == "" {
		fmt.Println("You did not provide a time string. Exiting program.")
		os.Exit(1)
	}

	fmt.Println("\nYour Protobuf Timestamp:\n", CreateProtobufTimestamp(*timeString))
}

func usage() {
	usage := `
This is a helper tool to make it easy for you as the developer to include Protobuf Timestamps in your Test Annotation Language.
This will take a date time string in the format of a RFC-3339.

more information: 
	(https://en.wikipedia.org/wiki/ISO_8601), 
	(https://tools.ietf.org/html/rfc3339), 
	(https://unix.stackexchange.com/questions/120484/what-is-a-standard-command-for-printing-a-date-in-rfc-3339-format)

Examples: 
	2018-08-23T11:24:52-07:00,
	2018-08-23T11:29:19-07:00,
	2018-08-23T11:29:29-07:00

Please provide the following flag:

-time=someTimeString`

	fmt.Println(usage)
}

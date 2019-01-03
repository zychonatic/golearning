package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose = kingpin.Flag("verbose", "Verbose mode").Short('v').Bool()
	logfile = kingpin.Flag("logfile", "Location of logfile").Short('l').Default("/var/log/test.log").String()
)

func main () {
	kingpin.Parse()
	fmt.Printf("%v, %s\n", *verbose, *logfile)
}

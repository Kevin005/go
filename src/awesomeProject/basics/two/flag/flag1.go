package main

import (
	"flag"
)

func main() {
	ParseArgs()
}

func ParseArgs() {
	var confFilepath string = "goku.conf"
	flag.StringVar(&confFilepath, "c", "./configure.conf", "Please provide a valid configuration file path")
	flag.Parse()
}

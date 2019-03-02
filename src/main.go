package main

import (
	"flag"
	"os"
	"strings"
	_ "time"
)

var (
	dir string
	ext string
	url string
)

func init() {
	flag.StringVar(&dir, "dir", "", "The directory to be used when reading the files")
	flag.StringVar(&url, "url", "0.0.0.0:4000", "The url that'll be used to sent the data")
	flag.StringVar(&ext, "ext", ".txt,.json", "Filter the searched files by extension")
}

func main() {
	flag.Parse()

	// This client should be a console, that way it would be much more flexible
	// about every part of the process.

	if flag.NFlag() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	if !strings.HasSuffix(dir, "/") {
		dir = dir + "/"
	}

	Init()
}

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	_ "sort"
)

var (
	filetypes = map[string]fileparam{
		"sample.txt": {20, 10, 3},
		"input.txt":  {1000, 1000, 3},
	}
	activeParam fileparam
)

type fileparam struct {
	records, top, productLimit int
}

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

func part1(in io.Reader) {
	fmt.Printf("part1 not implemented\n")
}

var (
	part   int
	file   string
	silent bool
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")
	flag.BoolVar(&silent, "s", false, "silent?")
}

func main() {
	flag.Parse()
	if file != "" {
		var err error
		if os.Stdin, err = os.Open(file); err != nil {
			panic(err)
		}
	}
	activeParam = filetypes[path.Base(file)]
	switch part {
	case 2:
		part2(os.Stdin)
	default:
		part1(os.Stdin)
	}
}

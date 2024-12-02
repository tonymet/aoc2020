package main

import (
	"flag"
	"fmt"
	_ "io"
	"os"
	_ "sort"
)

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

type rec []int

func log(f string, val ...any) {
	if silent {
		return
	}
	fmt.Printf(f, val...)
}

func part2() {
	fmt.Printf("part2 not implemented\n")
}

func part1() {
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
	switch part {
	case 2:
		part2()
	default:
		part1()
	}
}

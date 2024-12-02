package main

import (
	"flag"
	"fmt"
	_ "io"
	"os"
	_ "sort"
)

func part2() {
	fmt.Printf("part2 not implemented\n")
}

func part1() {
	fmt.Printf("part1 not implemented\n")
}

var (
	part int
	file string
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")

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

package main

import (
	"flag"
	"fmt"
	_ "io"
	_ "os"
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
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
}
func main() {
	flag.Parse()
	switch part {
	case 2:
		part2()
	default:
		part1()
	}
}

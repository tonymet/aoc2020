package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	_ "sort"
)

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

type stonesType []int64

func part1(in io.Reader) {
	fmt.Printf("part1 not implemented\n")
	initStones := make(stonesType, 0, 5)
	for {
		var cur int64
		_, err := fmt.Fscan(in, &cur)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		initStones = append(initStones, cur)
	}
	fmt.Printf("initStones: %+v", initStones)
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
		part2(os.Stdin)
	default:
		part1(os.Stdin)
	}
}

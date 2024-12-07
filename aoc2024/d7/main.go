package main

// test

import (
	"flag"
	"fmt"
	"github.com/tonymet/aoc2020/shared"
	"io"
	"os"
	_ "sort"
)

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

func part1(in io.Reader) {
	shared.LineProcessor(in, func(line io.Reader) {
		// scan values
		var (
			want   int64
			params []int64
		)
		params = make([]int64, 0, 8)
		_, err := fmt.Fscanf(line, "%d:", &want)
		if err == io.EOF {
			panic(io.ErrUnexpectedEOF)
		} else if err != nil {
			panic(err)
		}
		for {
			var p int64
			_, err := fmt.Fscan(line, &p)
			if err == io.EOF {
				// expected
				break
			} else if err != nil {
				panic(err)
			}
			params = append(params, p)
		}
		fmt.Printf("w: %d ", want)
		fmt.Printf("p: %+v\n", params)
	})
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

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	_ "sort"
)

func part2() {
	fmt.Printf("part2 not implemented\n")
}

func part1() {
	var row [5]int
	for {
		_, err := fmt.Scanf("%d %d %d %d %d", &row[0], &row[1], &row[2], &row[3], &row[4])
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("%+x\n", row)
	}
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

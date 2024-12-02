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

type rec [5]int

func part1() {
	var row rec
	for {
		_, err := fmt.Scanf("%d %d %d %d %d", &row[0], &row[1], &row[2], &row[3], &row[4])
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("%+x\n", row)
		fmt.Printf("desc : %+v\n", row.ascDesc())
	}
}

func (row rec) ascDesc() bool {
	// test if all asc or desc
	// test all desc
	// l , r
	// compare to adjascent
	// true if they meet, false if not
	l, r := 0, len(row)-1
	for {
		// if l == r, break
		if l == r {
			break
		}
		// different signs
		if (row[l+1]-row[l])*(row[r]-row[r-1]) < 0 {
			return false
		}
		l, r = l+1, r-1
	}
	return true
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

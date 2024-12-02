package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	_ "sort"
	"strings"
)

func part2() {
	fmt.Printf("part2 not implemented\n")
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

type rec []int

func part1() {
	lineReader := bufio.NewScanner(os.Stdin)
	var sum = 0
	for lineReader.Scan() {
		// scan line then scan the goods
		line := lineReader.Text()
		in := strings.NewReader(line)
		var val int
		var row rec = make(rec, 0, 8)
		for {
			_, err := fmt.Fscan(in, &val)
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			row = append(row, val)

		}
		if row.safe() {
			sum += 1
		}
		fmt.Printf("%+x, safe: %t\n", row, row.safe())
	}
	fmt.Printf("sum : %d\n", sum)
}

func (row rec) safe() bool {
	return row.ascDesc() && row.gap(1, 3)
}

func (row rec) ascDesc() bool {
	l, r := 0, len(row)-1
	for {
		if l == r {
			break
		}
		if (row[l+1]-row[l])*(row[r]-row[r-1]) < 0 {
			return false
		}
		l, r = l+1, r-1
	}
	return true
}

func (row rec) gap(tl, th int) bool {
	l, r := 0, len(row)-1
	for {
		// if l == r, break
		if l == r {
			break
		}
		// different signs
		if abs(row[l+1]-row[l]) < tl || abs(row[l+1]-row[l]) > th ||
			abs(row[r]-row[r-1]) < tl || abs(row[r]-row[r-1]) > th {
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

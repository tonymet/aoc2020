package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	_ "sort"
	"strconv"
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
		if row.safe(1, 3) {
			sum += 1
		}
		fmt.Printf("%v, safe: %t\n", row, row.safe(1, 3))
	}
	fmt.Printf("sum : %d\n", sum)
}

func (row rec) String() string {
	b := make([]byte, 0, len(row)*3)
	for _, v := range row {
		b = strconv.AppendInt(b, int64(v), 10)
		b = append(b, ' ')
	}
	return string(b)
}
func (row rec) safe(tl, th int) bool {
	l, r := 0, len(row)-1
	for {
		if l > r {
			break
		}
		// different signs
		if abs(row[l+1]-row[l]) < tl || abs(row[l+1]-row[l]) > th ||
			abs(row[r]-row[r-1]) < tl || abs(row[r]-row[r-1]) > th {
			return false
		}
		l, r = l+1, r-1
	}
	for i := 0; i < len(row)-2; i++ {
		if (row[i+1]-row[i])*(row[i+2]-row[i+1]) < 0 {
			return false
		}

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

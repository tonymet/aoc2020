package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func parseFile(input io.Reader) {
	lineReader := bufio.NewScanner(input)
	var sum = 0
	for lineReader.Scan() {
		// scan line then scan the goods
		line := lineReader.Text()
		in := strings.NewReader(line)
		var val int
		var row rec = make(rec, 0, len(line)/2)
		for {
			_, err := fmt.Fscan(in, &val)
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			row = append(row, val)
		}
		switch part {
		case 1:
			safe := row.safe(tlow, thigh)
			if safe {
				sum += 1
			}
			log("%v, safe: %t\n", row, safe)
		case 2:
			safe := row.safePart2()
			if safe {
				sum += 1
			}
			log("%v, safe: %t\n", row, safe)
		default:
			panic(fmt.Errorf("no part %d", part))
		}
	}
	log("sum : %d\n", sum)
}

func (row rec) String() string {
	b := make([]byte, 0, len(row)*3)
	for _, v := range row {
		b = strconv.AppendInt(b, int64(v), 10)
		b = append(b, ' ')
	}
	return string(b)
}

func (row rec) safePart2() bool {
	// iterate and remove an element
	// test for safety
	// if none are safe return false
	if row.safe(tlow, thigh) {
		return true
	}
	for i := 0; i < len(row); i++ {
		var testRow = make(rec, len(row))
		copy(testRow, row)
		testRow = append(testRow[:i], testRow[i+1:]...)
		if testRow.safe(tlow, thigh) {
			return true
		}
	}
	return false
}

func (row rec) safe(tl, th int) bool {
	l, r := 0, len(row)-1
	for {
		if l > r {
			break
		}
		// test gap
		if abs(row[l+1]-row[l]) < tl || abs(row[l+1]-row[l]) > th ||
			abs(row[r]-row[r-1]) < tl || abs(row[r]-row[r-1]) > th {
			return false
		}
		l, r = l+1, r-1
	}
	// test asc vs desc
	for i := 0; i < len(row)-2; i++ {
		if (row[i+1]-row[i])*(row[i+2]-row[i+1]) < 0 {
			return false
		}

	}
	return true
}

var (
	part        int
	file        string
	silent      bool
	tlow, thigh = 1, 3
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")
	flag.BoolVar(&silent, "s", false, "which exercise part?")
}

func main() {
	flag.Parse()
	var f *os.File
	if file != "" {
		var err error
		if f, err = os.Open(file); err != nil {
			panic(err)
		}
	}
	switch part {
	case 2:
		parseFile(f)
	default:
		parseFile(f)
	}
}

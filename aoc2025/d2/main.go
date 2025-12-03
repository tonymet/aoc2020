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

type myrange = [2]int64

type rangeReader struct {
	*bufio.Scanner
}

func NewScanner(in io.Reader) (rr rangeReader) {
	rr.Scanner = bufio.NewScanner(in)
	rr.Split(CSVSplitter)
	return
}

func (rr rangeReader) Range() (r myrange, err error) {
	t := rr.Text()
	parts := strings.Split(t, "-")
	for i, p := range parts {
		r[i], err = strconv.ParseInt(p, 10, 64)
		if err != nil {
			return
		}
	}
	return
}

// see bufio.Scanner for example splitter-wrappers
func CSVSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	loc := strings.IndexByte(string(data), ',')
	if loc == -1 {
		if !atEOF {
			return 0, nil, nil
		}
		return 0, nil, bufio.ErrFinalToken
	}
	return loc + 1, data[:loc], nil
}

func part1(in io.Reader) {
	reader := NewScanner(in)
	for reader.Scan() {
		rec, err := reader.Range()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%x\n", rec)
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
	part1(os.Stdin)
}

// if the len > pos in any direction,1 will be hit
// if the len > 100, 0 will be hit the
// pos - len  + len / 100

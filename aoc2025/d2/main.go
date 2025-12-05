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

type intRange = [2]int64

type rangeMeta struct {
	intRange
	strRange [2]string
}

type rangeReader struct {
	*bufio.Scanner
}

func (rm rangeMeta) String() string {
	return fmt.Sprintf("%s - %s (%d - %d )", rm.strRange[0], rm.strRange[1], rm.intRange[0], rm.intRange[1])
}

func NewScanner(in io.Reader) (rr rangeReader) {
	rr.Scanner = bufio.NewScanner(in)
	rr.Split(CSVSplitter)
	return
}

func (rr rangeReader) RangeMeta() (r rangeMeta, err error) {
	t := rr.Text()
	parts := strings.Split(t, "-")
	for i, p := range parts {
		r.intRange[i], err = strconv.ParseInt(p, 10, 64)
		r.strRange[i] = p
		if err != nil {
			return
		}
	}
	return
}

func cmpIntSlice(a, b []int64) (same bool) {
	same = true
	for i, _ := range a {
		if b[i] != a[i] {
			return false
		}
	}
	return
}

func divCmpInt(a int64) bool {
	// convert to int slice
	d := countDigits(a)
	intSlice := make([]int64, d)
	t := a
	for i := d - 1; i >= 0; i-- {
		intSlice[i] = t % 10
		t = t / 10
	}
	return divCmp(intSlice)
}

func divCmp(a []int64) (match bool) {
	match = false
	var last []int64
	for h := len(a) / 2; h >= 1; h-- {
		if len(a)%h != 0 {
			continue
		}
		last = a[0:h]
		for k := h; k <= len(a)-h; k += h {
			if match = cmpIntSlice(last, a[k:k+h]); !match {
				break
			}
			last = a[k : k+h]
		}
		if match {
			return true
		}
	}
	return
}

func halvesMatch(n int64) bool {
	if n < 0 {
		n = -n
	}
	if n == 0 {
		return false
	}
	totalDigits := countDigits(n)
	lowHalfLen := totalDigits / 2
	if lowHalfLen == 0 || totalDigits%2 == 1 {
		return false
	}
	var divisor int64 = 1
	for i := 0; i < lowHalfLen; i++ {
		divisor *= 10
	}
	lowSegment, highSegment := n%divisor, n/divisor
	return highSegment == lowSegment
}

func countDigits(n int64) (count int) {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
	}
	for n > 0 {
		n /= 10
		count++
	}
	return
}

func CSVSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := strings.IndexByte(string(data), ','); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data[:len(data)-1], nil
	}
	return 0, nil, nil
}

func sumInts(is []int64) (sum int64) {
	for _, v := range is {
		sum += v
	}
	return
}

func part2(in io.Reader) {
	reader := NewScanner(in)
	var sum int64
	for reader.Scan() {
		rec, err := reader.RangeMeta()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+s\n", rec)
		for i := rec.intRange[0]; i <= rec.intRange[1]; i++ {
			if divCmpInt(i) {
				fmt.Printf("found %d\t", i)
				sum += i
			}
		}
		fmt.Println("")
	}
	fmt.Printf("sums := %d\n", sum)
}
func part1(in io.Reader) {
	reader := NewScanner(in)
	var sum int64
	for reader.Scan() {
		rec, err := reader.RangeMeta()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+s\n", rec)
		for i := rec.intRange[0]; i <= rec.intRange[1]; i++ {
			if halvesMatch(i) {
				fmt.Printf("found %d\t", i)
				sum += i
			}
		}
		fmt.Println("")
	}
	fmt.Printf("sums := %d\n", sum)
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
	case 1:
		part1(os.Stdin)
	case 2:
		part2(os.Stdin)
	default:
		panic("no part")
	}
}

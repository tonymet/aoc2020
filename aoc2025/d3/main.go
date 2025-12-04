package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	part int
	file string
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")
}

func stringInts(s string) (r []int64, err error) {
	r = make([]int64, 0, 20)
	for _, c := range s {
		v, err := strconv.ParseInt(string(c), 10, 64)
		if err != nil {
			panic(err)
		}
		r = append(r, v)
	}
	return r, nil
}

func j(l, r int64) (v int64) {
	return 10*l + r
}
func findMax(v []int64) int64 {
	// lh & rh pointers until l reaches r-1 and keep max
	var max int64
	for lh := 0; lh != len(v)-1; lh++ {
		for rh := len(v) - 1; rh != lh; rh-- {
			cur := j(v[lh], v[rh])
			if cur > max {
				max = cur
			}
		}
	}
	return max
}

func part1(in io.Reader) {
	scanner := bufio.NewScanner(in)
	var sum int64
	for scanner.Scan() {
		t := scanner.Text()
		ints, err := stringInts(t)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+x\t", ints)
		max := findMax(ints)
		fmt.Printf("max: %d\n", max)
		sum += max
	}
	fmt.Printf("sum: %d\n", sum)
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

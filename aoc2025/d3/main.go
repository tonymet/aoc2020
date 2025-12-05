package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var (
	part  int
	file  string
	quiet bool
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")
	flag.BoolVar(&quiet, "q", true, "quiet?")
}

func stringInts(s string) (r []int64, err error) {
	r = make([]int64, 0, 40)
	for _, c := range s {
		v, err := strconv.ParseInt(string(c), 10, 64)
		if err != nil {
			panic(err)
		}
		r = append(r, v)
	}
	return r, nil
}

func join2(l, r int64) (v int64) {
	return 10*l + r
}

func bigj(v string, l, m, r int) *big.Int {
	ret := big.NewInt(0)
	var final strings.Builder
	final.Grow(110)
	final.WriteString(v[0:l])
	final.WriteString(v[l+1 : m])
	final.WriteString(v[m+1 : r])
	final.WriteString(v[r+1:])
	finalS := final.String()
	if len(finalS) != (len(v) - 3) {
		panic("wrong len")
	}
	ret.SetString(finalS, 10)
	return ret
}

func findMax(v []int64) int64 {
	// lh & rh pointers until l reaches r-1 and keep max
	var max int64
	for lh := 0; lh != len(v)-1; lh++ {
		for rh := len(v) - 1; rh != lh; rh-- {
			cur := join2(v[lh], v[rh])
			if cur > max {
				max = cur
			}
		}
	}
	return max
}

func findMax2(s string) *big.Int {
	max := big.NewInt(0)
	for lh := 0; lh != len(s)-2; lh++ {
		for m := lh + 1; m != len(s)-1; m++ {
			for rh := len(s) - 1; rh != m; rh-- {
				cur := bigj(s, lh, m, rh)
				if max.Cmp(cur) == -1 {
					max = cur
				}
			}
		}
	}
	return max
}

// intPowerOfTen calculates 10^k using integer multiplication.
// Returns the result and a boolean indicating if an overflow occurred.
func intPowerOfTen(k int) (int64, bool) {
	if k < 0 {
		return 0, true // Cannot represent negative powers (like 0.1) as integers
	}

	var result int64 = 1
	var base int64 = 10

	for i := 0; i < k; i++ {
		// Overflow check: If result > MaxInt64 / 10, the next multiplication will overflow.
		if result > math.MaxInt64/base {
			return 0, true // Overflow detected
		}
		result *= base
	}
	return result, false
}

func findMax3(v []int64) int64 {
	// l = 0 , r = end - 12, then 11, then 10
	// tlen l = last + 1, r = end - 11
	var o [12]int64
	l := 0
	for cap := 11; cap >= 0; cap-- {
		r := len(v) - cap - 1
		lmax := v[l]
		l++
		for j := l; j <= r; j++ {
			if v[j] > lmax {
				lmax = v[j]
				l = j + 1
			}
		}
		o[12-cap-1] = lmax
	}
	// multiply
	var ret int64
	for k := 0; k < 12; k++ {
		ret *= 10
		ret += o[k]
	}
	return ret
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
		//fmt.Printf("%v\t", ints)
		max := findMax(ints)
		//fmt.Printf("max: %d\n", max)
		sum += max
	}
	fmt.Printf("sum: %d\n", sum)
}

func part2(in io.Reader) {
	scanner := bufio.NewScanner(in)
	var sum int64
	for scanner.Scan() {
		t := scanner.Text()
		ints, err := stringInts(t)
		if err != nil {
			panic(err)
		}
		if !quiet {
			fmt.Printf("%+x\t", ints)
		}
		max := findMax3(ints)
		if !quiet {
			fmt.Printf("max: %d\n", max)
		}
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
	switch part {
	case 1:
		part1(os.Stdin)
	case 2:
		part2(os.Stdin)
	default:
		log.Fatalf("not found part : %d", part)
	}
}

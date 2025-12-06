package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type mRange = [2]int64
type mRanges = []mRange

type ACL struct {
	mRanges
}

const (
	stateRules = 0
	stateIds   = 1
)

func intToStr(s []string) (ret [2]int64, err error) {
	if len(s) > 2 {
		panic("oob")
	}
	for i, v := range s {
		ret[i], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			break
		}
	}
	return
}

func mustParse(s string) (r int64) {
	ret, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return ret
}

func (a ACL) CountRanges() (c int64) {
	uniq := make(map[int64]bool)
	for _, v := range a.mRanges {
		c += v[1] - v[0] + 1
		for i := v[0]; i <= v[1]; i++ {
			uniq[i] = true
		}
	}
	// count keys
	return int64(len(uniq))
}

func intSort(ints []int64) {
	sort.Slice(ints, func(i, j int) bool {
		return ints[i] < ints[j]
	})
}

func (a ACL) CountRanges2() (c int64) {
	// order lows , order highs.
	// interate from the lowest low then turn off if encounter the lowest high. turn back on with next lowest low pop off lows when they are passed
	uniq := make(map[int64]bool)
	lows, highs := make([]int64, 0, 30), make([]int64, 0, 30)
	// build lows
	for _, v := range a.mRanges {
		lows = append(lows, v[0])
	}
	intSort(lows)
	for _, v := range a.mRanges {
		highs = append(highs, v[1])
	}
	intSort(highs)
	// now

	fmt.Printf("full range: %d", highs[len(highs)-1]-lows[0])
	// idState := true
	// min := lows[0]
	// lows = lows[1:]
	// for ; len(lows) != 0 && len(highs) != 0; min++ {
	// 	switch {
	// 	case min < highs[0] && idState:
	// 		uniq[min] = true
	// 	case min == highs[0]:
	// 		idState = false
	// 		highs = highs[1:]
	// 		uniq[min] = true
	// 	case min == lows[0]:
	// 		idState = true
	// 		// pop the front off
	// 		lows = lows[1:]
	// 		uniq[min] = true
	// 	}
	// }

	_ = uniq
	//return int64(len(uniq))
	countx := int64(0)
	for len(lows) != 0 && len(highs) != 0 {
		l := lows[0]
		lows = lows[1:]
		h := highs[0]
		highs = highs[1:]
		countx += h - l + 1
		// subtract

		if len(lows) > 1 && h < lows[0] {
			countx -= lows[0] - h + 1
		}
	}

	return countx

}

func (a ACL) Test(id int64) (pass bool) {
	pass = false
	for _, r := range a.mRanges {
		if id >= r[0] && id <= r[1] {
			return true
		}
	}
	return
}
func (a *ACL) ReadRanges(in io.Reader) (count int64) {
	// read in file
	state := stateRules
	scanner := bufio.NewScanner(in)
	a.mRanges = make(mRanges, 0, 20)
	for scanner.Scan() {
		line := scanner.Text()
		switch state {
		case stateRules:
			if line == "" {
				state = stateIds
				continue
			}
			p := strings.Split(line, "-")
			ints, err := intToStr(p)
			if err != nil {
				panic(err)
			}
			a.mRanges = append(a.mRanges, ints)
		case stateIds:
			id := mustParse(line)
			var stat bool
			if stat = a.Test(id); stat {
				count++
			}
			fmt.Printf("id : %d, pass: %t\n", id, stat)
			_ = 4
		default:
			panic("wrong state")
		}
	}
	fmt.Printf("len acl: %d,\t count: %d\n", len(a.mRanges), count)
	return
}

func part2(in io.Reader) {
	_ = in
	fmt.Printf("part2 not implemented\n")
}

func part1(in io.Reader) {
	var acl ACL
	acl.ReadRanges(in)
	count := acl.CountRanges2()
	fmt.Printf("full count : %d\n", count)
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

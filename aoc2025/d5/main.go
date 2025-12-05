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
	fmt.Printf("part2 not implemented\n")
}

func part1(in io.Reader) {
	var acl ACL
	acl.ReadRanges(in)
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

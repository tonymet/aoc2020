package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	_ "sort"
	_ "strconv"
	"strings"

	"github.com/stevenle/topsort/v2"
)

var (
	STATE_RULES   = 1
	STATE_UPDATES = 2
	ruleList      ruleset
)

type ruleset []rule

type rule struct {
	l, r int64
}
type updateSpec []int64

func (u updateSpec) fix() updateSpec {
	topo := topsort.NewGraph[int64]()
	index := make(map[int64]bool)
	for _, v := range u {
		index[v] = true
	}

	for _, rule := range ruleList {
		_, ok1 := index[rule.l]
		_, ok2 := index[rule.r]
		if ok1 && ok2 {
			topo.AddEdge(rule.r, rule.l)
		}
	}
	path, err := topo.TopSort(topo.RootNode())
	if err != nil {
		return updateSpec{}
	}
	return path
}

func (u updateSpec) checkRules() bool {
	for _, rule := range ruleList {
		lVal, rVal := rule.l, rule.r
		var l, r int
		lFound, rFound := false, false
		for l = 0; l < len(u); l++ {
			if u[l] == lVal {
				lFound = true
				break
			}
		}
		for r = len(u) - 1; r >= 0; r-- {
			if u[r] == rVal {
				rFound = true
				break
			}
		}
		if !(lFound && rFound) {
			continue
		}
		// l passed R
		if l >= r {
			return false
		}
	}
	return true
}

func (u updateSpec) indexValue() int64 {
	p := len(u) / 2
	return u[p]
}

func part1(in io.Reader) {
	var l, r, d int64
	ruleList = make([]rule, 0)
	sum := int64(0)
	sum2 := int64(0)
	for {
		n, err := fmt.Fscanf(in, "%d|%d", &l, &r)
		if err == io.EOF {
			break
		}
		if n == 0 {
			break
		} else if err != nil {
			break
		}
		curRule := rule{l, r}
		ruleList = append(ruleList, curRule)
	}
	lineReader := bufio.NewScanner(in)
	for lineReader.Scan() {
		var updateRecord = make(updateSpec, 0)
		line := lineReader.Text()
		numReader := strings.NewReader(line)
		for {
			var c = make([]byte, 1)
			_, err := fmt.Fscanf(numReader, "%d", &d)
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Fscan(numReader, &c)
			}
			updateRecord = append(updateRecord, d)
			// discard comma
			numReader.Read(c)
		}
		check := updateRecord.checkRules()
		fmt.Printf("UR: %+v, checkRules: %t \n", updateRecord, check)
		if check {
			sum += updateRecord.indexValue()
		} else {
			fixed := updateRecord.fix()
			fixChecked := fixed.checkRules()
			fmt.Printf("fixed: %+v, checkRules: %t \n", fixed, fixChecked)
			sum2 += fixed.indexValue()
		}

	}

	fmt.Printf("sum1 := %d ", sum)
	fmt.Printf("sum2 := %d ", sum2)
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
		part1(os.Stdin)
	default:
		part1(os.Stdin)
	}
}

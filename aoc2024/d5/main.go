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

var (
	STATE_RULES   = 1
	STATE_UPDATES = 2
	state         int
)

type rule struct {
	l, r int64
}
type ruleset []rule

func part1(in io.Reader) {
	fmt.Printf("part1 not implemented\n")
	var l, r, d int64
	var state = STATE_RULES
	aRuleset := make(map[int64]ruleset)
forloop:
	for {
		switch state {
		case STATE_RULES:
			n, err := fmt.Fscanf(in, "%d|%d", &l, &r)
			if err == io.EOF {
				break forloop
			}
			if n == 0 {
				break forloop
			} else if err != nil {
				break forloop
			}
			curRule := rule{l, r}
			{
				_, ok := aRuleset[l]
				if !ok {
					aRuleset[l] = make(ruleset, 0)
				}
				aRuleset[l] = append(aRuleset[l], curRule)
			}
			{
				_, ok := aRuleset[r]
				if !ok {
					aRuleset[r] = make(ruleset, 0)
				}
				aRuleset[r] = append(aRuleset[r], curRule)
			}
			fmt.Printf("l , r = %d, %d\n", l, r)
		}
	}
	lineReader := bufio.NewScanner(in)
	for lineReader.Scan() {
		var updateRecord = make([]int64, 0)
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
		fmt.Printf("UR: %+v\n", updateRecord)
	}
	fmt.Printf("ruleset %+x\n", aRuleset)

	// read orderings into a hash. num -> [n]orderings. both sides will be keyed
	// list of all scoped orderings
	// iter list . use l-R traversal to confirm instance of ordering is true
	// violation return false.
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
		part2()
	default:
		part1(os.Stdin)
	}
}

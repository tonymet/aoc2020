package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	_ "sort"
	"strconv"
	"strings"
)

var (
	filetypes = map[string]fileparam{
		"sample.txt": {20, 10, 3},
		"input.txt":  {1000, 1000, 3},
	}
	activeParam fileparam
)

type machine struct {
	lights  []rune
	buttons [][]int
	weights [][]int
}

const (
	stateLights int = iota
	stateButtons
	stateWeights
)

type fileparam struct {
	records, top, productLimit int
}

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

func splitInt(s string) (value []int, err error) {
	parts := strings.Split(s, ",")
	value = make([]int, 0, len(s))
	for _, v := range parts {
		v, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return value, err
		}
		value = append(value, int(v))
	}
	return
}
func initialLights(l int) (r []rune) {
	r = make([]rune, l)
	for i := range r {
		r[i] = '.'
	}
	return
}

func testSolution(target []rune, buttons [][]int) bool {
	var counts = make([]int, len(target))
	for _, v := range buttons {
		for _, pos := range v {
			counts[pos]++
		}
	}
	for i, r := range target {
		switch r {
		case '.':
			// check even
			if counts[i]%2 != 0 {
				return false
			}
		case '#':
			if counts[i]%2 != 1 {
				return false
			}
		}
	}
	return true
}

func Subsets(data [][]int, f func([][]int)) {
	n := len(data)
	// Total number of subsets is 2^n
	numSubsets := 1 << n

	for i := 0; i < numSubsets; i++ {
		var subset [][]int
		for j := 0; j < n; j++ {
			// Check if the j-th bit is set in mask i
			if (i>>j)&1 == 1 {
				subset = append(subset, data[j])
			}
		}
		f(subset)
	}
}

func isolve(in [][]int, t []rune) int {
	min := len(in)
	Subsets(in, func(test [][]int) {
		if testSolution(t, test) {
			if len(test) < min {
				min = len(test)
			}
		}
	})
	return min

}

func part1(in io.Reader) {
	scanner := bufio.NewScanner(in)
	sum := 0
	for scanner.Scan() {
		var cur machine
		cur.lights = make([]rune, 0, 10)
		cur.buttons = make([][]int, 0, 10)
		line := scanner.Text()
		fields := strings.Fields(line)
		for _, s := range fields {
			switch s[0] {
			case '[':
				cur.lights = []rune(s[1 : len(s)-1])
			case '(':
				val, err := splitInt(s[1 : len(s)-1])
				if err != nil {
					panic(err)
				}
				cur.buttons = append(cur.buttons, val)
			case '{':
				//pass
			}
		}
		fmt.Printf("cur: %+v\n", cur)
		count := isolve(cur.buttons, cur.lights)
		sum += count
		fmt.Printf("soln: %d\n", count)
	}
	fmt.Printf("p1 sum: %d\n", sum)
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
	activeParam = filetypes[path.Base(file)]
	switch part {
	case 2:
		part2(os.Stdin)
	default:
		part1(os.Stdin)
	}
}

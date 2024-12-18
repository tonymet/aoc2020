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

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

type vec struct {
	x, y int
}
type button struct {
	name byte
	dir  vec
}

type solution struct {
	a, b, score int
}

func (v vec) mult(x int) vec {
	return vec{v.x * x, v.y * x}
}

func (v vec) div(x int) vec {
	return vec{v.x / x, v.y / x}
}

func (v vec) mod2(v2 vec) vec {
	return vec{v.x % v2.x, v.y % v2.y}
}
func (v vec) mod(n int) vec {
	return vec{v.x % n, v.y % n}
}
func (v vec) diff(v2 vec) vec {
	return vec{v.x - v2.x, v.y - v2.y}
}

func (b button) String() string {
	var buf = make([]byte, 0, 10)
	buf = append(buf, b.name, ':', ' ')
	buf = append(buf, " x:"...)
	buf = strconv.AppendInt(buf, int64(b.dir.x), 10)
	buf = append(buf, " y:"...)
	buf = strconv.AppendInt(buf, int64(b.dir.y), 10)
	return string(buf)
}

var (
	STATE_BUTTON      = 1
	STATE_PRIZE       = 2
	STATE_NEWLINE     = 3
	state         int = STATE_BUTTON
	ErrUnsolvable     = fmt.Errorf("unsolvable")
)

func solve(bA, bB button, p vec) (solution, error) {
	// test simple division % = 0 4 cases
	if (p.mod2(bA.dir) == vec{0, 0}) {
		return solution{a: p.x / bA.dir.x, b: 0}, nil
	} else if (p.mod2(bB.dir) == vec{0, 0}) {
		// TOOD best solution
		return solution{a: 0, b: p.x / bB.dir.x}, nil
	}

	// subtract A and test difference mult B % = 0
	cur := p
	for cur.x > 0 && cur.y > 0 {
		cur = cur.diff(bB.dir)
		gap := p.diff(cur)
		if (cur.mod2(bA.dir) == vec{0, 0}) {
			// TODO find best solution here
			// calculate bA facotr, bB Facto4
			// return
			bAFactor := cur.x / bA.dir.x
			bBFactor := gap.x / bB.dir.x
			if bAFactor > 100 || bBFactor > 100 {
				return solution{a: bAFactor, b: bBFactor, score: bAFactor*3 + bBFactor*1}, ErrUnsolvable
			}
			return solution{a: bAFactor, b: bBFactor, score: bAFactor*3 + bBFactor*1}, nil
		}
	}

	// confirm Y solution possible given X
	// if we get to the bottom all options have been tested no solution
	// err = unsolvable
	return solution{}, ErrUnsolvable
}

func part1(in io.Reader) {
	lineScanner := bufio.NewScanner(in)
	var activeButtons = make(map[byte]button, 2)
	sum := 0
	for lineScanner.Scan() {
		line := lineScanner.Text()
		if line == "" {
			continue
		}
		lineReader := strings.NewReader(line)

		switch line[0] {
		case 'B':
			var curButton button
			_, err := fmt.Fscanf(lineReader, "Button %c: X+%d, Y+%d", &curButton.name, &curButton.dir.x, &curButton.dir.y)
			if err != nil {
				panic(err)
			}
			activeButtons[curButton.name] = curButton
			fmt.Printf("B: %s\n", curButton)
		case 'P':
			var curPrize vec
			_, err := fmt.Fscanf(lineReader, "Prize: X=%d, Y=%d", &curPrize.x, &curPrize.y)
			_ = err
			fmt.Printf("P: %+v\n", curPrize)
			solution, err := solve(activeButtons['A'], activeButtons['B'], curPrize)
			if err != nil {
				fmt.Printf("error: %s", err)
			} else {
				fmt.Printf("solution: %+v\n", solution)
				sum += solution.score
			}
			activeButtons = make(map[byte]button)
		}
	}
	fmt.Printf("\nsum : %d\n", sum)
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

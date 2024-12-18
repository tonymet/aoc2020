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
type solutions []solution

func (a solutions) Len() int           { return len(a) }
func (a solutions) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a solutions) Less(i, j int) bool { return a[i].score < a[j].score }

func makeSol(a, b int) solution {
	return solution{a: a, b: b, score: a*3 + b*1}
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
	cur := p
	gap := vec{0, 0}
	solutions := make(solutions, 0, 5)
	for cur.x >= 0 && cur.y >= 0 {
		if (cur.mod2(bA.dir) == vec{0, 0} && gap.mod2(bB.dir) == vec{0, 0}) {
			bAFactor := cur.x / bA.dir.x
			bBFactor := gap.x / bB.dir.x
			if bAFactor <= 100 && bBFactor <= 100 && (bAFactor*bA.dir.y+bBFactor*bB.dir.y) == p.y {
				solutions = append(solutions, makeSol(bAFactor, bBFactor))
			}
		}
		cur = cur.diff(bB.dir)
		gap = p.diff(cur)
	}
	if len(solutions) == 0 {
		return solution{}, ErrUnsolvable
	}

	// sort and return bottom
	sort.Sort(solutions)
	return solutions[0], nil
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

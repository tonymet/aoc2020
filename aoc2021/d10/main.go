package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

type brackets = []rune

type bStack brackets

var bMap = map[rune]rune{
	'{': '}',
	'(': ')',
	'<': '>',
	'[': ']',
}

var part = 1

func score(test rune) int {
	switch test {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		panic("not here")
	}
}

func opener(test rune) bool {
	switch test {
	case '{':
		fallthrough
	case '(':
		fallthrough
	case '<':
		fallthrough
	case '[':
		return true
	default:
		return false
	}
}

func closer(test rune) bool {
	switch test {
	case '}':
		fallthrough
	case ')':
		fallthrough
	case '>':
		fallthrough
	case ']':
		return true
	default:
		return false
	}

}

func (b bStack) completion() (score int) {
	for i := len(b) - 1; i >= 0; i-- {
		score = 5*score + p2Score(b[i])
	}
	return
}

func p2Score(test rune) int {
	switch test {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	default:
		panic("not here")
	}
}
func (b *bStack) push(cur rune) {
	*b = append(*b, cur)
}

func (b *bStack) pop() (rune, error) {
	if len(*b) == 0 {
		return 0, errors.New("empty stack")
	}
	cur := (*b)[len(*b)-1]
	*b = (*b)[0 : len(*b)-1]
	return cur, nil
}

func (b *bStack) checkRune(test rune) bool {
	last, err := (*b).pop()
	if err != nil {
		return false
	}
	switch {
	case test == last:
		// continue
		return true
	case test != last:
		// we're good continue
		return false
	default:
		panic("should not get here")
	}
}

func parseAndSetup() {
	var (
		sum         = 0
		part2scores = make([]int, 0)
	)

file:
	for {
		var line = make(brackets, 0)
		var curStack = make(bStack, 0)
		var mismatch = false
		for {
			var (
				cur rune
			)
			n, err := fmt.Scanf("%c", &cur)
			if err == io.EOF {
				break file
			}
			if n != 1 || cur == '\n' || err != nil {
				// end line
				// not invalid, calculate score
				if !mismatch && part == 2 {
					part2score := curStack.completion()
					fmt.Printf("Part2 Score: %d\n", part2score)
					part2scores = append(part2scores, part2score)
				}
				break
			}
			line = append(line, cur)
			switch {
			case opener(cur):
				curStack.push(bMap[cur])
			case closer(cur):
				if !curStack.checkRune(cur) {
					fmt.Printf("mismatch: %c\n", cur)
					sum += score(cur)
					mismatch = true
					break
				}
			}
		}
		fmt.Printf("line: %v, curStack: %+v\n", string(line), string(curStack))
	}
	if part == 1 {
		fmt.Printf("Score: %d\n", sum)
	}
	if part == 2 {
		// sort and find midpoint
		sort.Ints(part2scores)
		fmt.Printf("part2 midpoint: %d", part2scores[len(part2scores)/2])

	}
}

func part1() {
	parseAndSetup()
	// sort the positions
	// binary search and take deltas
}
func part2() {
	parseAndSetup()
	// sort the positions
	// binary search and take deltas
}

func main() {
	if stdin := os.Getenv("STDIN"); len(stdin) != 0 {
		stdinFile, err := os.Open(stdin)
		if err != nil {
			panic(err)
		}
		os.Stdin = stdinFile

	}
	switch os.Getenv("PART") {
	case "2":
		part = 2
		part2()
	default:
		part1()
	}
}

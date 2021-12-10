package main

import (
	"fmt"
	"os"
	"sort"
)

func sumFromOrigin(positions []int, index int) (sum int) {
	for _, v := range positions {
		switch {
		case v > positions[index]:
			sum += v - positions[index]
		default:
			sum += positions[index] - v
		}
	}
	return
}

func min3(a, b, c int) (min int) {
	min = a
	for _, v := range []int{a, b, c} {
		if v < min {
			min = v
		}

	}
	return
}

func min(set []int) (min int) {
	min = set[0]
	for _, v := range set {
		if v < min {
			min = v
		}
	}
	return

}

func parseAndSetup() {
	var (
		positions = make([]int, 0)
	)

	for {
		var cur int
		n, err := fmt.Scanf("%d", &cur)
		if n != 1 || err != nil {
			break
		}
		positions = append(positions, cur)
	}
	fmt.Printf("len(pos): %d, pos: %+v\n", len(positions), positions)
	sort.Ints(positions)

	// sort the positions
	// binary search and take deltas
	l, r := 0, len(positions)-1
	for {
		ldiff, rdiff := sumFromOrigin(positions, l), sumFromOrigin(positions, r)
		min := min([]int{ldiff, rdiff})
		switch {
		case r-l <= 1:
			fmt.Printf("Pos: %d, solution: %d\n", positions[r], rdiff)
			return
		case ldiff == min:
			// smaller than left, move right
			r = r - ((r - l) / 2)
		case rdiff == min:
			l = ((r - l) / 2) + l
		}
	}
}

func part1() {
	parseAndSetup()
}
func part2() {}

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
		part2()
	default:
		part1()
	}
}

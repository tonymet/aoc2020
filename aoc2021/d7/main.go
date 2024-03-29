package main

import (
	"fmt"
	"os"
	"sort"
)

var (
	positions = make([]int, 0)
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

func funnySumFromPoint(positions []int, point int) (sum int) {
	for _, v := range positions {
		switch {
		case v > point:
			sum += funnyDelta(v - point)
		default:
			sum += funnyDelta(point - v)
		}
	}
	return
}
func funnySumFromOrigin(positions []int, index int) (sum int) {
	for _, v := range positions {
		switch {
		case v > positions[index]:
			sum += funnyDelta(v - positions[index])
		default:
			sum += funnyDelta(positions[index] - v)
		}
	}
	return
}

func funnyDelta(n int) (sum int) {
	for i := 0; i <= n; i++ {
		sum += i
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
}

func part1() {
	parseAndSetup()
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
func part2() {
	parseAndSetup()
	// sort the positions
	// binary search and take deltas
	minSet, maxSet := positions[0], positions[len(positions)-1]
	_, _ = minSet, maxSet
	l, r := minSet, maxSet
	for {
		ldiff, rdiff := funnySumFromPoint(positions, l), funnySumFromPoint(positions, r)
		min := min([]int{ldiff, rdiff})
		switch {
		case r-l <= 1:
			fmt.Printf("Pos: %d, solution: %d\n", r, min)
			return
		case ldiff == min:
			r = r - ((r - l) / 2)
		case rdiff == min:
			l = ((r - l) / 2) + l
		}
	}
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
		part2()
	default:
		part1()
	}
}

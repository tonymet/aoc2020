package main

import (
	"fmt"
	"io"
	"os"
)

const (
	SIZE int = 1000
	//SIZE int = 10
	X = 0
	Y = 1
)

type tGrid = [SIZE][SIZE]int
type tPoint = [2]int

var grid tGrid

func traceGrid(grid *tGrid, p1 tPoint, p2 tPoint) {
	switch {
	case p1[X] == p2[X] && p1[Y] >= p2[Y]:
		for y := p2[Y]; y <= p1[Y]; y++ {
			grid[y][p1[X]]++
		}
	case p1[X] == p2[X] && p1[Y] < p2[Y]:
		for y := p2[Y]; y >= p1[Y]; y-- {
			grid[y][p1[X]]++
		}
		// y case (outer)
	case p1[Y] == p2[Y] && p1[X] <= p2[X]:
		for x := p2[X]; x >= p1[X]; x-- {
			grid[p1[Y]][x]++
		}
	case p1[Y] == p2[Y] && p1[X] > p2[X]:
		for x := p2[X]; x <= p1[X]; x++ {
			grid[p1[Y]][x]++
		}
	case p1[Y] == p2[Y] || p1[X] == p2[X]:
		panic("this shouldn't happen")
	default:
		traceSlope(grid, p2, p1)
		fmt.Printf("not horiz or vertical: %+v, %+v\n", p1, p2)
	}
}

func addPoint(p2, p1 tPoint) tPoint {
	return tPoint{p2[X] + p1[X], p2[Y] + p1[Y]}
}

func eqPoint(p2, p1 tPoint) bool {
	return p2[Y] == p1[Y] && p2[X] == p1[X]
}

func slopeVec(p2, p1 tPoint) (vec tPoint) {
	switch {
	case p2[Y] >= p1[Y]:
		vec[Y] = -1
	case p2[Y] < p1[Y]:
		vec[Y] = 1
	}
	switch {
	case p2[X] >= p1[X]:
		vec[X] = -1
	case p2[X] < p1[X]:
		vec[X] = 1
	}
	return
}

func traceSlope(grid *tGrid, p1 tPoint, p2 tPoint) {
	var (
		slope = slopeVec(p2, p1)
		cur   tPoint
	)
	for cur = (tPoint{p2[X], p2[Y]}); !eqPoint(cur, p1); cur = addPoint(cur, slope) {
		grid[cur[Y]][cur[X]]++
	}
	// set the point itself
	grid[cur[Y]][cur[X]]++

}

func printGrid(grid tGrid) {
	for _, yVal := range grid {
		fmt.Println(yVal)
	}
}

func part1Solution(grid tGrid) (twoCount int) {
	for _, row := range grid {
		for _, cell := range row {
			if cell >= 2 {
				twoCount++
			}
		}
	}
	return
}

func parseAndSetup() {
	var (
		p1, p2 tPoint
	)

	for n, err := fmt.Scanf("%d,%d -> %d,%d\n", &p1[0], &p1[1], &p2[0], &p2[1]); err != io.EOF; n, err = fmt.Scanf("%d,%d -> %d,%d", &p1[0], &p1[1], &p2[0], &p2[1]) {
		if n != 4 {
			fmt.Printf("error reading line\n")
		}

		fmt.Printf("start : %+v, end: %+v \n", p1, p2)
		traceGrid(&grid, p1, p2)
	}
	printGrid(grid)
	fmt.Printf("Part 1 Solution: %d\n", part1Solution(grid))
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

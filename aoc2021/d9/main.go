package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type twod struct {
	x, y int
}

type callback = func(x, y int)

type seenType = map[twod]bool

const (
	//maxX, maxY = 9, 4
	maxX, maxY = 99, 99
)

var (
	grid [maxY + 1][maxX + 1]int
	//mask [maxY + 1][maxX + 1]bool
)

func countBasinFrom(x, y int, seen seenType) int {
	if x < 0 || x > maxX || y < 0 || y > maxY || grid[y][x] == 9 {
		return 0
	}
	seen[twod{x, y}] = true
	basinSum := 1
	up, down, left, right := twod{x, y - 1}, twod{x, y + 1}, twod{x - 1, y}, twod{x + 1, y}
	for _, t := range []twod{up, down, left, right} {
		// don't double back
		if ok := seen[t]; ok {
			continue
		}
		basinSum += countBasinFrom(t.x, t.y, seen)
	}
	return basinSum
}

func part2Sum(basinSizes []int) int {
	// 3 largest basins
	// multiply
	sort.Ints(basinSizes)
	return basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3]

}
func part2() {
	parseAndSetup()
	basinSizes := make([]int, 0)
	// find low points
	// calc basins
	// part2 sum
	iterGrid(func(x, y int) {
		seen := make(seenType)
		basinSizes = append(basinSizes, countBasinFrom(x, y, seen))
	})
	fmt.Printf("part2Sum: %d\n", part2Sum(basinSizes))
}

func isLowest(x, y int) bool {
	cmp := grid[y][x]
	up, down, left, right := twod{x, y - 1}, twod{x, y + 1}, twod{x - 1, y}, twod{x + 1, y}

	for _, t := range []twod{up, down, left, right} {
		// check if out of bounds
		// check if lower
		if t.x < 0 || t.x > maxX || t.y < 0 || t.y > maxY {
			continue
		}
		if grid[t.y][t.x] <= cmp {
			return false
		}
	}
	return true
}

func parseAndSetup() {
	var cur int
	x, y := 0, 0

file:
	for {
		n, err := fmt.Scanf("%1d", &cur)
		switch {
		case err == io.EOF:
			break file
		case n == 0:
			continue file
		}
		grid[y][x] = cur
		x = (x + 1) % (maxX + 1)
		if x == 0 {
			y++
		}
	}

}

// iterate over grid and yield to callback if isLowest() == true
func iterGrid(cb callback) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			// test and add sum
			if isLowest(x, y) {
				cb(x, y)
			}
		}
	}
}
func part1() {
	parseAndSetup()
	var part1Sum int
	// iterate over both
	iterGrid(func(x, y int) {
		fmt.Printf("lowest  (x, y ,val): (%d, %d, %d)\n", x, y, grid[y][x])
		part1Sum += grid[y][x] + 1
	})
	fmt.Printf("\nisLower (1,0) %+x\n", isLowest(1, 0))
	fmt.Printf("\nisLower (9,0) %+x\n", isLowest(9, 0))
	fmt.Printf("part1Sum: %d\n", part1Sum)
	// iterate over grid ,  check borders, add to maks
	//fmt.Printf("%+x\n", grid)
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

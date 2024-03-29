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
	maxX, maxY = 9, 9
	//maxX, maxY = 99, 99
)

var (
	grid [maxY + 1][maxX + 1]int
	//mask [maxY + 1][maxX + 1]bool
)

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
		//seen := make(seenType)
		//basinSizes = append(basinSizes, countBasinFrom(x, y, seen))
	})
	fmt.Printf("part2Sum: %d\n", part2Sum(basinSizes))
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
			y++
			x = 0
			continue file
		}
		grid[y][x] = cur
		x++
	}
}

// iterate over grid and yield to callback if isLowest() == true
func iterGrid(cb callback) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			// test and add sum
			cb(x, y)
		}
	}
}
func part1() {
	parseAndSetup()
	//var part1Sum int
	// iterate over both
	iterGrid(func(x, y int) {
		//fmt.Printf("lowest  (x, y ,val): (%d, %d, %d)\n", x, y, grid[y][x])
		//part1Sum += grid[y][x] + 1
	})
	fmt.Printf("%+v\n", grid)
	step()
	fmt.Printf("%+v\n", grid)

}

func increaseNeighbors(pos twod) {
	/*
		up, down, left, right, ul, ur, bl, br := twod{x, y - 1}, twod{x, y + 1}, twod{x - 1, y},
			twod{x + 1, y}, twod{x - 1, y - 1}, twod{x + 1, y - 1}, twod{x - 1, y + 1}, twod{x + 1, y + 1}
	*/
	/*
		for _, v := range []twod{up, down, left, right, ul, ur, bl, br} {
			grid[v.y][v.x] += 1
		}
	*/
	for y := pos.y - 1; y <= pos.y+1; y++ {
		for x := pos.x - 1; x <= pos.x+1; x++ {
			// skip cur
			if pos == (twod{x, y}) || y < 0 || x < 0 || y > maxY || x > maxX {
				continue
			}
			grid[y][x]++
		}
	}
}

func step() {
	// iterate over grid
	// if < 9, increase by 1
	// if == 9 , increase neighbors by 1, then set self to zero
	// if zero , stop

	// increaseO
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			grid[y][x] += 1
			if grid[y][x] >= 9 && grid[y][x] != 0 {
				increaseNeighbors(twod{x, y})
				grid[y][x] = 0
			}
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

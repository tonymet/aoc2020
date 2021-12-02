package main

import (
	"flag"
	"fmt"
	"io"
)

var grid []string

type d2 struct {
	x, y int
}

var lineWidth = 32

func readGrid() {
	grid = make([]string, 0)
	var (
		y   int
		cur string
	)
	for {
		_, err := fmt.Scanf("%32s", &cur)
		if err == io.EOF {
			break
		}
		grid = append(grid, cur)
		y++
	}
	lineWidth = len(grid[0])
	fmt.Printf("lineWidth = %d\n", lineWidth)
}

func scanGrid(slope d2) int64 {
	var (
		spot  d2 = d2{0, 0}
		count int64
	)
	for spot.y < len(grid) {
		//fmt.Printf("line: %d ; spot: %+v : %s\n", spot.y, spot, string(grid[spot.y][spot.x]))
		if grid[spot.y][spot.x] == '#' {
			count++
		}
		spot.x = (spot.x + slope.x) % (lineWidth)
		spot.y += slope.y
	}
	return count
}

func main() {
	readGrid()
	pFlag := flag.Int("p", 1, "p 1 or 2")
	flag.Parse()
	switch *pFlag {
	case 1:
		count := scanGrid(d2{3, 1})
		fmt.Printf("\nGrid = %+v", grid)
		fmt.Printf("\nCount = %d", count)
	case 2:
		var product int64 = 1
		slopes := []d2{
			{1, 1},
			{3, 1},
			{5, 1},
			{7, 1},
			{1, 2},
		}
		for _, s := range slopes {
			count := scanGrid(s)
			product *= count
			fmt.Printf("\nCount = %d", count)
		}
		fmt.Printf("\nproduct = %d", product)
	default:
		panic("-p wrong")
	}
}

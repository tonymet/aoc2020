package main

import (
	"fmt"
)

var grid []string

type d2 struct {
	x, y int
}

const (
	lineWidth = 32
	x         = 3
	y         = 1
)

var count = 0

func readGrid() {
	grid = make([]string, 0)
	var (
		y   int
		cur string
	)
	for {
		_, err := fmt.Scanf("%32s", &cur)
		grid = append(grid, cur)
		if err != nil {
			break
		}
		y++
	}
}

func scanGrid() {
	var spot d2
	spot.x, spot.y = 0, 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < lineWidth; x++ {
			if spot.x == x && spot.y == y {
				fmt.Print(string(grid[y][x]))
				spot.x = (spot.x + 3) % (lineWidth - 1)
				spot.y++
				if string(grid[y][x]) == "#" {
					count++
				}
			}
		}
	}

}

func main() {
	readGrid()
	scanGrid()
	fmt.Printf("\nCount = %d", count)
}

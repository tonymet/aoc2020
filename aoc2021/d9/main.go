package main

import (
	"fmt"
	"io"
	"os"
)

func slidingSum(page []int) int {
	if page[3] > page[0] {
		return 1
	} else if page[0] > page[3] {
		return -1
	}
	return 1
}

func part2() {
}
func part1() {
	var cur int
	var grid [5][10]int
	x, y := 0, 0

	for {
		n, err := fmt.Scanf("%1d", &cur)
		fmt.Printf("%+v", cur)
		if err == io.EOF {
			break
		}
		if n == 0 {
			continue
		}
		grid[y][x] = cur
		x = (x + 1) % 10
		if x == 0 {
			y++
		}
	}
	fmt.Printf("%+x\n", grid)
}
func main() {
	switch os.Getenv("PART") {
	case "2":
		part2()
	default:
		part1()
	}
}

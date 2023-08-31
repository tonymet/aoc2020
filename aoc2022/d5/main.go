package main

import (
	"fmt"
	"io"
	"os"
	"slices"
)

// bounds for the tower b/c we are streaming the file
const (
	WIDTH  int = 3
	HEIGHT int = 3
)

// utilities for parsing the tower into data structures
func keep(index int) bool {
	return ((index - 1) % 4) == 0
}

func colToTower(col int) int {
	return col / 4
}
func isLetter(l byte) bool {
	return l >= 65 && l <= 90
}

func part2() {
}
func part1() {
	var b []byte = make([]byte, 1)
	var towers [][]byte = make([][]byte, WIDTH)
	for i := 0; i < WIDTH; i++ {
		towers[i] = make([]byte, 0, HEIGHT)
	}
	var row, col int
	for {
		_, err := os.Stdin.Read(b)
		if err == io.EOF {
			break
		}
		// keep will skip over benign characters
		if keep(col) {
			fmt.Printf("b:%s, row: %d, col %d, toTower: %d\n", string(b), row, col, colToTower(col))
			if isLetter(b[0]) {
				towers[colToTower(col)] = append(towers[colToTower(col)], b[0])
			}
		}
		// end of tower reading
		if col == 0 && b[0] == '\n' {
			break
		}
		// update row , col values
		col++
		if b[0] == '\n' {
			row++
			col = 0
		}
	}
	for i := range towers {
		slices.Reverse(towers[i])
	}
	fmt.Printf("towers: %+v\n", towers)

	var moves = make([][3]int, 0)
	var cur [3]int
	for {
		_, err := fmt.Scanf("move %d from %d to %d\n", &cur[0], &cur[1], &cur[2])
		if err != nil {
			break
		}
		fmt.Printf("curMove: %+v\n", cur)
		moves = append(moves, cur)
	}
	fmt.Printf("moves: %+v\n", moves)
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

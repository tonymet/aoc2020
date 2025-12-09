package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	_ "sort"
)

type fileParam struct {
	cols, rows int
}

const (
	target = '@'
)

var (
	files map[string]fileParam = map[string]fileParam{
		"sample": {cols: 10, rows: 10},
		"input":  {cols: 140, rows: 140},
	}
	filetype string
	newline  = "\r\n"
	offsets  = [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
)

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

func part1(in io.Reader) {
	R, C := files[filetype].rows, files[filetype].cols
	grid := make([][]byte, 0, C)
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}
	//fmt.Printf("%s\n", grid)
	totalCount := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			tCount := 0
			if grid[r][c] != target {
				continue
			}
			for _, offset := range offsets {
				dr, dc := offset[0], offset[1]
				nr, nc := r+dr, c+dc
				if nr >= 0 && nr < R && nc >= 0 && nc < C {
					if grid[nr][nc] == target {
						tCount++
					}
				}
			}
			if tCount < 4 {
				totalCount++
			}
		}
	}
	fmt.Printf("totalCount: %d\n", totalCount)
}

var (
	part   int
	file   string
	silent bool
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")
	flag.StringVar(&filetype, "type", "", "filetype")
	flag.BoolVar(&silent, "s", false, "silent?")

}

func main() {
	flag.Parse()
	if file != "" {
		var err error
		if os.Stdin, err = os.Open(file); err != nil {
			panic(err)
		}
	}
	switch part {
	case 2:
		part2(os.Stdin)
	default:
		part1(os.Stdin)
	}
}

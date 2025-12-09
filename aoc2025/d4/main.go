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

var (
	files map[string]fileParam = map[string]fileParam{
		"sample": {cols: 10, rows: 10},
		"input":  {cols: 100, rows: 100},
	}
	filetype string
	newline  = "\r\n"
)

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

func part1(in io.Reader) {
	rows, cols := files[filetype].rows, files[filetype].cols
	mapData := make([][]byte, 0, cols)
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		mapData = append(mapData, []byte(line))
	}
	fmt.Printf("%s\n", mapData)
	_ = rows

	// for r :=0 ; r < rows; r++{
	// 	for c:=0; c < cols; c++{
	// 		tCount := 0

	// 	}
	// }
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

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	_ "sort"
)

var (
	filetypes = map[string]fileparam{
		"sample.txt": {8, 2},
		"input.txt":  {496, 2},
	}
	activeParam fileparam
)

type pt struct {
	x, y int
}

type fileparam struct {
	rows, cols int
}

func area(p1, p2 pt) (area int) {
	dx, dy := p1.x-p2.x+1, p1.y-p2.y+1
	a := dx * dy
	if a < 0 {
		return -1 * a
	}
	return a
}

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

func part1(in io.Reader) {
	pts := make([]pt, 0, activeParam.rows)
	for {
		var cur pt
		_, err := fmt.Fscanf(in, "%d,%d", &cur.x, &cur.y)
		if err == io.EOF {
			break
		}
		pts = append(pts, cur)
		//fmt.Printf("%d,%d\n", cur.x, cur.y)
	}
	maxA := 0
	for i := 0; i < len(pts); i++ {
		for j := 0; j < len(pts); j++ {
			if i == j {
				continue
			}
			a := area(pts[i], pts[j])
			if a > maxA {
				maxA = a
			}
			//fmt.Printf("i: %+x, j: %+x, a: %d\t", pts[i], pts[j], a)
		}
	}
	fmt.Printf("\nmaxA: %d\n", maxA)
}

var (
	part   int
	file   string
	silent bool
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")
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
	activeParam = filetypes[path.Base(file)]
	switch part {
	case 2:
		part2(os.Stdin)
	default:
		part1(os.Stdin)
	}
}

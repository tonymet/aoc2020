package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	_ "sort"
	"strconv"
)

func part2() {
	fmt.Printf("part2 not implemented\n")
}

const (
	MAX_Y = 140
	MAX_X = 140
	// MAX_X = 10
	// MAX_Y = 10
)

type matrix [MAX_Y][MAX_X]byte
type vec struct {
	x, y int
}

func part1() {
	var mtx matrix
	var symbolList = make([]vec, 0)
	var sumList = make([]int64, 0)

	var (
		x int = 0
		y int = 0
	)

	for {
		var curChar byte
		_, err := fmt.Scanf("%c", &curChar)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("%c", curChar)
		if curChar == '\n' {
			y += 1
			x = 0
			continue
		} else {
			mtx[y][x] = curChar
			x += 1
		}
	}
	// find special chars
	for yidx := 0; yidx < MAX_Y; yidx += 1 {
		buf := 0
		for xidx := 0; xidx < MAX_X; xidx += 1 {
			cur := mtx[yidx][xidx]
			if buf > 0 {
				buf -= 1
				continue
			}
			switch {
			case cur >= '0' && cur <= '9':
				// skip
				val, l, err := readNum(&mtx, yidx, xidx)
				// check and add
				if err != nil {
					continue
				} else {
					buf = l
					if symbolCheck(&mtx, vec{xidx, yidx}, l) {
						sumList = append(sumList, val)
					}
				}
				fmt.Printf("val %+v, l%+v\n", val, l)
			case cur == '.':
				buf = 0
			default:
				// special char
				symbolList = append(symbolList, vec{xidx, yidx})
				buf = 0
			}
		}
	}
	// find adjascent part nums and sum
	// fmt.Printf("mtx: %+v\n", mtx)
	fmt.Printf("symbolList %+v\n", symbolList)
	fmt.Printf("len symbolList %+d\n", len(symbolList))
	fmt.Printf("sumList %+v\n", sumList)
	fmt.Printf("len sumList %d\n", len(sumList))
	var sum int64 = 0
	for _, val := range sumList {
		sum += val
	}
	fmt.Printf("sum %+v\n", sum)
}

func symbolCheck(mtx *matrix, p vec, l int) bool {
	for yidx := p.y - 1; yidx < MAX_Y && yidx <= p.y+1; yidx += 1 {
		if yidx < 0 {
			continue
		}
		for xidx := p.x - 1; xidx < MAX_X && xidx < p.x+l+1; xidx += 1 {
			if xidx < 0 {
				continue
			}
			cur := mtx[yidx][xidx]
			switch {
			case cur >= '0' && cur <= '9':
				continue
			case cur == '.':
				continue
			default:
				return true
			}

		}
	}
	return false
}

func readNum(mtx *matrix, y, x int) (int64, int, error) {
	slice := mtx[y][x:MAX_X]
	bound := 0
	for i, c := range slice {
		if !(c >= '0' && c <= '9') {
			break
		}
		bound = i
	}
	// parse int slice
	if intVal, err := strconv.ParseInt(string(slice[0:bound+1]), 10, 64); err != nil {
		return 0, 0, err
	} else {
		return intVal, bound + 1, nil
	}
}

var (
	part int
	file string
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")
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
		part2()
	default:
		part1()
	}
}

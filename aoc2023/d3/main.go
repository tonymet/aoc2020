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
	// part 1 to parse and build matrix
	part1()
	// loop symbols where = *
	// find adjacent numbers
	// if len = 2 multiply
	// sum it up
	sum := int64(0)
	for _, sym := range symbolList {
		if mtx[sym.y][sym.x] != '*' {
			continue
		}
		ret, val := findNums(vec{sym.x, sym.y})
		if len(ret) == 2 {
			fmt.Printf("bounds of *: %+v, mult %d\n", ret, val)
			sum += val
		}
	}
	fmt.Printf("summysum: %d\n", sum)

}

func isNum(c byte) bool {
	return c >= '0' && c <= '9'
}

func findNums(p vec) ([]int64, int64) {
	var ret = make([]int64, 0, 9)
	var valMap = make(map[int64]bool)
	for yidx := p.y - 1; yidx < maxDim && yidx <= p.y+1; yidx += 1 {
		if yidx < 0 {
			continue
		}
		for xidx := p.x - 1; xidx < maxDim && xidx <= p.x+1; xidx += 1 {
			if xidx < 0 {
				continue
			}
			if valMtx[yidx][xidx] > 0 {
				valMap[valMtx[yidx][xidx]] = true
			}
		}
	}
	//return ret
	for k := range valMap {
		ret = append(ret, k)
	}
	var sum int64 = ret[0]
	for _, v := range ret[1:] {
		sum = sum * v
	}
	return ret, sum
}

const (
	MAX_S = 10
	MAX_1 = 140
)

var (
	maxDim     = MAX_S
	mtx        matrix
	symbolList []vec
	valMtx     [][]int64
	sumList    []int64
)

type matrix [][]byte
type vec struct {
	x, y int
}

func part1() {
	mtx = make(matrix, maxDim)
	for i := 0; i < maxDim; i++ {
		mtx[i] = make([]byte, maxDim)
	}
	symbolList = make([]vec, 0)
	sumList = make([]int64, 0)
	valMtx = make([][]int64, maxDim)
	for i := 0; i < maxDim; i++ {
		valMtx[i] = make([]int64, maxDim)
	}

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
	for yidx := 0; yidx < maxDim; yidx += 1 {
		buf := 0
		for xidx := 0; xidx < maxDim; xidx += 1 {
			cur := mtx[yidx][xidx]
			if buf > 0 {
				buf -= 1
				continue
			}
			switch {
			case isNum(cur):
				// skip
				val, l, err := readNum(&mtx, yidx, xidx)
				// check and add
				if err != nil {
					continue
				} else {
					buf = l - 1
					for i := xidx; i < xidx+l; i++ {
						valMtx[yidx][i] = val
					}
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
	for yidx := p.y - 1; yidx < maxDim && yidx <= p.y+1; yidx += 1 {
		if yidx < 0 {
			continue
		}
		for xidx := p.x - 1; xidx < maxDim && xidx < p.x+l+1; xidx += 1 {
			if xidx < 0 {
				continue
			}
			cur := (*mtx)[yidx][xidx]
			switch {
			case isNum(cur):
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
	slice := (*mtx)[y][x:maxDim]
	bound := 0
	for i, c := range slice {
		if !(isNum(c)) {
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
	part   int
	file   string
	format string
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")
	flag.StringVar(&format, "format", "sample", "sample|full")

}
func main() {
	flag.Parse()
	if file != "" {
		var err error
		if os.Stdin, err = os.Open(file); err != nil {
			panic(err)
		}
	}

	switch format {
	case "sample":
		maxDim = 10
	default:
		maxDim = 140

	}
	switch part {
	case 2:
		part2()
	default:
		part1()
	}
}

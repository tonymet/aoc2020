package main

import (
	"flag"
	"fmt"
	"io"
	"sort"
)

func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func part2() {
	var l, r int
	lcol, rmap := make([]int, 0), make(map[int]int, 0)

	for {
		_, err := fmt.Scanf("%d %d", &l, &r)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		lcol = append(lcol, l)
		rmap[r] += 1
	}
	// calculate lookup
	sum := 0
	for _, e := range lcol {
		v, ok := rmap[e]
		if !ok {
			v = 0
		}
		sum += e * v
	}
	fmt.Printf("sum: %d\n", sum)
}

func part1() {
	var l, r int
	lcol, rcol := make([]int, 0), make([]int, 0)
	for {
		_, err := fmt.Scanf("%d %d", &l, &r)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		lcol = append(lcol, l)
		rcol = append(rcol, r)
	}
	// sort both
	sort.Ints(lcol)
	sort.Ints(rcol)
	sum := 0
	for i := range lcol {
		sum += Abs(lcol[i] - rcol[i])
	}
	fmt.Printf("sum: %d\n", sum)
}

var (
	part int
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
}
func main() {
	flag.Parse()
	switch part {
	case 2:
		part2()
	default:
		part1()
	}
}

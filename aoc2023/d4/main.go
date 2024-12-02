package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	_ "sort"
)

func part2() {
	fmt.Printf("part2 not implemented\n")
}

type bound struct {
	min, max int64
}
type vec struct {
	x, y, z int64
}
type vecF struct {
	x, y float64
}
type rec struct {
	p, v vec
}

func part1() {
	for {
		var curRec rec
		_, err := fmt.Scanf("%d, %d, %d @ %d, %d, %d", &curRec.p.x, &curRec.p.y, &curRec.p.z, &curRec.v.x, &curRec.v.y, &curRec.v.z)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("rec: %+v\n", curRec)
	}
}

func parseRec(in io.Reader) (curRec rec) {
	_, err := fmt.Fscanf(in, "%d, %d, %d @ %d, %d, %d", &curRec.p.x, &curRec.p.y, &curRec.p.z, &curRec.v.x, &curRec.v.y, &curRec.v.z)
	if err == io.EOF {
		return
	} else if err != nil {
		panic(err)
	}

	return
}

var (
	part     int
	file     string
	boundary bound
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")

}

func (r rec) solveY(x float64) vecF {
	slope := float64(r.v.y) / float64(r.v.x)
	diffX := float64(r.p.x) - x
	return vecF{x, float64(r.p.y) - (slope * diffX)}
}

/*
func findCrossing(xb, yb bound, r1, r2 rec) {
	// bisect xb.min to xb.max , see if they cross before one is out of bound
	x := xb.min
	d := (r1.solve)
}
*/

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
	case 0:
		boundary = bound{7, 27}
		part1()
	default:
		boundary = bound{200000000000000, 400000000000000}
		part1()
	}
}

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

type vec struct {
	x, y, z int64
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

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	_ "sort"
	"strconv"
)

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

type stonesType []int64

// If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
// If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
// If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
func digLen(s int64) int {
	return len(strconv.FormatInt(s, 10))
}
func splitStone(s int64) []int64 {
	toString := strconv.FormatInt(s, 10)
	strLen := len(toString)
	l, r := toString[:strLen-1], toString[strLen-1:]
	lInt, err := strconv.ParseInt(l, 10, 64)
	if err != nil {
		panic(err)
	}
	rInt, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		panic(err)
	}
	ret := make([]int64, 2)
	ret[0], ret[1] = lInt, rInt
	return ret
}

func blink(in stonesType) stonesType {
	// rules
	ret := make(stonesType, 0, len(in)*5)
	for _, s := range in {
		switch {
		case s == 0:
			ret = append(ret, 1)
		case digLen(s)%2 == 0:
			ret = append(ret, splitStone(s)...)
		default:
			ret = append(ret, s*2024)
		}

	}
	return ret
}

func part1(in io.Reader) {
	fmt.Printf("part1 not implemented\n")
	initStones := make(stonesType, 0, 5)
	for {
		var cur int64
		_, err := fmt.Fscan(in, &cur)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		initStones = append(initStones, cur)
	}
	fmt.Printf("initStones: %+v", initStones)
	var r = initStones
	for range 25 {
		r := blink(r)
		fmt.Printf("r: %+v", r)
	}
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
	switch part {
	case 2:
		part2(os.Stdin)
	default:
		part1(os.Stdin)
	}
}

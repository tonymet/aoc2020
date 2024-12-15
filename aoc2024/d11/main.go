package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/tonymet/aoc2020/shared"
)

type stonesType []int64
type indexType map[int64]int64

func mergeIndexFactor(l, r indexType, f int64) {
	for k := range r {
		l[k] += r[k] * f
	}
}

func sumValues(s indexType) (sum int64) {
	for _, v := range s {
		sum += v
	}
	return
}

func splitInteger(num int64) (left, right int64, digits int) {
	temp := num
	for temp > 0 {
		digits++
		temp /= 10
	}
	divisor := int64(shared.Pow10(digits / 2))
	left = num / divisor
	right = num % divisor
	return
}

func index(stones stonesType) indexType {
	index := make(indexType)
	for _, v := range stones {
		index[v] += 1
	}
	return index
}
func blink(in stonesType) stonesType {
	ret := make(stonesType, 0, len(in)*2)
	for _, s := range in {
		if s == 0 {
			ret = append(ret, 1)
		} else if l, r, digits := splitInteger(s); digits%2 == 0 {
			ret = append(ret, l, r)
		} else {
			ret = append(ret, s*2024)
		}
	}
	return ret
}

func parseStones(in io.Reader) stonesType {
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
	return initStones
}

func stoners(in io.Reader) {
	initStones := parseStones(in)
	fmt.Printf("initStones: %+v\n", initStones)
	stoneIndex := index(initStones)
	var cap = [2]int64{5, 5}
	switch part {
	case 1:
	case 2:
		cap = [2]int64{5, 15}
	default:
		panic("no cap")
	}
	for range cap[0] {
		curIndex := make(indexType)
		for k, v := range stoneIndex {
			cur := stonesType{k}
			for range cap[1] {
				cur = blink(cur)
			}
			mergeIndexFactor(curIndex, index(cur), v)
		}
		stoneIndex = curIndex
		fmt.Printf("curIndex sum: %d\n", sumValues(curIndex))
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
	stoners(os.Stdin)
}

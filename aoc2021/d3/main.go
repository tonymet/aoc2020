package main

import (
	"fmt"
	"io"
	"os"
)

type threed struct {
	x, y, z int
}

type bitTally map[int]int

//var masks = []int{0b01, 0b10, 0b100, 0b1000, 0b10000}
//var masks = []int{0b01, 0b10, 0b100, 0b1000, 0b10000, 0b1000000, 0b10000000,
//0b100000000, 0b1000000000, 0b10000000000, 0b100000000000, 0b1000000000000}

const width = 12

const invertMask = 0b111111111111

//const invertMask = 0b11111

var masks = genMasks(width)
var maskCounters = make(bitTally)

func xorAll(vals []int) (sum int) {
	sum = vals[0]
	for _, val := range vals[1:] {
		sum = sum ^ val
	}
	return sum
}

func genMasks(width int) []int {
	masks := make([]int, 0, width)
	for i := 0; i < width; i++ {
		masks = append(masks, 0b1<<i)
	}
	return masks
}

func tallyBits(val int) error {
	for _, mask := range masks {
		if (mask & val) == mask {
			maskCounters[mask]++
		}
	}
	return nil
}

func tallyToInt(target int) (val int) {
	mask := masks[len(masks)-1]
	for i := len(masks) - 1; i >= 0; i-- {
		fmt.Printf("mask: %d , val: %d\n", masks[i], maskCounters[masks[i]])

		if maskCounters[masks[i]] > target {
			val = mask | val
		}
		mask = mask >> 1
	}
	return val
}

func part1() {
	var (
		val  int
		vals []int
		i    int
	)
	vals = make([]int, 0, 100)

	for n, err := fmt.Scanf("%b", &val); err != io.EOF; n, err = fmt.Scanf("%b", &val) {
		if n != 1 {
			fmt.Printf("error reading line\n")
		}
		vals = append(vals, val)
		tallyBits(val)
		fmt.Printf("value : %d, binary: %b \n", val, val)
		i++
	}
	fmt.Printf("vals : %+v, sum: %+v\n", vals, xorAll(vals))
	fmt.Printf("maskCounters : %+v\n", maskCounters)
	fmt.Printf("masks : %+v\n", masks)
	fmt.Printf("tally : %+v\n", tallyToInt(6))
	fmt.Printf("len : %+v\n", i)

	tally := tallyToInt(i / 2)
	fmt.Printf("invert tally : %+v\n", invertMask^tally)
	fmt.Printf("solution:: %+v\n", tally*(invertMask^tally))

}

func part2() {}
func main() {
	switch os.Getenv("PART") {
	case "2":
		part2()
	default:
		part1()
	}
}

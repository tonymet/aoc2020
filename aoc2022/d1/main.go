package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type ElfTally struct {
	tallyList [][]int
	counts    map[int]int
}

func maxCounts(c map[int]int) int {
	var max int
	for _, v := range c {
		if v >= max {
			max = v
		}

	}
	return max
}

func maxCounts3(c map[int]int) []int {
	var allCounts []int = make([]int, 0, len(c))

	for _, v := range c {
		allCounts = append(allCounts, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(allCounts)))
	return []int{allCounts[0], allCounts[1], allCounts[2]}
}

func sum(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return sum
}

func part2() {

}
func part1() {
	var tl ElfTally
	tl.tallyList = make([][]int, 1)
	tl.counts = make(map[int]int)
	tl.tallyList[0] = make([]int, 0)
	var cur, elf int
	for _, err := fmt.Scanf("%d\n", &cur); err != io.EOF; _, err = fmt.Scanf("%d\n", &cur) {
		if err != nil {
			fmt.Printf("new elf.  prev, elf = %+v\n", tl.tallyList[elf])
			elf++
			tl.tallyList = append(tl.tallyList, make([]int, 0))
			continue
		}
		fmt.Printf("cur: %d, \n", cur)
		tl.tallyList[elf] = append(tl.tallyList[elf], cur)
		tl.counts[elf] += cur
	}
	fmt.Printf("TallyList : %+v\n", tl)
	fmt.Printf("max : %+v\n", maxCounts(tl.counts))
	max3 := maxCounts3(tl.counts)
	fmt.Printf("top3 : %+v\n", max3)
	fmt.Printf("sum : %+v\n", sum(max3))

}
func main() {
	switch os.Getenv("PART") {
	case "2":
		part2()
	default:
		part1()
	}
}

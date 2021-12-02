package main

import (
	"fmt"
	"io"
	"os"
)

func slidingSum(page []int) int {
	if (page[0] + page[1] + page[2]) > (page[1] + page[2] + page[3]) {
		return -1
	} else if (page[0] + page[1] + page[2]) < (page[1] + page[2] + page[3]) {
		return 1
	}
	return 0
}

func part2() {
	var (
		page    []int
		i       int
		cur     int
		greater int
	)
	page = make([]int, 4, 4)
	for _, err := fmt.Scanf("%d\n", &cur); err != io.EOF; _, err = fmt.Scanf("%d\n", &cur) {
		page = append(page[1:4], cur)
		fmt.Printf("page: %+v", &page)
		if i > 2 {
			switch slidingSum(page) {
			case -1:
				fmt.Printf("decreasing: \n")
			case 1:
				fmt.Printf("increasing: \n")
				greater++
			case 0:
				fmt.Printf("same: \n")
			}
		}
		i++
	}
	fmt.Printf("greater: %d\n", greater)

}
func part1() {
	var cur, prev, greater int
	for _, err := fmt.Scanf("%d\n", &cur); err != io.EOF; _, err = fmt.Scanf("%d\n", &cur) {
		if cur > prev && prev != 0 {
			greater++
		}
		fmt.Printf("cur: %d, prev: %d, greater: %d\n", cur, prev, greater)
		prev = cur
	}
}
func main() {
	switch os.Getenv("PART") {
	case "2":
		part2()
	default:
		part1()
	}
}

package main

import (
	"fmt"
	"io"
	"os"
)

type bounds [4]int

func part1() {

	var total int
	for {
		var cur bounds
		_, err := fmt.Scanf("%d-%d,%d-%d\n", &cur[0], &cur[1], &cur[2], &cur[3])
		if err == io.EOF {
			break
		}
		if contains(cur) {
			total++
		}
		fmt.Printf("cur: %+v, contains: %+v\n", cur, contains(cur))
	}
	fmt.Printf("total: %d\n", total)
}

func contains(c bounds) bool {
	if c[0] <= c[2] && c[1] >= c[3] {
		return true
	}
	if c[2] <= c[0] && c[3] >= c[1] {
		return true
	}
	return false
}

func part2() {}

func main() {
	if stdin := os.Getenv("STDIN"); len(stdin) != 0 {
		stdinFile, err := os.Open(stdin)
		if err != nil {
			panic(err)
		}
		os.Stdin = stdinFile

	}
	switch os.Getenv("PART") {
	case "2":
		part2()
	default:
		part1()
	}
}

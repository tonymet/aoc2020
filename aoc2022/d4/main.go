package main

import (
	"fmt"
	"io"
	"os"
)

type bounds [4]int

var part int

func run() {
	var total int
	var containsFunc func(bounds) bool
	switch part {
	case 2:
		containsFunc = contains2
	default:
		containsFunc = contains
	}
	for {
		var cur bounds
		_, err := fmt.Scanf("%d-%d,%d-%d\n", &cur[0], &cur[1], &cur[2], &cur[3])
		if err == io.EOF {
			break
		}
		if containsFunc(cur) {
			total++
		}
		fmt.Printf("cur: %+v, contains: %+v\n", cur, containsFunc(cur))
	}
	fmt.Printf("total part %d: %d\n", part, total)
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

func contains2(c bounds) bool {
	// equal & overlap
	if c[0] == c[2] && c[1] == c[3] {
		return true
	}
	// find left
	var l, r [2]int
	if c[0] <= c[2] {
		l[0], l[1] = c[0], c[1]
		r[0], r[1] = c[2], c[3]
	} else {
		r[0], r[1] = c[0], c[1]
		l[0], l[1] = c[2], c[3]
	}

	// if right side of the left > left side of the right
	if l[1] >= r[0] {
		return true
	}
	return false
}

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
		part = 2
		run()
	default:
		part = 1
		run()
	}
}

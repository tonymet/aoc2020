package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	_ "sort"
)

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

type rec []int

func log(f string, val ...any) {
	if silent {
		return
	}
	fmt.Printf(f, val...)
}

func part2() {
	fmt.Printf("part2 not implemented\n")
}

func part1(in io.ReadSeeker) {
	scanner := bufio.NewScanner(in)
	patternSplit := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		loc := mulPattern.FindIndex(data)
		if loc == nil {
			if !atEOF {
				return 0, nil, nil
			}
			// If we have reached the end, return the last token.
			return 0, nil, bufio.ErrFinalToken
		}
		// Otherwise, return the token before the comma.
		return loc[1], data[loc[0]:loc[1]], nil
	}
	scanner.Split(patternSplit)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("token: %s\n", text)
	}
}

var (
	part       int
	file       string
	silent     bool
	mulPattern = regexp.MustCompile(`(?m)mul\(\d{1,5},\d{1,5}\)`)
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
		part2()
	default:
		part1(os.Stdin)
	}
}

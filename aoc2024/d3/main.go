package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	_ "sort"

	"github.com/tonymet/aoc2020/shared"
)

func log(f string, val ...any) {
	if silent {
		return
	}
	fmt.Printf(f, val...)
}

func part2(in io.Reader) {
	sum := int64(0)
	scanner := bufio.NewScanner(in)
	patternSplit := shared.SplitterFactory(mulPattern2)
	scanner.Split(patternSplit)
	do := true
	for scanner.Scan() {
		text := scanner.Text()
		switch text {
		case "do()":
			do = true
		case "don't()":
			do = false
		default:
			v := execMul(text)
			log("token: %s, v= %d\n", text, v)
			if do {
				sum += v
			}
		}
	}
	log("sum: %d\n", sum)
}

func part1(in io.ReadSeeker) {
	sum := int64(0)
	scanner := bufio.NewScanner(in)
	patternSplit := shared.SplitterFactory(mulPattern)
	scanner.Split(patternSplit)
	for scanner.Scan() {
		text := scanner.Text()
		v := execMul(text)
		log("token: %s, v= %d\n", text, v)
		sum += v
	}
	log("sum: %d\n", sum)
}

func execMul(x string) int64 {
	var l, r int64
	if _, err := fmt.Sscanf(x, "mul(%d,%d)", &l, &r); err != nil {
		return 0
	}
	return l * r
}

var (
	part        int
	file        string
	silent      bool
	mulPattern  = regexp.MustCompile(`(?m)mul\(\d{1,5},\d{1,5}\)`)
	mulPattern2 = regexp.MustCompile(`(?m)do\(\)|don't\(\)|mul\(\d+,\d+\)`)
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

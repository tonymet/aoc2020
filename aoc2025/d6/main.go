package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	_ "sort"
	"strings"
)

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

type IntScanner struct {
	bufio.Scanner
}

type fileParam struct {
	col, row int
}

func (is *IntScanner) ScanInts(recLen int) []int {
	t := is.Text()
	reader := strings.NewReader(t)
	ints := make([]int, 0, recLen)
	for {
		var cur int
		_, err := fmt.Fscanf(reader, "%d", &cur)
		if err == io.EOF {
			break
		}
		ints = append(ints, cur)
	}
	return ints
}

func (is *IntScanner) ScanOp(recLen int) []rune {
	t := is.Text()
	reader := strings.NewReader(t)
	runes := make([]rune, 0, recLen)
	for {
		var cur rune
		_, err := fmt.Fscanf(reader, "%c", &cur)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		switch cur {
		case '+':
			fallthrough
		case '*':
			runes = append(runes, cur)
		}
	}
	return runes
}

func part1(in io.Reader) {
	var scanner IntScanner
	scanner.Scanner = *bufio.NewScanner(in)
	i := 0
	state := stateScanNum
	fileSpec := files[fileType]
	allInts := make([][]int, 0, fileSpec.row)
	var ops []rune
	for scanner.Scan() {
		switch state {
		case stateScanNum:
			cur := scanner.ScanInts(fileSpec.col)
			allInts = append(allInts, cur)
			fmt.Printf("%+x\n", cur)
			i++
			if i >= fileSpec.row {
				state = stateScanOp
			}
		case stateScanOp:
			ops = scanner.ScanOp(fileSpec.col)
			fmt.Printf("%s\n", string(ops))
		}
	}
	// do the math
	finalSum := 0
	for col, op := range ops {
		var opVal int
		switch op {
		case '*':
			opVal = 1
			for _, v := range allInts {
				opVal *= v[col]
			}
		case '+':
			opVal = 0
			for _, v := range allInts {
				opVal += v[col]
			}
		}
		finalSum += opVal
	}
	fmt.Printf("finalsum: %d\n", finalSum)
}

var (
	part     int
	file     string
	silent   bool
	fileType string
	files    map[string]fileParam = map[string]fileParam{
		"sample": {col: 4, row: 3},
		"input":  {col: 1000, row: 4},
	}
	opFunc = map[rune]func(a, b int) int{
		'+': func(a, b int) int { return a + b },
		'*': func(a, b int) int { return a * b },
	}
)

const (
	stateScanNum int = iota
	stateScanOp
	recLen int = 4
	numRec int = 3
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")
	flag.BoolVar(&silent, "s", false, "silent?")
	flag.StringVar(&fileType, "type", "sample", "sample or input")

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

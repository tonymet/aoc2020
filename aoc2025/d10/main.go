package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	_ "sort"
	"strconv"
	"strings"
)

var (
	filetypes = map[string]fileparam{
		"sample.txt": {20, 10, 3},
		"input.txt":  {1000, 1000, 3},
	}
	activeParam fileparam
)

type machine struct {
	lights  []rune
	buttons [][]int
	weights [][]int
}

const (
	stateLights int = iota
	stateButtons
	stateWeights
)

type fileparam struct {
	records, top, productLimit int
}

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

func splitInt(s string) (value []int, err error) {
	parts := strings.Split(s, ",")
	value = make([]int, 0, len(s))
	for _, v := range parts {
		v, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return value, err
		}
		value = append(value, int(v))
	}
	return
}

func part1(in io.Reader) {
	fmt.Printf("part1 not implemented\n")
	// parse
	//var state = stateLights
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		var cur machine
		cur.lights = make([]rune, 0, 10)
		cur.buttons = make([][]int, 0, 10)
		line := scanner.Text()
		//lineReader := strings.NewReader(line)
		fields := strings.Fields(line)
		for _, s := range fields {
			switch s[0] {
			case '[':
				cur.lights = []rune(s[1 : len(s)-1])
			case '(':
				val, err := splitInt(s[1 : len(s)-1])
				if err != nil {
					panic(err)
				}
				cur.buttons = append(cur.buttons, val)
			case '{':
				//pass
			}
		}
		fmt.Printf("cur: %+v\n", cur)
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
	activeParam = filetypes[path.Base(file)]
	switch part {
	case 2:
		part2(os.Stdin)
	default:
		part1(os.Stdin)
	}
}

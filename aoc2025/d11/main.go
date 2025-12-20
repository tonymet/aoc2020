package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	_ "sort"
	"strings"
)

type mymap map[string][]string
type myroute []string
type fftStat map[*[]string]struct {
	fft, dac bool
}

var (
	filetypes = map[string]fileparam{
		"sample.txt": {20, 10, 3},
		"input.txt":  {1000, 1000, 3},
	}
	activeParam fileparam
	numRoutes   int
	tfftStat    fftStat
)

type fileparam struct {
	records, top, productLimit int
}

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

func (mr myroute) String() string {
	var b strings.Builder
	for _, v := range mr {
		b.WriteString(v + " ")
	}
	return b.String()
}

func (m mymap) String() string {
	var b strings.Builder
	for k, v := range m {
		b.WriteString(k + ":\t")
		for _, e := range v {
			b.WriteString(e + "\t")
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func solve(km mymap, start string, route []string) {
	next := km[start]
	if next[0] == "out" {
		numRoutes++
		return
	}
	for _, v := range next {
		solve(km, v, route)
	}

}
func part1(in io.Reader) {
	scanner := bufio.NewScanner(in)
	keyMap := make(map[string][]string)
	tfftStat = make(fftStat)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		keyMap[parts[0]] = strings.Fields(parts[1])
	}
	if !silent {
		fmt.Printf("keymap:\n%s\n", mymap(keyMap))
	}
	solve(keyMap, "you", make([]string, 0, 10))
	fmt.Printf("route count: %d\n", numRoutes)
}

var (
	part   int
	file   string
	silent bool
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")
	flag.BoolVar(&silent, "s", true, "silent?")
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

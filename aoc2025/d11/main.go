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
type fftStat map[*[]string]fftTrack
type fftTrack struct {
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
	iter        int
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

func solve2(km mymap, start, end string, t fftTrack) {
	next := km[start]
	if next[0] == end {
		// if t.dac && t.fft {
		// 	numRoutes++
		// }
		numRoutes++
		return
	}
	for _, v := range next {
		if v == "out" {
			continue
		}
		if v == "fft" {
			t.fft = true
		}
		if v == "dac" {
			t.dac = true
		}
		iter++
		solve2(km, v, end, t)
	}

}
func solve(km mymap, start, end string, t fftTrack) {
	next := km[start]
	if next[0] == end {
		numRoutes++
		return
	}
	for _, v := range next {
		if v == "out" {
			continue
		}
		solve(km, v, end, t)
	}

}

// func parse(in io.Reader) mymap{

// }
func parse(in io.Reader) mymap {
	scanner := bufio.NewScanner(in)
	keyMap := make(map[string][]string)
	tfftStat = make(fftStat)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		keyMap[parts[0]] = strings.Fields(parts[1])
	}
	return keyMap
}
func part1(in io.Reader) {
	keyMap := parse(in)
	if !silent {
		fmt.Printf("keymap:\n%s\n", mymap(keyMap))
	}
	solve(keyMap, "you", "out", fftTrack{})
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

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

type vec struct {
	x, y int
}
type button struct {
	name string
	dir  vec
}

var (
	STATE_BUTTON      = 1
	STATE_PRIZE       = 2
	STATE_NEWLINE     = 3
	state         int = STATE_BUTTON
)

func part1(in io.Reader) {
	lineScanner := bufio.NewScanner(in)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		if line == "" {
			continue
		}
		lineReader := strings.NewReader(line)

		switch line[0] {
		case 'B':
			var curButton button
			_, err := fmt.Fscanf(lineReader, "Button %s X+%d, Y+%d", &curButton.name, &curButton.dir.x, &curButton.dir.y)
			curButton.name = curButton.name[0:1]
			_ = err
			fmt.Printf("B: %+v\n", curButton)
		case 'P':
			var curPrize vec
			_, err := fmt.Fscanf(lineReader, "Prize: X=%d, Y=%d", &curPrize.x, &curPrize.y)
			_ = err
			fmt.Printf("P: %+v\n", curPrize)
		}
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
	switch part {
	case 2:
		part2(os.Stdin)
	default:
		part1(os.Stdin)
	}
}

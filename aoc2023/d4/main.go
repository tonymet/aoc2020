package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	_ "sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

type rec struct {
	winners []int64
	card    []int64
	id      int64
}

func (r rec) String() string {
	b := make([]byte, 0, 64)
	b = append(b, []byte("Card ")...)
	b = strconv.AppendInt(b, int64(r.id), 10)
	b = append(b, ' ')
	for _, v := range r.winners {
		b = strconv.AppendInt(b, int64(v), 10)
		b = append(b, ' ')
	}
	b = append(b, ' ', '|', ' ')
	for _, v := range r.card {
		b = strconv.AppendInt(b, int64(v), 10)
		b = append(b, ' ')
	}
	return string(b)
}

func log(f string, val ...any) {
	if silent {
		return
	}
	fmt.Printf(f, val...)
}

var (
	STATE_START     = 1
	STATE_ID        = 2
	STATE_WIN       = 3
	STATE_CARD      = 4
	state       int = STATE_ID
)

func part2() {
	fmt.Printf("part2 not implemented\n")
}

func part1(in io.Reader) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		// each line
		var curRec rec
		curRec.winners = make([]int64, 0)
		curRec.card = make([]int64, 0)
		line := scanner.Text()
		lineScanner := bufio.NewScanner(strings.NewReader(line))
		lineScanner.Split(bufio.ScanWords)
		state = STATE_WIN
		for lineScanner.Scan() {
			// scan tokens
			switch lineScanner.Text() {
			case "Card":
				lineScanner.Scan()
				t := lineScanner.Text()
				if v, err := strconv.ParseInt(t[:len(t)-1], 10, 64); err != nil {
					fmt.Printf("no parse")
				} else {
					curRec.id = v
				}
			case "|":
				state = STATE_CARD
				fmt.Print("  | ")
			default:
				word := lineScanner.Text()
				if v, err := strconv.ParseInt(word, 10, 64); err != nil {
					fmt.Printf("no parse: %+v", word)
				} else {
					switch state {
					case STATE_WIN:
						curRec.winners = append(curRec.winners, v)
					case STATE_CARD:
						curRec.card = append(curRec.card, v)
					}
				}
			}
		}
		fmt.Println(curRec)
	}
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
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
		part2()
	default:
		part1(os.Stdin)
	}
}

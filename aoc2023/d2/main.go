package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	hands []hand
	id    int
}

type hand struct {
	r, b, g int
}

var bounds = hand{
	12, 14, 13,
}

func possibleGame(g game) bool {
	for _, h := range g.hands {
		if h.r > bounds.r || h.b > bounds.b || h.g > bounds.g {
			return false
		}
	}

	return true
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxHand(g game) (r hand) {
	for _, h := range g.hands {
		r.r = maxInt(h.r, r.r)
		r.b = maxInt(h.b, r.b)
		r.g = maxInt(h.g, r.g)
	}
	return
}

func power(h hand) int {
	return h.r * h.b * h.g
}

func part1() {
	var reGame, reColor = regexp.MustCompile(`^Game (\d+):`), regexp.MustCompile(`(?m)(\d+) (\w+)`)
	scanner := bufio.NewScanner(os.Stdin)
	var total, total2 int
	for scanner.Scan() {
		curLine := scanner.Text()
		var curGame = game{make([]hand, 0), 0}

		matches := reGame.FindStringSubmatch(curLine)
		if i, err := strconv.ParseInt(matches[1], 10, 64); err != nil {
			panic(err)
		} else {
			fmt.Printf("game : %d\n", i)
			curGame.id = int(i)
		}
		start := strings.IndexAny(curLine, ":")
		hands := strings.Split(curLine[start+2:], ";")

		for _, h := range hands {
			colors := reColor.FindAllStringSubmatch(h, -1)
			var curHand hand
			for _, icolor := range colors {
				if val, err := strconv.ParseInt(icolor[1], 10, 64); err != nil {
					panic(err)
				} else {
					switch icolor[2] {
					case "red":
						curHand.r = int(val)
					case "blue":
						curHand.b = int(val)
					case "green":
						curHand.g = int(val)
					}
				}
			}
			curGame.hands = append(curGame.hands, curHand)
		}
		fmt.Printf("curGame: %+x\n", curGame)
		fmt.Printf("curGame possible: %t\n", possibleGame(curGame))
		maxHandValue := maxHand(curGame)
		fmt.Printf("maxHand possible: %x\n", maxHandValue)
		if possibleGame(curGame) {
			total += curGame.id
		}
		total2 += power(maxHandValue)
	}
	fmt.Printf("total : %d\n", total)
	fmt.Printf("total p2 : %d\n", total2)
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
		fallthrough
	default:
		part1()
	}
}

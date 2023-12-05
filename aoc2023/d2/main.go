package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var part = 1

type game struct {
	hands []hand
	i     int
}

type hand struct {
	r, b, g int
}

func part1() {

	var reGame = regexp.MustCompile(`^Game (\d+):`)
	var reColor = regexp.MustCompile(`(?m)(\d+) (\w+)`)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		curLine := scanner.Text()
		var curGame game
		curGame.hands = make([]hand, 0)

		matches := reGame.FindStringSubmatch(curLine)
		if i, err := strconv.ParseInt(matches[1], 10, 64); err != nil {
			panic(err)
		} else {
			fmt.Printf("game : %d\n", i)
			curGame.i = int(i)
		}
		start := strings.IndexAny(curLine, ":")
		rest := curLine[start+2:]
		hands := strings.Split(rest, ";")

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
	}
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
		fallthrough
	default:
		part1()
	}
}

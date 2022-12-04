package main

import (
	"fmt"
	"io"
	"os"
)

//  first col: A Rock, B for Paper, and C for Scissors.
//  second column,  X for Rock, Y for Paper, and Z for Scissors.

const (
	win  int = 6
	lose int = 0
	draw int = 3
)

var oValMap = map[rune]int{
	'W': win,
	'L': lose,
	'D': draw,
}

var lMap = map[rune]rune{
	'A': 'R',
	'B': 'P',
	'C': 'S',
}
var rMap = map[rune]rune{
	'X': 'R',
	'Y': 'P',
	'Z': 'S',
}
var rMapP2 = map[rune]rune{
	'X': 'L',
	'Y': 'D',
	'Z': 'W',
}

var rValMap = map[rune]int{
	'R': 1,
	'P': 2,
	'S': 3,
}

var eMap = map[rune]int{
	'X': lose,
	'Y': draw,
	'Z': win,
}

type move [2]rune

func (m move) String() string {
	return fmt.Sprintf("%s %s\n", string(m[0]), string(m[1]))
}

func part2() {
	var cur, trans move
	var i, total int
	var moves []move
	moves = make([]move, 0)

	for _, err := fmt.Scanf("%c %c\n", &cur[0], &cur[1]); err != io.EOF; _, err = fmt.Scanf("%c %c\n", &cur[0], &cur[1]) {
		var outcomeVal, moveVal int
		trans[0], trans[1] = lMap[cur[0]], rMapP2[cur[1]]
		moves = append(moves, trans)
		fmt.Printf("cur: %+v, trans: %+v \n", cur, trans)
		movePlayed := play2(trans)
		moveVal, outcomeVal = rValMap[movePlayed], oValMap[trans[1]]
		fmt.Printf("movePlayed: %s, moveVal : %d, outcomeval: %d\n", string(movePlayed), moveVal, outcomeVal)
		total += moveVal + outcomeVal
		i++
	}
	fmt.Printf("moves : %+v\n", moves)
	fmt.Printf("total : %d\n", total)
}

func play(m move) int {
	if m[0] == m[1] {
		return draw
	}
	switch {
	case m[0] == 'R' && m[1] == 'P':
		return win
	case m[0] == 'R' && m[1] == 'S':
		return lose
	case m[0] == 'P' && m[1] == 'R':
		return lose
	case m[0] == 'P' && m[1] == 'S':
		return win
	case m[0] == 'S' && m[1] == 'R':
		return win
	case m[0] == 'S' && m[1] == 'P':
		return lose
	default:
		panic("bad combo")
	}
}

func play2(m move) rune {
	switch {
	case m[0] == 'R' && m[1] == 'W':
		return 'P'
	case m[0] == 'R' && m[1] == 'L':
		return 'S'
	case m[0] == 'R' && m[1] == 'D':
		return 'R'
	case m[0] == 'P' && m[1] == 'W':
		return 'S'
	case m[0] == 'P' && m[1] == 'L':
		return 'R'
	case m[0] == 'P' && m[1] == 'D':
		return 'P'
	case m[0] == 'S' && m[1] == 'W':
		return 'R'
	case m[0] == 'S' && m[1] == 'L':
		return 'P'
	case m[0] == 'S' && m[1] == 'D':
		return 'S'
	default:
		panic("bad combo")
	}
}

func part1() {
	var cur, trans move
	var i, total int
	var moves []move
	moves = make([]move, 0)

	for _, err := fmt.Scanf("%c %c\n", &cur[0], &cur[1]); err != io.EOF; _, err = fmt.Scanf("%c %c\n", &cur[0], &cur[1]) {
		trans[0], trans[1] = lMap[cur[0]], rMap[cur[1]]
		moves = append(moves, trans)
		fmt.Printf("cur: %+v \n", cur)
		fmt.Printf("outcome: %d, moveVal: %d\n", play(trans), rValMap[trans[1]])
		total += play(trans) + rValMap[trans[1]]
		i++
	}
	fmt.Printf("moves : %+v\n", moves)
	fmt.Printf("total : %d\n", total)
}
func main() {
	switch os.Getenv("PART") {
	case "2":
		part2()
	default:
		part1()
	}
}

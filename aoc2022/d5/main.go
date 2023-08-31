package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

//  first col: A Rock, B for Paper, and C for Scissors.
//  second column,  X for Rock, Y for Paper, and Z for Scissors.

const (
	WIDTH  int = 3
	HEIGHT int = 3
)

type LMR [3]string

type LMRuniq [3]uniqType

func charVal(c rune) int {
	// a = 97, z = 123
	// A = 65, Z=91

	if c < 65 || c > 123 {
		panic(fmt.Sprintf("out of range: %d %s ", c, string(c)))
	}
	if c > 91 && c < 97 {
		panic(fmt.Sprintf("out of range: %d %s ", c, string(c)))
	}
	switch {
	case c >= 97:
		return int(c) - 96
	case c >= 65:
		return int(c) - 64 + 26
	default:
		panic("out of bounds")
	}
}

type uniqType map[rune]bool
type foundType []rune

func uniq(s string) (u uniqType) {
	u = make(uniqType)
	for _, v := range s {
		u[v] = true
	}
	return
}

func uniq2(l, r string) (found foundType) {
	uniqL := uniq(l)
	uniqR := uniq(r)
	found = make(foundType, 0)
	for k := range uniqR {
		_, ok := uniqL[k]
		if ok {
			found = append(found, k)
		}
	}
	return
}

func (m uniqType) String() string {
	s := make([]string, len(m))
	for v := range m {
		s = append(s, string(v))
	}
	return fmt.Sprintf("%v\n", strings.Join(s, " "))
}
func (m foundType) String() string {
	s := make([]string, len(m))
	for _, v := range m {
		s = append(s, string(v))
	}
	return fmt.Sprintf("%v\n", strings.Join(s, " "))
}

func keep(index int) bool {
	return ((index - 1) % 4) == 0
}

func colToTower(col int) int {
	return col / 4
}
func isLetter(l byte) bool {
	return l >= 65 && l <= 90
}

func (m foundType) rawString() string {
	s := make([]string, len(m))
	for _, v := range m {
		s = append(s, string(v))
	}
	return fmt.Sprintf("%v\n", strings.Join(s, ""))
}

func part2() {
}
func part1() {
	var b []byte = make([]byte, 1)
	var towers [][]byte = make([][]byte, WIDTH)
	for i := 0; i < WIDTH; i++ {
		towers[i] = make([]byte, 0, HEIGHT)
	}
	var row, col int
	for {
		_, err := os.Stdin.Read(b)
		if err == io.EOF {
			break
		}
		if keep(col) {
			fmt.Printf("b:%s, row: %d, col %d, toTower: %d\n", string(b), row, col, colToTower(col))
			if isLetter(b[0]) {
				towers[colToTower(col)] = append(towers[colToTower(col)], b[0])
			}
		}
		if col == 0 && b[0] == '\n' {
			break
		}
		col++
		if b[0] == '\n' {
			row++
			col = 0
		}
	}
	for i := range towers {
		slices.Reverse(towers[i])
	}
	fmt.Printf("towers: %+v\n", towers)

	var moves = make([][3]int, 0)
	var cur [3]int
	for {
		_, err := fmt.Scanf("move %d from %d to %d\n", &cur[0], &cur[1], &cur[2])
		if err != nil {
			break
		}
		if err == io.EOF {
			break
		}
		fmt.Printf("curMove: %+v\n", cur)
		moves = append(moves, cur)
	}
	fmt.Printf("moves: %+v\n", moves)
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
		part2()
	default:
		part1()
	}
}

package d5

import (
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
)

// bounds for the tower b/c we are streaming the file
var (
	WIDTH  *int
	HEIGHT *int
)

type towersTypeA = [][]byte
type towerTypeA = []byte
type towerType []byte
type towersClass struct {
	towerStorage towersTypeA
}

func pop(t *towerTypeA) (cur byte) {
	cur = (*t)[len(*t)-1]
	(*t) = (*t)[:len(*t)-1]
	return
}

func push(t *towerTypeA, e byte) {
	(*t) = append(*t, e)
}

func printTops(t towersTypeA) {
	for _, eachTower := range t {
		if len(eachTower) == 0 {
			panic("eachTower is empty")
		}
		fmt.Print(string(eachTower[len(eachTower)-1]))
	}
	fmt.Print("\n")
}

func (ts *towersClass) move(n, f, t int) {
	for i := 0; i < n; i++ {
		if len(ts.towerStorage[f]) == 0 {
			panic("from is empty")
		}
		e := pop(&ts.towerStorage[f])
		push(&ts.towerStorage[t], e)
	}
}

func (ts *towersClass) move2(n, f, t int) {
	// move to temporary , reverse and then push
	tmp := make([]byte, 0)
	for i := 0; i < n; i++ {
		if len(ts.towerStorage[f]) == 0 {
			panic("from is empty")
		}
		e := pop(&ts.towerStorage[f])
		tmp = append(tmp, e)
	}
	slices.Reverse(tmp)
	ts.towerStorage[t] = append(ts.towerStorage[t], tmp...)
}

// utilities for parsing the tower into data structures
func keep(index int) bool {
	return ((index - 1) % 4) == 0
}

func colToTower(col int) int {
	return col / 4
}
func isLetter(l byte) bool {
	return l >= 65 && l <= 90
}

func parseAndSolve() {
	var b []byte = make([]byte, 1)
	var towers [][]byte = make(towersTypeA, *WIDTH)
	for i := 0; i < *WIDTH; i++ {
		towers[i] = make(towerType, 0, *HEIGHT)
	}
	var aTower towersClass
	var row, col int
	for {
		_, err := os.Stdin.Read(b)
		if err == io.EOF {
			break
		}
		// keep will skip over benign characters
		if keep(col) {
			fmt.Printf("b:%s, row: %d, col %d, toTower: %d\n", string(b), row, col, colToTower(col))
			if isLetter(b[0]) {
				towers[colToTower(col)] = append(towers[colToTower(col)], b[0])
			}
		}
		// end of tower reading
		if col == 0 && b[0] == '\n' {
			break
		}
		// update row , col values
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
	aTower.towerStorage = towers

	var moves = make([][3]int, 0)
	var cur [3]int
	for {
		_, err := fmt.Scanf("move %d from %d to %d\n", &cur[0], &cur[1], &cur[2])
		if err != nil {
			break
		}
		fmt.Printf("curMove: %+v\n", cur)
		moves = append(moves, cur)
	}

	switch os.Getenv("PART") {
	case "2":
		for _, m := range moves {
			aTower.move2(m[0], m[1]-1, m[2]-1)
		}
		fmt.Printf("AFTER moves: %+v\n", aTower.towerStorage)
		printTops(aTower.towerStorage)
	default:
		// perform the moves
		for _, m := range moves {
			aTower.move(m[0], m[1]-1, m[2]-1)
		}
		fmt.Printf("AFTER moves: %+v\n", aTower.towerStorage)
		printTops(aTower.towerStorage)
	}
}

func init() {
	WIDTH = flag.Int("width", 3, "width")
	HEIGHT = flag.Int("height", 3, "height")
}

func main() {
	flag.Parse()
	if stdin := os.Getenv("STDIN"); len(stdin) != 0 {
		stdinFile, err := os.Open(stdin)
		if err != nil {
			panic(err)
		}
		os.Stdin = stdinFile
	}
	parseAndSolve()
}

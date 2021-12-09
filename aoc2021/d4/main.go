package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var bingoNumbers []int
var bingoLookup = map[int][]boardRowCol{}
var boards [100]*tBoard

type tBoard = [5][5]int
type tBoardPtr = *[5][5]int

type tWinner = struct {
	board *tBoard
	num   int
}

type boardRowCol struct {
	boardRowCol *[5]int
	board       *tBoard
}

type boardRowColPtr = *boardRowCol
type xy struct {
	x, y int
}

func bingoTest(rowCol [5]int) bool {
	for _, v := range rowCol {
		if v != -1 {
			return false
		}
	}
	return true
}

func sumBoard(board tBoard) (sum int) {
	for _, row := range board {
		for _, cell := range row {
			if cell != -1 {
				sum += cell
			}
		}
	}
	return
}

func part1Solution(board tBoard, num int) int {
	return sumBoard(board) * num
}

func indexOf(haystack [5]int, needle int) (int, error) {
	for i, v := range haystack {
		if needle == v {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}

func part3() {
	var line string
	fmt.Printf("hello debugger\n")
	fmt.Scanf("%s", &line)
	fmt.Printf("hello debugger line: %+v\n", line)

}

func part2() {
	// play the game
	//var bingoWinners = make([]tWinner, 0)
	parseAndSetup()

	// index all boards
	// we will remove the seen ones
	bingoBoards := make(map[*tBoard]bool)
	for _, board := range boards {
		bingoBoards[board] = false
	}
	for _, num := range bingoNumbers {
		for _, rowCol := range bingoLookup[num] {
			// skip the completed cards
			if _, ok := bingoBoards[rowCol.board]; !ok {
				continue
			}
			rowColRecord := rowCol.boardRowCol
			fmt.Printf("num: %+v, rowCol: %+v\n", num, rowCol)
			// find index
			if index, err := indexOf(*rowColRecord, num); err == nil {
				rowColRecord[index] = -1
			} else {
				panic("index num not found")
			}

			// check bingo
			if bingoTest(*rowColRecord) {
				fmt.Printf("Bingo: %+v, solution: %+v\n", rowCol.board, part1Solution(*rowCol.board, num))
				// remove the board
				delete(bingoBoards, rowCol.board)
				if len(bingoBoards) == 0 {
					fmt.Printf("FINAL Bingo: %+v, solution: %+v\n", rowCol.board, part1Solution(*rowCol.board, num))
					return
				}
			}
		}
	}
}

func parseAndSetup() {
	for {
		var cur int
		n, err := fmt.Scanf("%d,", &cur)
		bingoNumbers = append(bingoNumbers, cur)
		bingoLookup[cur] = make([]boardRowCol, 0)
		if n != 1 {
			panic("end of reading")
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	// read the gap
	fmt.Scanf("\n")

	var i, x, y int
boards:
	for {
		var curBoard tBoard
		var curCols tBoard
		for {
			var cur int
			if x == 5 {
				break
			}
			_, err := fmt.Scanf("%d", &cur)
			curCols[y][x], curBoard[x][y] = cur, cur
			fmt.Printf("cur: %+v, curBoardRow: %+v, curBoardCol: %+v\n", cur, curBoard[x], curCols[y])
			bingoLookup[cur] = append(bingoLookup[cur], boardRowCol{&curBoard[x], &curBoard}, boardRowCol{&curCols[y], &curBoard})
			if err == io.EOF {
				break boards
			}
			y = (y + 1) % 5
			if y == 0 {
				x++
			}
		}
		boards[i] = &curBoard
		i++
		x, y = 0, 0
		// read off blank line
		fmt.Scanf("\n")
	}
}
func part1() {
	// play the game
	//var bingoWinners = make([]tWinner, 0)
	parseAndSetup()
	for _, num := range bingoNumbers {
		for _, rowCol := range bingoLookup[num] {
			rowColRecord := rowCol.boardRowCol
			fmt.Printf("num: %+v, rowCol: %+v\n", num, rowCol)
			// find index
			if index, err := indexOf(*rowColRecord, num); err == nil {
				rowColRecord[index] = -1
			} else {
				panic("index num not found")
			}

			// check bingo
			if bingoTest(*rowColRecord) {
				fmt.Printf("Bingo: %+v, solution: %+v\n", rowCol.board, part1Solution(*rowCol.board, num))
				return
			}
		}
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
	case "3":
		part3()
	case "2":
		part2()
	default:
		part1()
	}
}

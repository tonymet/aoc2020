package main

import (
	"errors"
	"fmt"
	"io"
)

type execError struct {
	code int
}

var ErrLoop = errors.New("loop")
var ErrExited = errors.New("exited")
var ErrOOB = errors.New("OOB")
var ErrNotFound = errors.New("Not Found")

type inst struct {
	cmd          string
	arg, visited int
}

var theProgram []inst

func scanFile() {
	theProgram = make([]inst, 0)
	for {
		var cur inst
		_, err := fmt.Scanf("%s %d", &cur.cmd, &cur.arg)
		if err != nil {
			fmt.Errorf("error reading")
			if err == io.EOF {
				fmt.Errorf("End of program\n")
				break
			}
		}
		theProgram = append(theProgram, cur)

	}
	//fmt.Printf("theProgram len: %d\n", len(theProgram))
}

func exec(curProgram []inst) (int, error) {
	// iterate at begining
	// case on instruction
	// increment visit
	// store accumulator
	// error when visited > 0
	var (
		accumulator = 0
		i           = 0
		cur         *inst
	)
	//fmt.Printf("len curProgram: %d\n", len(curProgram))
	for {
		if i >= len(curProgram) {
			fmt.Printf("out of bounds : %+v, acc: %d\n", cur, accumulator)
			return accumulator, ErrOOB
		}
		cur = &curProgram[i]
		cur.visited++
		if cur.visited > 1 {
			//fmt.Printf("already seen : %+v, acc: %d\n", cur, accumulator)
			return accumulator, ErrLoop
		}
		switch cur.cmd {
		case "jmp":
			i += cur.arg
		case "acc":
			accumulator += cur.arg
			i++
		case "nop":
			i++
		case "exit":
			//fmt.Printf("exited : %+v, acc: %d\n", cur, accumulator)
			return accumulator, ErrExited
		default:
			panic("cmd err")
		}
	}
}

func resetProgram(curProgram *[]inst) {
	for i := range *curProgram {
		(*curProgram)[i].visited = 0
	}
}
func part2Exec(curProgram []inst) (int, error) {
	// add exit command
	// brute force changes
	// run until exit is hit
	curProgram = append(curProgram, inst{"exit", 0, 0})
	for i, cur := range curProgram {
		copyOfProgram := make([]inst, len(curProgram))
		copy(copyOfProgram, curProgram)
		resetProgram(&copyOfProgram)
		switch cur.cmd {
		case "jmp":
			copyOfProgram[i].cmd = "nop"
		case "nop":
			copyOfProgram[i].cmd = "jmp"
		}
		acc, err := exec(copyOfProgram)
		if err == ErrExited {
			return acc, err
		}
	}
	return 0, ErrNotFound
}

func main() {
	scanFile()
	acc, err := exec(theProgram)
	if err == ErrLoop {
		fmt.Printf("Part1 loop: accumulator = %d\n", acc)
	}
	acc, err = part2Exec(theProgram)
	if err == ErrExited {
		fmt.Printf("Part2 exited: accumulator = %d\n", acc)
	} else {
		panic(err)
	}
}

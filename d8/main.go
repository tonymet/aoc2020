package main

import (
	"flag"
	"fmt"
	"io"
)

var part2 *bool

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

func exec(curProgram []inst) {
	// iterate at begining
	// case on instruction
	// increment visit
	// store accumulator
	// error when visited > 0
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if recover() != nil {
			fmt.Printf("error out of bounds\n")
		}
	}()
	var (
		accumulator = 0
		i           = 0
		cur         *inst
	)
	//fmt.Printf("len curProgram: %d\n", len(curProgram))
	for {
		if i >= len(curProgram) {
			fmt.Printf("out of bounds : %+v, acc: %d\n", cur, accumulator)
			return
		}
		cur = &curProgram[i]
		cur.visited++
		if cur.visited > 1 {
			fmt.Printf("already seen : %+v, acc: %d\n", cur, accumulator)
			return
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
			fmt.Printf("exited : %+v, acc: %d\n", cur, accumulator)
			return
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
func part2Exec(curProgram []inst) {
	// add exit command
	// brute force changes
	// run until exit is hit
	curProgram = append(curProgram, inst{"exit", 0, 0})
	for i, cur := range curProgram {
		copyOfProgram := make([]inst, len(curProgram))
		copy(copyOfProgram, curProgram)
		switch cur.cmd {
		case "jmp":
			fmt.Printf("line %d changing jmp to nop\n", i)
			copyOfProgram[i].cmd = "nop"
			exec(copyOfProgram)
		case "nop":
			fmt.Printf("line %d changing nop to jmp\n", i)
			copyOfProgram[i].cmd = "jmp"
			exec(copyOfProgram)
		}
	}
}

func main() {
	part2 = flag.Bool("part2", false, "part2")
	scanFile()
	exec(theProgram)
	part2Exec(theProgram)
}

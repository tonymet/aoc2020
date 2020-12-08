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
	fmt.Printf("theProgram len: %d\n", len(theProgram))
}

func exec() {
	// iterate at begining
	// case on instruction
	// increment visit
	// store accumulator
	// error when visited > 0
	var accumulator = 0
	i := 0
	cur := &theProgram[i]
	for {
		cur.visited++
		if cur.visited > 1 {
			fmt.Printf("already seen : %+v, acc: %d", cur, accumulator)
			return
		}
		switch cur.cmd {
		case "jmp":
			i += cur.arg
			cur = &theProgram[i]
		case "acc":
			accumulator += cur.arg
			i++
			cur = &theProgram[i]
		case "nop":
			i++
			cur = &theProgram[i]
		default:
			panic("cmd err")
		}
	}
}

func main() {
	part2 = flag.Bool("part2", false, "part2")
	scanFile()
	exec()
}

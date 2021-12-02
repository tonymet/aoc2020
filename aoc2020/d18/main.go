package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type istack []int
type sstack []string

var ErrorEOL = errors.New("end of line")

func (is *istack) pop() int {
	r := (*is)[len(*is)-1]
	*is = (*is)[:len(*is)-1]
	return r
}
func (is *istack) push(a int) {
	*is = append(*is, a)
}
func (is *sstack) push(a string) {
	*is = append(*is, a)
}

func (is *sstack) pop() string {
	r := (*is)[len(*is)-1]
	*is = (*is)[:len(*is)-1]
	return r
}

func eval(op string, l, r int) int {
	switch op {
	case "+":
		return l + r
	case "*":
		return l * r
	default:
		panic("bad op: " + op)
	}
}

func scanFile() int {
	var (
		args istack
		ops  sstack
		sum  int
		//parens sstack
	)

	for {
		var cur string
		_, err := fmt.Scanf("%1s", &cur)
		if err == io.EOF {
			break
		}
		switch cur {
		case "", "\n":
			for len(args) > 1 && len(ops) > 0 {
				r := args.pop()
				l := args.pop()
				op := ops.pop()
				args.push(eval(op, l, r))
			}
			fmt.Printf("line result: %d\n", args[len(args)-1])
			sum += args.pop()
			// reset
			args, ops = make(istack, 0), make(sstack, 0)
		case "+", "*":
			ops.push(cur)
		case "(":
			call := scanFile()
			args.push(call)
			// call recursive
		case ")":
			for len(args) > 1 && len(ops) > 0 {
				r := args.pop()
				l := args.pop()
				op := ops.pop()
				args.push(eval(op, l, r))
			}
			return args.pop()
		default:
			if arg, err := strconv.ParseInt(cur, 10, 32); err == nil {
				args.push(int(arg))
			}
			// if there is an arg and an op, pop both , eval and push the result to arg
			for len(args) > 1 && len(ops) > 0 {
				r := args.pop()
				l := args.pop()
				op := ops.pop()
				args.push(eval(op, l, r))
			}
		}
	}
	fmt.Printf("sum: %d\n", sum)
	//return args[len(args)-1]
	return 0
}

func main() {
	scanFile()
}

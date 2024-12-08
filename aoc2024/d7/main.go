package main

// test

import (
	"flag"
	"fmt"
	"github.com/tonymet/aoc2020/shared"
	"io"
	"math/rand"
	"os"
)

var (
	OP_TYPE_VALUE = 1
	OP_TYPE_OP    = 2
)

type paramType []int64
type opType func(int64, int64) int64
type opValue struct {
	typeOf int
	v      int64
	op     func(int64, int64) int64
}

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

func calcStack(opsStack []opValue) int64 {
	for {
		if len(opsStack) == 1 {
			return opsStack[0].v
		}
		// shift off and push
		curOps := opsStack[0:3]
		opsStack = opsStack[3:]
		v := curOps[1].op(curOps[0].v, curOps[2].v)
		opsStack = append([]opValue{{typeOf: OP_TYPE_VALUE, v: v}}, opsStack...)
	}
}

func (p paramType) bruteForce(want int64) bool {
	// brute force all variations and see if we return
	ops := []opType{
		func(a, b int64) int64 { return a + b },
		func(a, b int64) int64 { return a * b },
	}
	//po
	// initial ops

	for pos := 1; pos < (2000 * (len(p) - 1)); pos++ {
		opsStack := make([]opValue, 0, len(p)+len(p)-1)
		for i := 0; i < len(p); i++ {
			operand := opValue{v: p[i], typeOf: OP_TYPE_VALUE}
			opsStack = append(opsStack, operand)
			if i == len(p)-1 {
				continue
			}
			operator := opValue{typeOf: OP_TYPE_OP, op: ops[(i+rand.Int())%2]}
			opsStack = append(opsStack, operator)
		}
		// calculate
		if calcStack(opsStack) == want {
			return true
		}
	}
	return false
}

func part1(in io.Reader) {
	sum := int64(0)
	shared.LineProcessor(in, func(line io.Reader) {
		// scan values
		var (
			want   int64
			params paramType
		)
		params = make(paramType, 0, 8)
		_, err := fmt.Fscanf(line, "%d:", &want)
		if err == io.EOF {
			panic(io.ErrUnexpectedEOF)
		} else if err != nil {
			panic(err)
		}
		for {
			var p int64
			_, err := fmt.Fscan(line, &p)
			if err == io.EOF {
				// expected
				break
			} else if err != nil {
				panic(err)
			}
			params = append(params, p)
		}
		fmt.Printf("w: %d ", want)
		viable := params.bruteForce(want)
		fmt.Printf("p: %+v, viable: %t\n", params, viable)
		if viable {
			sum += want
		}
	})
	fmt.Printf("sum: %d\n", sum)
}

var (
	part   int
	file   string
	silent bool
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")
	flag.BoolVar(&silent, "s", false, "silent?")

}

func main() {
	flag.Parse()
	if file != "" {
		var err error
		if os.Stdin, err = os.Open(file); err != nil {
			panic(err)
		}
	}
	switch part {
	case 2:
		part2(os.Stdin)
	default:
		part1(os.Stdin)
	}
}

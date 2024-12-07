package main

import (
	"flag"
	"fmt"
	s "github.com/tonymet/aoc2020/shared"
	"io"
	"os"
	_ "sort"
	"strings"
)

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

type vec struct {
	x, y int
}

type agent struct {
	pos, dir vec
}

type gridType [][]byte

var (
	grid     gridType
	bounds   vec
	theAgent agent
	NUL      byte = 0x00
	up            = vec{0, -1}
	right         = vec{1, 0}
	down          = vec{0, 1}
	left          = vec{-1, 0}
)

func (g gridType) String() string {
	var b strings.Builder
	for _, v := range g {
		b.Write(v)
		b.WriteByte('\n')
	}
	return b.String()
}

func peek(from, dir vec) byte {
	next := vec{from.x + dir.x, from.y + dir.y}
	if next.x < 0 || next.y < 0 ||
		next.y >= bounds.y || next.x >= bounds.x {
		return NUL
	}
	return grid[next.y][next.x]
}

// walk and count steps
func walk() int64 {
	r := int64(1)
	cur := theAgent
	grid[cur.pos.y][cur.pos.x] = 'X'
outer:
	for {
		switch peek(cur.pos, cur.dir) {
		case '#':
			cur.dir = turn(cur.dir)
		case NUL:
			break outer
		case 'X':
			r--
			fallthrough
		default:
			//walk
			cur.pos.x, cur.pos.y = cur.pos.x+cur.dir.x, cur.pos.y+cur.dir.y
			grid[cur.pos.y][cur.pos.x] = 'X'
			//s.Log("%s", grid)
			r++
		}
	}
	return r
}

func turn(from vec) vec {
	switch from {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	default:
		panic("not valid")
	}
}

func part1(in io.Reader) {
	grid = make(gridType, bounds.y)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]byte, bounds.x)
	}
	var (
		x = 0
		y = 0
	)
	for {
		var cur byte
		if _, err := fmt.Fscanf(in, "%c", &cur); err == io.EOF {
			break
		}
		switch cur {
		case '\n':
			x = 0
			y++
			// add newline
		case '^':
			theAgent.pos = vec{x, y}
			theAgent.dir = up
			fallthrough
		default:
			grid[y][x] = cur
			x++
		}
	}
	fmt.Printf("grid: \n%+v\n", grid)
	count := walk()
	fmt.Printf("count: %d\n", count)

}

var (
	part   int
	file   string
	silent bool
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")
	flag.BoolVar(&s.Silent, "s", false, "silent?")

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
	case 0:
		bounds.x, bounds.y = 10, 10
		part1(os.Stdin)
	default:
		bounds.x, bounds.y = 130, 130
		part1(os.Stdin)
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Rot struct {
	dir rune
	len int64
}

type dial struct {
	pos int64
}

var (
	d dial = dial{pos: 50}
	c int64
)

func (d *dial) op(r Rot) {
	// if L , subtract value mod 99
	// if R , add value mod 99
	switch r.dir {
	case 'L':
		d.pos = (d.pos - r.len) % 100
		if d.pos < 0 {
			d.pos += 100
		}
	case 'R':
		d.pos = (d.pos + r.len) % 100
		if d.pos < 0 {
			d.pos += 100
		}
	}
}

type RScanner struct {
	bufio.Scanner
}

func (rs *RScanner) Rot() (r Rot, err error) {
	t := rs.Text()
	if len(t) < 2 {
		return r, fmt.Errorf("parsing error")
	}
	r.dir = rune(t[0])
	r.len, err = strconv.ParseInt(string(t[1:]), 10, 32)
	return r, err
}

func (rs Rot) String() string {
	return fmt.Sprintf("%c%d", rs.dir, rs.len)
}

func part1(in io.Reader) {
	lineScanner := bufio.NewScanner(in)
	rScanner := RScanner{*lineScanner}
	for rScanner.Scan() {
		if rot, err := rScanner.Rot(); err != nil {
			panic(err)
		} else {
			fmt.Printf("rot : %s\t", rot)
			d.op(rot)
			fmt.Printf("d.pos : %d\n", d.pos)
			if d.pos == 0 {
				c++
			}
		}
	}
	fmt.Printf("count : %d\n", c)
}

var (
	part int
	file string
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")
}

func main() {
	flag.Parse()
	if file != "" {
		var err error
		if os.Stdin, err = os.Open(file); err != nil {
			panic(err)
		}
	}
	part1(os.Stdin)
}

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
	c  int64
	fc int64
)

func iAbs(b int64) int64 {
	if b < 0 {
		return -1 * b
	}
	return b
}

func (d *dial) op(r Rot) (zc int64) {
	lpos, llen := d.pos, r.len
	switch r.dir {
	case 'L':
		d.pos = (d.pos - r.len) % 100
		if d.pos < 0 {
			d.pos += 100
		}
		r := llen % 100
		d := iAbs(llen) - r
		zc += d / 100
		if lpos > 0 && lpos-r <= 0 {
			zc++
		}
	case 'R':
		d.pos = (d.pos + r.len) % 100
		if lpos+llen > 99 {
			zc += (lpos + llen) / 100
		}
	}
	return
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
	d := dial{pos: 50}
	for rScanner.Scan() {
		if rot, err := rScanner.Rot(); err != nil {
			panic(err)
		} else {
			fmt.Printf("d.pos : %d\t, rot : %s\t", d.pos, rot)
			zc := d.op(rot)
			fmt.Printf("d.pos : %d, d.zc : %d\n", d.pos, zc)
			if d.pos == 0 {
				c++
			}
			if zc < 0 {
				panic("oob")
			}
			fc += zc
		}
	}
	fmt.Printf("count : %d\t, fc: %d\n", c, fc)
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

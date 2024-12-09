package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	_ "sort"
	"strconv"
)

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

func parseInt(c byte) int64 {
	if r, err := strconv.ParseInt(string(c), 10, 64); err != nil {
		panic(err)
	} else {
		return r
	}
}

func part1(in io.Reader) {
	fmt.Printf("part1 not implemented\n")
	// read size creat eblock switch state
	state = STATE_FILE
	id := int64(0)
	var (
		root     = block{id: -1}
		curBlock = &root
		end      *block
	)
	for {
		var curLen byte
		_, err := fmt.Fscanf(in, "%c", &curLen)
		if err == io.EOF || curLen == '\n' {
			break
		} else if err != nil {
			panic(err)
		}
		var b = block{id: id,
			len:   parseInt(curLen),
			empty: (id % 2) != 0,
			next:  nil,
		}
		id++
		b.prev = curBlock
		curBlock.next = &b
		curBlock = &b
		end = curBlock
	}
	fmt.Println(root.next)
	fmt.Printf("curScore: %d", root.next.score())
	// defrag
	// l - r iteration
	for l, r := root.nextFree(), end; l.id != r.id; l, r = l.next, r.prev {
		// disconnect r
		// insert to l
		// continue until the end

	}

}

// multiplying each of these blocks' position with the file ID number it
// contains. The leftmost block is in position 0. If a block contains free space,
// skip it instead.
func (b *block) score() int64 {
	score := int64(0)
	pos := int64(0)
	b.iter(func(cur *block) {
		score += pos * cur.id
		pos++
	})
	return score
}

/*
func moveByte(to *block, from *block) error{
	// handle empty case
	if !to.empty || from.len < 1{
		panic("error: to is not empty or from.len < 1")
	}
	to.


}
*/

func (b *block) nextFree() *block {
	if b.empty || b.next == nil {
		return b
	}
	for {
		b = b.next
		if b.empty || b.next == nil {
			return b
		}
	}

}

func (b *block) String() string {
	buf := make([]byte, 0, 50)
	b.iter(func(cur *block) {
		buf = strconv.AppendInt(buf, cur.len, 10)
	})
	return string(buf)
}

// linked list of blocks
// first block is a file ID = 0
// then space
// then file ID = 1 and so on
type block struct {
	id    int64
	len   int64
	empty bool
	next  *block
	prev  *block
}

// ops
// head
// cur
// insert

func (b *block) insertAfter(new *block) *block {
	new.next = b.next
	b.next = new
	return new
}

func (b *block) iter(iterator func(*block)) {
	for c := b; true; c = c.next {
		iterator(c)
		if c.next == nil {
			break
		}
	}
}

var (
	part     int
	file     string
	silent   bool
	state    int
	errorEOL = errors.New("End of List")
)

const (
	STATE_FILE  = 1
	STATE_SPACE = 2
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

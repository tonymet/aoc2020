package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	_ "sort"
	"strconv"
)

type fsType []int64
type blockType []int64

var (
	EMPTY = int64(-1)
)

func (fs fsType) String() string {
	var output = make([]byte, 0, len(fs))
	for _, v := range fs {
		if v == EMPTY {
			output = append(output, '.')
			continue
		}
		output = strconv.AppendInt(output, v, 10)
	}
	return string(output)
}

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

func (fs fsType) cksum() (r int64) {
	for _, v := range fs {
		if v == -1 {
			continue
		}
		r += v
	}
	return
}

func (fs fsType) score() (r int64) {
	for i, v := range fs {
		if v == -1 {
			return
		}
		r += int64(i) * v
	}
	return
}

func (fs fsType) defrag() fsType {
	// defrag l , r
	for l, r := 0, len(fs)-1; l <= r-2; {
		if fs[l] != EMPTY {
			l++
			continue
		}
		if fs[r] == EMPTY {
			r--
			continue
		}
		fs[l] = fs[r]
		fs[r] = EMPTY
	}
	return fs
}

func part1(in io.Reader) {
	fs := readInput(in)
	//fmt.Printf("fs: %s", fs)
	fmt.Printf("orig cksum: %d\n", fs.cksum())
	fs.defrag()
	//fmt.Printf("defrag: %s\n", fs)
	fmt.Printf("score: %d\n", fs.score())
	fmt.Printf("cksum: %d\n", fs.cksum())
}

func readInput(in io.Reader) fsType {
	fmt.Printf("part1 not implemented\n")
	var (
		fs    = make(fsType, 0, 20000*10)
		id    = int64(0)
		i     = int(0)
		empty = false
	)
	for {
		var curLen byte
		empty = (i % 2) != 0
		_, err := fmt.Fscanf(in, "%c", &curLen)
		if err == io.EOF || curLen == '\n' {
			break
		} else if err != nil {
			panic(err)
		}
		block := make(blockType, parseInt(curLen))
		var initValue int64 = id
		if empty {
			initValue = EMPTY
		} else {
			id++
		}
		initVal(block, initValue)
		fs = append(fs, block...)
		i++
	}
	return fs
}

func initVal(b []int64, v int64) {
	for i := range b {
		b[i] = v
	}
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

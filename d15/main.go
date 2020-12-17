package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// start with starting numbs
// consider previous
// if that was first time, say 0
// if not say dfifference between previous and now

// slice of spoken nums in order
// hash of nums with latest turn spoken
// question: what is 2020th num spoken (idx by 1)

type numRecorder struct {
	order []int
	// map of number to index of when spoken (by 0)
	latest map[int]int
	counts map[int]int
}

func (nr *numRecorder) up(num int) int {
	nr.order = append(nr.order, num)
	prev := nr.latest[num]
	nr.latest[num] = len(nr.order) - 1
	nr.counts[num]++
	return prev
}

func (nr *numRecorder) play() {

	var prev = 0
	for i := len(nr.order) - 1; i < endIndex; i++ {
		cur := nr.order[i]
		//fmt.Printf("counts: %d,  latest:%d\n", nr.counts[cur], nr.latest[cur])
		if c, ok := nr.counts[cur]; ok && c == 1 {
			prev = nr.up(0)
		} else if ok && c > 1 {
			prev = nr.up(i - prev)
		} else {
			panic("not here")
		}
	}
}

const endIndex = 2020 - 1

func scanFile() numRecorder {
	nr := numRecorder{make([]int, 0), make(map[int]int), make(map[int]int)}
	var line string
	for {
		_, err := fmt.Scanf("%s\n", &line)
		if err == io.EOF {
			break
		}
		for _, v := range strings.Split(line, ",") {
			if iVal, err := strconv.ParseInt(v, 10, 64); err == nil {
				nr.up(int(iVal))
			} else {
				panic(err)
			}
		}
	}
	return nr
}

func main() {
	nr := scanFile()
	nr.play()
	//fmt.Printf("%+v\n", nr)
	fmt.Printf("order: %+v\n", nr.order)
	fmt.Printf("2020: %d\n", nr.order[endIndex])
}

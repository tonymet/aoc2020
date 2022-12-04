package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//  first col: A Rock, B for Paper, and C for Scissors.
//  second column,  X for Rock, Y for Paper, and Z for Scissors.

const (
	win  int = 6
	lose int = 0
	draw int = 3
)

func charVal(c rune) int {
	// a = 97, z = 123
	// A = 65, Z=91

	if c < 65 || c > 123 {
		panic(fmt.Sprintf("out of range: %d %s ", c, string(c)))
	}
	if c > 91 && c < 97 {
		panic(fmt.Sprintf("out of range: %d %s ", c, string(c)))
	}
	switch {
	case c >= 97:
		return int(c) - 96
	case c >= 65:
		return int(c) - 64 + 26
	default:
		panic("out of bounds")
	}
}

type uniqType map[rune]bool
type foundType []rune

func uniq(s string) (u uniqType) {
	u = make(uniqType)
	for _, v := range s {
		u[v] = true
	}
	return
}

func (m uniqType) String() string {
	s := make([]string, len(m))
	for v := range m {
		s = append(s, string(v))
	}
	return fmt.Sprintf("%v\n", strings.Join(s, " "))
}
func (m foundType) String() string {
	s := make([]string, len(m))
	for _, v := range m {
		s = append(s, string(v))
	}
	return fmt.Sprintf("%v\n", strings.Join(s, " "))
}

func part1() {
	// read line into string, with len
	// read first half into index
	// compare second half into index
	// lookup value in map and sum

	total := 0
	for {
		var cur string
		_, err := fmt.Scanf("%s\n", &cur)
		if err == io.EOF {
			break
		}
		fmt.Printf("%s: %d\n", cur, len(cur))
		end, mid := len(cur), len(cur)/2
		l, r := cur[0:mid], cur[mid:end]
		u := uniq(l)
		uniqR := uniq(r)
		var found = make(foundType, 0)
		for k := range uniqR {
			_, ok := u[k]
			if ok {
				found = append(found, k)
			}
		}
		fmt.Printf("%s\n%s \n\n", l, r)
		fmt.Printf("found: %v\n", found)
		for _, c := range found {
			total += charVal(c)
		}
	}
	fmt.Printf("total := %d\n", total)
}

func part2() {}

func main() {
	switch os.Getenv("PART") {
	case "2":
		part2()
	default:
		part1()
	}
}

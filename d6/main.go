package main

import (
	"fmt"
	"io"
)

type form struct {
	fields map[rune]int
}

func (f *form) parseForm(formLine string) {
	for _, c := range formLine {
		f.fields[c]++
	}
}

func (f *form) checkSize(expected int) int {
	c := 0
	for _, v := range f.fields {
		if v >= expected {
			c++
		}
	}
	return c
}

func main() {
	var (
		c, c2, groupSize int
		groupForm        form
	)
	groupForm.fields = make(map[rune]int)
	for {
		var line string
		n, err := fmt.Scanf("%s", &line)

		if n > 0 {
			groupForm.parseForm(line)
			groupSize++
		}

		if err != nil {
			// new record do check
			fmt.Printf("count: %d, groupSize: %d\n", len(groupForm.fields), groupSize)
			c += len(groupForm.fields)
			c2 += groupForm.checkSize(groupSize)
			groupForm.fields = make(map[rune]int)
			groupSize = 0
			if err == io.EOF {
				break
			}
		}
	}
	fmt.Printf("sum: %d \n, sum2: %d\n", c, c2)
}

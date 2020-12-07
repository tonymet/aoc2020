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
func main() {
	var (
		c         int
		groupForm form
	)
	groupForm.fields = make(map[rune]int)
	for {
		var line string
		_, err := fmt.Scanf("%s", &line)
		groupForm.parseForm(line)

		if err != nil {
			// new record do check
			fmt.Printf("count: %d\n", len(groupForm.fields))
			c += len(groupForm.fields)
			groupForm.fields = make(map[rune]int)
			if err == io.EOF {
				break
			}
		}
	}
	fmt.Printf("sum: %d \n", c)
}

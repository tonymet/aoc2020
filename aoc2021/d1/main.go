package main

import (
	"fmt"
	"io"
)

func main() {
	var cur, prev, greater int
	for _, err := fmt.Scanf("%d\n", &cur); err != io.EOF; _, err = fmt.Scanf("%d\n", &cur) {
		if cur > prev && prev != 0 {
			greater++
		}
		fmt.Printf("cur: %d, prev: %d, greater: %d\n", cur, prev, greater)
		prev = cur
	}
}

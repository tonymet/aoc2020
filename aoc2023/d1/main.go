package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func part1() {
	// var cur, i int
	var curLine string
	curNum := make([]string, 2)
	var sum int64
	for {
		_, err := fmt.Scanf("%s\n", &curLine)
		if err == io.EOF {
			break
		}
		// fmt.Printf("line = %s\n", curLine)
		var l, r int
		for l = 0; l <= len(curLine)-1; l++ {
			if curLine[l] >= '0' && curLine[l] <= '9' {
				// fmt.Printf("curLine-L: %s\n", string(curLine[l]))
				curNum[0] = string(curLine[l])
				break
			}
		}
		for r = len(curLine) - 1; r >= 0; r-- {
			if curLine[r] >= '0' && curLine[r] <= '9' {
				// fmt.Printf("curLine-R: %s\n", string(curLine[r]))
				curNum[1] = string(curLine[r])
				break
			}
		}
		// fmt.Printf("curNum: %x\n", curNum)
		fullNum := strings.Join(curNum, "")
		// fmt.Printf("fullNum: %s\n", fullNum)

		if i, err := strconv.ParseInt(fullNum, 10, 64); err == nil {
			fmt.Printf("i: %d\n", i)
			sum += i
		} else {
			panic(err)
		}
	}
	fmt.Printf("sum: %d\n", sum)
}
func main() {
	switch os.Getenv("PART") {
	case "2":
		// part2()
	default:
		part1()
	}
}

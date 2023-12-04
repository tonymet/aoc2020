package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var part = 1

var numMap = map[string]string{
	"zero":  "0o",
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "4",
	"five":  "5e",
	"six":   "6",
	"seven": "7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func replaceNumsInString(s string) string {
	for k := range numMap {
		s = strings.Replace(s, k, numMap[k], -1)
	}
	return s
}

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
		curLine = replaceNumsInString(curLine)
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
	if stdin := os.Getenv("STDIN"); len(stdin) != 0 {
		stdinFile, err := os.Open(stdin)
		if err != nil {
			panic(err)
		}
		os.Stdin = stdinFile
	}
	switch os.Getenv("PART") {
	case "2":
		part = 2
		fallthrough
	default:
		part1()
	}
}

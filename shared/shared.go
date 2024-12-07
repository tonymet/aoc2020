package shared

import (
	"bufio"
	"fmt"
	"regexp"
)

var (
	Silent bool
)

func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

type rec []int

func Log(f string, val ...any) {
	if Silent {
		return
	}
	fmt.Printf(f, val...)
}

// generate a bufio splitter based on the regex
func SplitterFactory(re *regexp.Regexp) func([]byte, bool) (int, []byte, error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		loc := re.FindIndex(data)
		if loc == nil {
			// try again
			if !atEOF {
				return 0, nil, nil
			}
			return 0, nil, bufio.ErrFinalToken
		}
		return loc[1], data[loc[0]:loc[1]], nil
	}
}

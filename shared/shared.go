package shared

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
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

func Powerset(data [][]int, f func([][]int)) {
	n := len(data)
	// Total number of subsets is 2^n
	numSubsets := 1 << n

	for i := 0; i < numSubsets; i++ {
		var subset [][]int
		for j := 0; j < n; j++ {
			// Check if the j-th bit is set in mask i
			if (i>>j)&1 == 1 {
				subset = append(subset, data[j])
			}
		}
		f(subset)
	}
}

func LineProcessor(in io.Reader, processor func(io.Reader)) {
	lineScanner := bufio.NewScanner(in)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		lineReader := strings.NewReader(line)
		processor(lineReader)
	}
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

func stringInts(s string) (r []int64, err error) {
	r = make([]int64, 0, 40)
	for _, c := range s {
		v, err := strconv.ParseInt(string(c), 10, 64)
		if err != nil {
			panic(err)
		}
		r = append(r, v)
	}
	return r, nil
}

package main

import (
	"fmt"
	"io"
	"os"
)

type LetterWindow struct {
	windowScope []byte
	windowCount map[byte]int
}

func sumCounts(l map[byte]int) (sum int) {
	fmt.Printf("%+v\n", l)
	for k := range l {
		sum += l[k]
	}
	return
}

func testWindow(lw LetterWindow) bool {
	if len(lw.windowCount) < 4 {
		return false
	}
	for _, v := range lw.windowCount {
		if v != 1 {
			return false
		}
	}
	return true
}

func parseAndSolve() {

	var lw LetterWindow
	lw.windowCount = make(map[byte]int)
	lw.windowScope = make([]byte, 0, 4)
	b := make([]byte, 1)
	var row, col int
	for {
		_, err := os.Stdin.Read(b)
		if err == io.EOF {
			break
		}
		if b[0] == '\n' {
			fmt.Printf("%+v\n", lw.windowCount)
			col = 0
			row++
			lw.windowCount = make(map[byte]int)
		}
		lw.windowCount[b[0]]++
		lw.windowScope = append(lw.windowScope, b[0])
		if col >= 3 {
			// test window
			// if not adequate, update
			l := lw.windowScope[0]
			if testWindow(lw) {
				fmt.Printf("lw: %+v\n", lw)
				fmt.Printf("found: col:%d \n", col+1)
				break
			} else {
				if lw.windowCount[l] > 1 {
					lw.windowCount[l]--
				} else {
					delete(lw.windowCount, l)
				}
				lw.windowScope = lw.windowScope[1:]
				if len(lw.windowScope) != 3 {
					panic("windowscope insonsidtent")

				}
			}
		}
		col++
	}
}
func main() {
	if stdin := os.Getenv("STDIN"); len(stdin) != 0 {
		stdinFile, err := os.Open(stdin)
		if err != nil {
			panic(err)
		}
		os.Stdin = stdinFile
	}
	parseAndSolve()

}

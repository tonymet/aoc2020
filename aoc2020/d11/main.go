package main

import (
	"fmt"
	"io"
)

const (
	maxX = 100
	maxY = 91
)

type gridType [92][]byte

func scanFile() (gridType, error) {
	var gr gridType
	for y, err := 0, error(nil); err != io.EOF; y++ {
		//var cur string
		_, err = fmt.Scanln(&gr[y])
		fmt.Println(string(gr[y]))
	}
	return gr, nil
}
func main() {
	_, err := scanFile()
	if err != nil {
		panic("err reading file")
	}

}

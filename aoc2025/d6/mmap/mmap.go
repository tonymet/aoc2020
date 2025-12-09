package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"syscall"
)

type fileParam struct {
	cols, rows int
}

var (
	files map[string]fileParam = map[string]fileParam{
		"sample": {cols: 16, rows: 3},
		"input":  {cols: 3753, rows: 4},
	}
	filetype string
	filename string
)

// A function to read the file into memory using mmap.
// Returns a byte slice representing the memory-mapped file content.
func mmapFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}
	size := int(fi.Size())
	data, err := syscall.Mmap(
		int(file.Fd()),
		0, // offset (start from the beginning of the file)
		size,
		syscall.PROT_READ,
		syscall.MAP_SHARED,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to mmap file: %w", err)
	}
	return data, nil
}

func init() {
	flag.StringVar(&filename, "f", "input.txt", "filename")
	flag.StringVar(&filetype, "type", "sample", "type of file")
}
func byteToInt(charVal [][]byte) (ret []int) {
	ret = make([]int, 0, len(charVal))
	for _, v := range charVal {
		//[]byte
		val, err := strconv.ParseInt(strings.TrimSpace(string(v)), 10, 64)
		if err != nil {
			panic(err)
		}
		ret = append(ret, int(val))
	}
	return

}

func forceInt(v []byte) int {
	val, err := strconv.ParseInt(strings.TrimSpace(string(v)), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(val)
}

func dividerCol(col []byte) (empty bool) {
	return strings.TrimSpace(string(col)) == ""
}

func main() {
	flag.Parse()
	rows, cols := files[filetype].rows, files[filetype].cols
	data, err := mmapFile(filename)
	if err != nil {
		fmt.Println("Mmap error:", err)
		return
	}
	opsArr := strings.Fields(string(data[cols*rows : cols*(rows+1)-1]))
	// RtoL order
	slices.Reverse(opsArr)
	fmt.Printf("%s\n", opsArr)
	curCol := 0
	stackInt := make([]int, 0)
	sumVal := 0
	// read matrix L to R, char-col at a time
	for pos := cols - 1; pos >= 0; pos-- {
		val := make([]byte, 0, rows)
		for j := rows; j > 0; j-- {
			val = append(val, data[(cols*(rows-j))+pos])
		}
		//	fmt.Printf("%s \n", string(val))
		if val[0] == '\n' {
			continue
		}
		if !dividerCol(val) {
			stackInt = append(stackInt, forceInt(val))
		}
		if dividerCol(val) || pos == 0 {
			opVal := 0
			switch opsArr[curCol] {
			case "*":
				opVal = 1
				for _, v := range stackInt {
					opVal *= v
				}
			case "+":
				opVal = 0
				for _, v := range stackInt {
					opVal += v
				}
			}
			stackInt = make([]int, 0)
			sumVal += opVal
			curCol++
		}
	}
	fmt.Printf("sumVal: %d\n", sumVal)

	// Ensure to unmap the memory when finished
	defer func() {
		if err := syscall.Munmap(data); err != nil {
			fmt.Println("Munmap error:", err)
		}
	}()
}

package main

import (
	"errors"
	"fmt"
	"io"
	"sort"
)

type recordType []int

var ErrorNotFound = errors.New("Not found")

var records recordType

func checkTargetSum(target int, poolRef []int) bool {
	pool := make([]int, len(poolRef))
	copy(pool, poolRef)
	sort.Ints(pool)
	l, r := 0, len(pool)-1
	for pool[l]+pool[r] != target && l != r {
		if pool[l]+pool[r] > target {
			r--
		} else {
			l++
		}
	}
	if l == r {
		return false
	}
	return true
}
func scanFile() {
	records = make(recordType, 0)
	for {
		var cur int
		_, err := fmt.Scanf("%d", &cur)
		if err == io.EOF {
			//fmt.Printf("end of file")
			return
		}
		records = append(records, cur)
	}
}

func findFalseTarget(records recordType) (int, error) {
	// target = r + 1
	for l, r := 0, 24; r != len(records)-2; l, r = l+1, r+1 {
		var target = records[r+1]
		result := checkTargetSum(target, records[l:r+1])
		//fmt.Printf("%t %d : %+v\n", result, target, records[l:r+1])
		if !result {
			return target, nil
		}
	}
	return 0, ErrorNotFound
}

func main() {
	scanFile()
	target, err := findFalseTarget(records)
	if err != nil {
		panic("Error, target not found \n")
	}
	fmt.Printf("target: %d\n ", target)
}

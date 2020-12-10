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

func findContig(target int, pool []int) ([]int, error) {
	// start at R
	// push R
	// sum queue
	//if sum > target
	// shift off queue
	// push R-1
	// dead case is if len(queue) == len(pool) or if r = 0

	testStack := make([]int, 0, len(pool))
	l, r := len(pool)-2, len(pool)-1
	testStack = append(testStack, pool[l])
	testStack = append(testStack, pool[r])
	for sumSlice(testStack) != target && r != 1 {
		for sumSlice(testStack) < target && l != 0 {
			l--
			if pool[l] == target {
				continue
			}
			testStack = testStack[0 : len(testStack)+1]
			copy(testStack[1:], testStack)
			testStack[0] = pool[l]
		}
		for sumSlice(testStack) > target && r != 1 {
			r--
			if pool[r] == target {
				continue
			}
			testStack = testStack[0 : len(testStack)-1]
		}
	}
	if sumSlice(testStack) == target {
		return testStack, nil
	}
	return nil, errors.New("not Found")
}

func minSlice(s []int) int {
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func maxSlice(s []int) int {
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

func sumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
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
		panic("Error, findFalseTarget target not found \n")
	}
	fmt.Printf("target: %d\n ", target)

	testStack, err2 := findContig(756008079, records)
	if err2 != nil {
		panic("Error, findContig target not found \n")
	}
	max, min := maxSlice(testStack), minSlice(testStack)
	fmt.Printf("findContig target: %+v\n, min:%d, max: %d, sum: %d ", testStack, min, max, max+min)

}

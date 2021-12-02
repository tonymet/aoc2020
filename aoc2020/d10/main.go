package main

import (
	"errors"
	"fmt"
	"io"
	"sort"
)

type recordType []int

func scanFile() (recordType, error) {
	records := make(recordType, 0)
	for {
		var cur int
		_, err := fmt.Scanf("%d", &cur)
		if err == io.EOF {
			//fmt.Printf("end of file")
			return records, nil
		} else if err != nil {
			return recordType{}, errors.New("Error reading file")
		}
		records = append(records, int(cur))
	}
}

func setupChain(records recordType) recordType {
	records = append(records, 0)
	sort.Ints(records)
	max := records[len(records)-1]
	// add +3 for device joltage
	records = append(records, max+3)
	return records

}

func findChains(records recordType) map[int]int {
	// add zero joltage
	deltaMap := make(map[int]int)
	records = setupChain(records)
	for i, prev := 1, 0; i < len(records); prev, i = i, i+1 {
		deltaMap[records[i]-records[prev]]++
	}
	fmt.Printf("records: %+v\n", records)
	fmt.Printf("deltaMap: %+v\n", deltaMap)
	fmt.Printf("product: %d\n", deltaMap[3]*deltaMap[1])
	max := records[len(records)-1]
	fmt.Printf("max: %d\n", max)
	return deltaMap
}

func genCounts(records recordType) int {

	// setup
	records = setupChain(records)
	max := records[len(records)-1]
	trackCounts := make(map[int]int)
	trackCounts[0] = 1

	recordIndex := make(map[int]bool)
	// make index
	for _, v := range records {
		recordIndex[v] = true
	}
	for _, cur := range records {
		for _, v := range []int{1, 2, 3} {
			if recordIndex[cur+v] {
				trackCounts[cur+v] += trackCounts[cur]
			}
		}
	}
	fmt.Printf("max: %d\n", max)
	return trackCounts[max]
}

func main() {
	records, _ := scanFile()
	records2 := make(recordType, len(records))
	copy(records2, records)
	findChains(records)
	//records2 := []int{0, 1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24, 25, 28, 31, 32, 33, 34, 35, 38, 39, 42, 45, 46, 47, 48, 49, 52}
	//fmt.Printf("genCounts test: %d  \n", genCounts(records2))
	fmt.Printf("genCounts actual: %d  \n", genCounts(records2))
}

package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// structure
// read times. sort
// from cur ts +1 if % bustime == 0 . else nexto
// DS : map [bus] --> time

type scheduler struct {
	targetTime    int64
	runningBusses []int64
	checkedTimes  map[int64]map[int64]bool
}

func stringsToInts(s []string, len int64) []int64 {
	r := make([]int64, 0)
	for _, e := range s {
		if i, err := strconv.ParseInt(e, 10, 32); err == nil {
			r = append(r, int64(i))
		} else {
			r = append(r, -1)
		}
	}
	return r
}

func scanFile() scheduler {
	theSchedule := scheduler{
		0,
		make([]int64, 0),
		make(map[int64]map[int64]bool),
	}
	// read Time
	n, err := fmt.Scanf("%d", &theSchedule.targetTime)
	if err != nil || n != 1 {
		panic("err reading targetTiem")
	}
	// csv
	var timeLine string
	n, err = fmt.Scanf("%s", &timeLine)
	if err != nil || n != 1 {
		panic("err reading timeLine")
	}
	theSchedule.runningBusses = stringsToInts(strings.Split(timeLine, ","), 0)
	fmt.Printf("runningInts %+v\n", theSchedule.runningBusses)

	if err != nil || n != 1 {
		panic("err reading csv")
	}
	fmt.Printf("theSchedule: %+v", theSchedule)
	return theSchedule
}

func sortInt64(s []int64) {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
}

func part1(theSchedule scheduler) (int64, int64) {
	sortInt64(theSchedule.runningBusses)
	for t := theSchedule.targetTime; true; t++ {
		// start at t
		for _, bus := range theSchedule.runningBusses {
			if bus == -1 {
				continue
			}
			if t%bus == 0 {
				return bus, (t - theSchedule.targetTime)
			}
		}
	}
	panic("end")
}

type timeIndexType struct {
	offset, bus int64
}

func product(theSchedule scheduler) int64 {
	product := int64(1)
	for _, v := range theSchedule.runningBusses {
		if v == -1 {
			continue
		}
		product *= v
	}
	return product
}
func part2(theSchedule scheduler) int64 {
	// for each busID
	timeIndex := make([]timeIndexType, 0)
	for k, v := range theSchedule.runningBusses {
		if v == -1 {
			continue
		}
		timeIndex = append(timeIndex, timeIndexType{int64(k), v})
	}
	fmt.Printf("timeINdex: %+v\n", timeIndex)
	var t int64
	start := int64(1e14)
	product := product(theSchedule)
	for t = start; true; t++ {

		if t%product == 0 {
			return t
		}
		/*
			for i, slot := range timeIndex {
				// if no match, break
				if t%int64(slot.bus) != int64(slot.offset) {
					break
				}
				if i+1 == len(timeIndex) {
					return t
				}
			}
		*/
		if t%1e7 == 0 {
			fmt.Printf("tested %d\n", t)
		}
	}
	return -1
}

func main() {
	fmt.Printf("hello\n")
	var pFlag = flag.Int("p", 1, "1 or 2")
	flag.Parse()
	if *pFlag == 1 {
		// idenitfy closes buss.
		// subtract departer from current time
		// multiply by bus ID
		theSchedule := scanFile()
		bus, delta := part1(theSchedule)
		fmt.Printf("part 1\n")
		fmt.Printf("bus: %d, delta %d, product: %d\n", bus, delta, bus*delta)
	} else if *pFlag == 2 {
		fmt.Printf("part 2\n")
		theSchedule := scanFile()
		time := part2(theSchedule)
		fmt.Printf("time: %d\n", time)

	} else {
		panic("wrong flag")
	}

}

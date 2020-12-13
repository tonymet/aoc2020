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
	targetTime    int
	runningBusses []int
	checkedTimes  map[int]map[int]bool
}

func stringsToInts(s []string, len int) []int {
	r := make([]int, 0)
	for _, e := range s {
		if i, err := strconv.ParseInt(e, 10, 32); err == nil {
			r = append(r, int(i))
		}
	}
	return r
}

func scanFile() scheduler {
	theSchedule := scheduler{
		0,
		make([]int, 0),
		make(map[int]map[int]bool),
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

func part1(theSchedule scheduler) (int, int) {
	sort.Ints(theSchedule.runningBusses)
	for t := theSchedule.targetTime; true; t++ {
		// start at t
		for _, bus := range theSchedule.runningBusses {
			if t%bus == 0 {
				return bus, (t - theSchedule.targetTime)
			}
		}
	}
	panic("end")
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

	} else {
		panic("wrong flag")
	}

}

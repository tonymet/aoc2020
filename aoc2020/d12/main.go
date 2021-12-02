package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
)

type recordType []cmd
type cmd struct {
	cmd   string
	value int
}

type coord struct {
	x, y int
}

type ship struct {
	counts   coord
	waypoint coord
}

type positionTracker struct {
	movements map[string]int
	bearing   string
}

var coordinatesToDegrees = map[string]int{"N": 0, "E": 90, "S": 180, "W": 270}
var degreesToCoordinates = map[int]string{0: "N", 90: "E", 180: "S", 270: "W"}

func rotate(cmd string, c coord) coord {
	switch cmd {
	case "L90":
		fallthrough
	case "R270":
		c.x, c.y = -1*c.y, c.x
	case "L180":
		fallthrough
	case "R180":
		c.x, c.y = -1*c.x, -1*c.y
	case "R90":
		fallthrough
	case "L270":
		c.x, c.y = c.y, -1*c.x
	}
	return c
}

func intAbs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func newBearing(curBearing string, turnCmd cmd) string {
	deg, ok := coordinatesToDegrees[curBearing]
	if !ok {
		panic(errors.New("coord not found"))
	}
	var newDeg int
	switch turnCmd.cmd {
	case "R":
		newDeg = (deg + turnCmd.value) % 360
	case "L":
		newDeg = intAbs(deg-turnCmd.value+360) % 360
	default:
		panic("wrong cmd")
	}
	newCoord, ok := degreesToCoordinates[newDeg]
	if !ok {
		panic(fmt.Errorf("bad newDeg: %d, oldDeg: %d", newDeg, deg))
	}
	return newCoord
}

func execMoves(records recordType) (positionTracker, error) {
	// inintial position
	tracker := positionTracker{
		make(map[string]int),
		"E",
	}
	for _, move := range records {
		switch move.cmd {
		case "L":
			fallthrough
		case "R":
			tracker.bearing = newBearing(tracker.bearing, move)
		case "N":
			fallthrough
		case "S":
			fallthrough
		case "E":
			fallthrough
		case "W":
			tracker.movements[move.cmd] += move.value
		case "F":
			tracker.movements[tracker.bearing] += move.value
		default:
			panic("bad cmd")
		}
	}
	return tracker, nil
}

func scanFile() (recordType, error) {
	results := make(recordType, 0)
	for {
		var curCmd cmd
		_, err := fmt.Scanf("%1s%d", &curCmd.cmd, &curCmd.value)
		if err == io.EOF {
			return results, nil
		}
		results = append(results, curCmd)
	}
}

func part2() (coord, error) {
	theShip := ship{coord{0, 0}, coord{10, 1}}

	var err error
	var curCmd cmd
	for {
		_, err = fmt.Scanf("%1s%d", &curCmd.cmd, &curCmd.value)
		fmt.Printf("curCmd: %+v\n", curCmd)
		if err == io.EOF {
			fmt.Printf("end of file \n")
			break
		} else if err != nil {
			return coord{0, 0}, errors.New("error reading file")
		}
		switch curCmd.cmd {
		case "L":
			fallthrough
		case "R":
			theShip.waypoint = rotate(fmt.Sprintf("%s%d", curCmd.cmd, curCmd.value), theShip.waypoint)
		case "N":
			theShip.waypoint.y += curCmd.value
		case "S":
			theShip.waypoint.y -= curCmd.value
		case "E":
			theShip.waypoint.x += curCmd.value
		case "W":
			theShip.waypoint.x -= curCmd.value
		case "F":
			theShip.counts.x += theShip.waypoint.x * curCmd.value
			theShip.counts.y += theShip.waypoint.y * curCmd.value
		default:
			panic("bad cmd")
		}
		fmt.Printf("theShip: %+v\n", theShip)
	}
	return theShip.counts, nil
}

func sumFinal(c coord) int {
	return intAbs(c.x) + intAbs(c.y)
}

func main() {
	fmt.Printf("hello\n")
	var pFlag = flag.Int("p", 1, "1 or 2")
	flag.Parse()
	if *pFlag == 1 {
		results, err := scanFile()
		if err != nil {
			panic(err)
		}
		//fmt.Printf("results: %+v\n", results)
		tracker, err2 := execMoves(results)
		if err2 != nil {
			panic(err)
		}
		fmt.Printf("tracker: %+v\n", tracker)
		netNorth := tracker.movements["N"] - tracker.movements["S"]
		netEast := tracker.movements["E"] - tracker.movements["W"]
		manhattan := intAbs(netNorth) + intAbs(netEast)
		fmt.Printf("N/S pos : %d | E/W pos :%d,  manhattan :%d", netNorth, netEast, manhattan)
	} else if *pFlag == 2 {
		part2Counts, err3 := part2()
		if err3 != nil {
			panic(err3)
		}
		fmt.Printf("part2: %+v, abs: %d", part2Counts, sumFinal(part2Counts))

	} else {
		panic("wrong flag")
	}
}

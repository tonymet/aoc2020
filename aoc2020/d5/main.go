package main

import (
	"flag"
	"fmt"
	"io"
	"sort"
)

var part2 *bool
var allSeats []int

func scanFile() {
	var (
		seatPath string
		maxID    int = 0
	)
	allSeats = make([]int, 0)

	for {
		n, err := fmt.Scanf("%10s\n", &seatPath)
		if n != 1 {
			fmt.Errorf("error reading file")
		}
		seatID := seatID(seatPath)
		if seatID > maxID {
			maxID = seatID
		}
		if *part2 {
			fmt.Printf("%d seatID\n", seatID)
			allSeats = append(allSeats, seatID)
		} else {
			fmt.Printf("seatPath: %s, seatID: %d\n", seatPath, seatID)
		}
		if err == io.EOF {
			break
		}
	}
	sort.Ints(allSeats)
	fmt.Printf("MaxID: %d", maxID)

}
func getRow(rowid string) int {
	l, r := 0, 128
	d, i := 64, 0

	for d >= 1 {
		//fmt.Printf("op, r, l , d , i : %s, %s, %s, %s, %s\n", string(rowid[i]), r, l, d, i)
		cur := rowid[i]
		switch cur {
		case 'F':
			r = r - d
		case 'B':
			l = l + d
		default:
			fmt.Errorf("unexpected: %s", cur)
		}
		d = d >> 1
		i++
	}
	return l

}

func getCol(colid string) int {
	l, r := 0, 7
	d, i := 4, 0

	for d >= 1 {
		cur := colid[i]
		switch cur {
		case 'L':
			r = r - d
		case 'R':
			l = l + d
		default:
			fmt.Errorf("unexpected: %s", cur)
		}
		d = d >> 1
		i++
		//fmt.Printf("r, l , d , i : %s, %s, %s, %s\n", r, l, d, i)
	}
	return l
}

func seatID(seatPath string) int {
	row := getRow(seatPath[0:7])
	col := getCol(seatPath[7:10])
	return (row * 8) + col
}

func findMissing() {
	i := allSeats[0]
	for _, v := range allSeats {
		if i != v {
			fmt.Printf("Expected %d, was missing, %d is here\n", i, v)
			i = v
		}
		i++
	}

}
func main() {
	part2 = flag.Bool("part2", false, "part2")
	flag.Parse()
	scanFile()
	findMissing()
}

package main

import "fmt"

func getRow(rowid string) int {
	l, r := 0, 128
	d, i := 64, 0

	for d >= 1 {
		fmt.Printf("op, r, l , d , i : %s, %s, %s, %s, %s\n", string(rowid[i]), r, l, d, i)
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
func main() {
	fmt.Println("hello")
	row := getRow("FFFBBBF")
	col := getCol("RRR")
	seatid := seatID("BBFFBBFRLL")
	fmt.Printf("row: %d, col: %d, seatID: %s\n", row, col, seatid)

}

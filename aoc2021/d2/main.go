package main

import (
	"fmt"
	"io"
	"os"
)

type threed struct {
	x, y, z int
}

func part2() {
}
func part1() {
	var (
		dir string
		val int
		th  threed
	)

	for _, err := fmt.Scanf("%s %d\n", &dir, &val); err != io.EOF; _, err = fmt.Scanf("%s %d\n", &dir, &val) {
		switch dir {
		case "forward":
			th.z += val
		case "down":
			th.y -= val
		case "up":
			th.y += val
		default:
			panic("out of bounds:")
		}
		fmt.Printf("threed : %+v\n", th)
	}
	fmt.Printf("solution : %d\n", th.y*-1*th.z)

}
func main() {
	switch os.Getenv("PART") {
	case "2":
		part2()
	default:
		part1()
	}
}

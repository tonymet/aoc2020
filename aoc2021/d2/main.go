package d2

import (
	"fmt"
	"io"
	"os"
)

type threed struct {
	x, y, z int
}

func part2() {
	var (
		dir string
		val int
		aim threed
		pos threed
	)

	for _, err := fmt.Scanf("%s %d\n", &dir, &val); err != io.EOF; _, err = fmt.Scanf("%s %d\n", &dir, &val) {
		switch dir {
		case "forward":
			pos.z += val
			pos.y += val * aim.y
		case "down":
			aim.y -= val
		case "up":
			aim.y += val
		default:
			panic("out of bounds:")
		}
		fmt.Printf("aim : %+v, pos: %+v\n", aim, pos)
	}
	fmt.Printf("solution : %d\n", pos.y*-1*pos.z)

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

package shared

var powers = []int64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000}

func Pow10(e int) int64 {
	return powers[e]
}

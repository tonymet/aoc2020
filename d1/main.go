package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	nums := make([]int64, 0)
	for i := 0; stdin.Scan(); i++ {
		v, err := strconv.ParseInt(stdin.Text(), 10, 64)
		if err == nil {
			nums = append(nums, v)
		}
	}

	l, r := 0, len(nums)-1
	for nums[l]+nums[r] != 2020 {
		fmt.Printf("l %d, r %d, sum %d\n", l, r, nums[l]+nums[r])
		if nums[l]+nums[r] > 2020 {
			r--
		} else {
			l++
		}
	}
	fmt.Printf("nums: %d, %d   mult: %d\n", nums[l], nums[r], nums[l]*nums[r])

	fmt.Printf("%x", nums)
	if err := stdin.Err(); err != nil {
		log.Println(err)
	}
}

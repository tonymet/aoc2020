package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum2(nums []int64) {
	if len(nums) == 0 {
		return
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

}

func sum3(nums []int64) {
	if len(nums) == 0 {
		return
	}
outer:
	for _, i := range nums {
		for _, j := range nums {
			for _, k := range nums {
				if i+j+k == 2020 {
					fmt.Printf("%d + %d + %d = %d ,  product= %d ", i, j, k, i+j+k, i*j*k)
					break outer
				}

			}
		}
	}

}
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	nums := make([]int64, 0)
	for i := 0; stdin.Scan(); i++ {
		v, err := strconv.ParseInt(stdin.Text(), 10, 64)
		if err == nil {
			nums = append(nums, v)
		}
	}
	sum2(nums)
	sum3(nums)

	fmt.Printf("%x", nums)
	if err := stdin.Err(); err != nil {
		log.Println(err)
	}
}

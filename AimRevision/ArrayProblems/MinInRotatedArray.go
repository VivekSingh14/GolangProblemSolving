package main

import (
	"fmt"
	"math"
)

func main14() {
	arr := []int{2, 1}

	minInRotatedArray(arr)
}

func minInRotatedArray(nums []int) {
	l := 0
	h := len(nums) - 1
	min := math.MaxInt
	for l <= h {
		if nums[l] < min {
			min = nums[l]
			l++
		} else if nums[h] < min {
			min = nums[h]
			h--
		} else {
			l++
			h--
		}
	}

	fmt.Println(min)
}

package main

import "fmt"

// binary search without using recursion and insert

func main12() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		v := nums[m]
		switch {
		case v < target:
			l = m + 1
		case v > target:
			r = m - 1
		default:
			return m
		}
	}
	return l
}

package main

import (
	"fmt"
)

func main12() {
	arr := []int{1, -2, -3, 0, 7, -8, -2}

	maxProduct(arr)
}

func maxProduct(nums []int) {
	max_ending_here := nums[0]
	min_ending_here := nums[0]
	max_so_far := nums[0]
	for i := 1; i < len(nums); i++ {
		temp := max(max(nums[i], nums[i]*max_ending_here), nums[i]*min_ending_here)
		min_ending_here = min(min(nums[i], nums[i]*max_ending_here), nums[i]*min_ending_here)
		max_ending_here = temp
		max_so_far = max(max_so_far, max_ending_here)
	}
	fmt.Println(max_so_far)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

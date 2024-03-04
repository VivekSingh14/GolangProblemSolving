package main

import "fmt"

/* --
Prefix Sum algorithm
-- */

func main11() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
}

func runningSum(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] + nums[i]
	}
	return res
}

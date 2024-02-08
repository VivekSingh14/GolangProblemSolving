package main

import "fmt"

func main37() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(minJumps(nums))
}

func minJumps(nums []int) bool {

	farthest := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		if nums[i] == 0 && i < n-1 && i == farthest {
			return false
		}
	}

	return true
}

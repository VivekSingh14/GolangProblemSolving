package main

import "fmt"

func main4() {
	A := []int{1, 2, 3, 3, 3, 4, 5, 5}

	fmt.Println(removeDuplicates(A))
}
func removeDuplicates(nums []int) int {
	slow := 0
	fast := 0
	res := 0

	for fast < len(nums) {
		if nums[fast] == nums[slow] {
			fast++
		} else if nums[fast] != nums[slow] {
			nums[slow+1] = nums[fast]
			slow++
			fast++
			res++
		}
	}
	return res
}

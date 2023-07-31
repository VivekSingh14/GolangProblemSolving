package main

import "fmt"

func main1() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	res := 0
	fast := 1
	slow := 0

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
	fmt.Println(nums)
	return res + 1
}

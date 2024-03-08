package main

import "fmt"

func main15() {
	nums := []int{2, 3, 1, 0}
	//10, 14, 6, 13
	//13, 3, -1 , 7

	fmt.Println(waysToSplitArray(nums))
}

func waysToSplitArray(nums []int) int {
	count := 0
	prefix := make([]int, len(nums))
	postfix := make([]int, len(nums))
	prefix[0] = nums[0]
	postfix[len(nums)-1] = nums[len(nums)-1]

	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		postfix[i] = postfix[i+1] + nums[i]
	}

	for i := 0; i < len(nums)-1; i++ {
		if prefix[i] >= postfix[i+1] {
			count++
		}
	}

	return count
}

package main

import "fmt"

func main() {
	arr := []int{1, 1, 2}

	num := removeDuplicates(arr)

	fmt.Println(num)

}

func removeDuplicates(nums []int) int {
	slow := 0
	fast := 1
	count := 0
	for slow < fast && fast < len(nums) {
		if nums[slow] == nums[fast] {
			fast = fast + 1
		} else {
			slow = fast
			fast = fast + 1
			count = count + 1
		}
	}
	return count + 1
}

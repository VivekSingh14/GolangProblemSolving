package main

import "fmt"

func main6() {
	nums := []int{1, 3, 5, 8, 9, 2, 6, 7, 6, 8, 9}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {

	count := 0
	max := len(nums) - 1

	for max != 0 {
		for i := 0; i < max; i++ {
			if i+nums[i] >= max {
				max = i
				count++
				break
			}
		}
	}
	return count
}

package main

import "fmt"

func main5() {

	arr := []int{1, 0, 3, 4, 0, 6, 9, 0, 11}
	//arr := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
	fmt.Println(moveZeroes(arr))
}

func moveZeroes(nums []int) []int {

	if len(nums) <= 2 && nums[len(nums)-1] == 0 {
		return nums
	}

	if len(nums) == 2 && nums[0] == 0 {
		nums[0] = nums[1]
		nums[1] = 0
		return nums
	} else if len(nums) == 2 {
		return nums
	} else {

		j := 1
		i := 0
		for j < len(nums) && i < len(nums) {
			if nums[i] == 0 {
				nums[i] = nums[j]
				nums[j] = 0
				j = j + 1
			} else {
				i = i + 1
				j = i + 1
			}
		}
	}
	return nums

}

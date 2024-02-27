package main

import "fmt"

func main5() {
	nums := []int{1, 2, 3, 1}

	fmt.Println(containsDuplicate(nums))
}

// all elements are positive
func containsDuplicate(nums []int) bool {
	res := make([]bool, 10)
	for i := 0; i < len(nums); i++ {
		if res[nums[i]] {
			return true
		} else {
			res[nums[i]] = true
		}
	}
	return false
}

// other case, considering negative numbers or numbers greater than 10
func containsDuplicateWithMap(nums []int) bool {
	map1 := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if map1[nums[i]] {
			return true
		}
		map1[nums[i]] = true
	}

	return false
}

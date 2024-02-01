package main

import "fmt"

func main32() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))

}

func containsDuplicate(nums []int) bool {
	map1 := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if _, ok := map1[nums[i]]; ok {
			return true
		}
		map1[nums[i]] = true
	}

	return false
}

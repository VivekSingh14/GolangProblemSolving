package main

import "fmt"

func main3() {
	nums := []int{4, 1, 2, 1, 2}

	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	map1 := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if map1[nums[i]] == 0 {
			map1[nums[i]] = 1
		} else {
			value := map1[nums[i]]
			map1[nums[i]] = value + 1
		}
	}

	for key, value := range map1 {
		if value == 1 {
			return key
		}
	}
	return 0
}

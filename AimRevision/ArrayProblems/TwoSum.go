package main

import "fmt"

func main6() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {

	var res []int

	map1 := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if value, ok := map1[target-nums[i]]; ok {
			res = append(res, value)
			res = append(res, i)
		}
		map1[nums[i]] = i
	}
	return res
}

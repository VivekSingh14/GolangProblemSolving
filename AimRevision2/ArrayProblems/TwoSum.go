/*
find out sum of pair in an array equals to target
array may or may not be in a sorted order
*/
package main

import "fmt"

func main1() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(targetSum(nums, target))
	fmt.Println(targetSum2(nums, target))
}

//In this function, Two pointer algo is used considering array is sorted.
func targetSum(nums []int, target int) []int {

	start := 0
	end := len(nums) - 1
	var res []int

	for start < end {
		if nums[start]+nums[end] == target {
			res = append(res, start)
			res = append(res, end)
			break
		} else if nums[start]+nums[end] < target {
			start++
		} else {
			end--
		}
	}
	return res
}

// in this function, map is used to find out target sum considering array is not sorted.
func targetSum2(nums []int, target int) []int {
	var res []int
	map1 := make(map[int]int)

	for i, value := range nums {
		if k, ok := map1[target-value]; ok {
			res = append(res, k)
			res = append(res, i)
		}
		map1[value] = i
	}
	return res
}

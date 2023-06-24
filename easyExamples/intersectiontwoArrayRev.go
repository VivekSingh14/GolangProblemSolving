package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{3, 4, 6, 2, 8, 5, 9}
	nums2 := []int{6, 3, 2, 7, 5, 1}

	fmt.Println(intersectionOfArray(nums1, nums2))
}

func intersectionOfArray(nums1 []int, nums2 []int) []int {
	nums1 = sortData(nums1)
	nums2 = sortData(nums2)
	max := 0
	if len(nums1) < len(nums2) {
		max = len(nums1)
	} else {
		max = len(nums2)
	}
	var res []int
	i := 0
	j := 0

	for i < max && j < max {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i = i + 1
			j = j + 1
		} else if nums1[i] < nums2[j] {
			i = i + 1
		} else {
			j = j + 1
		}
	}

	return res
}

func sortData(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}

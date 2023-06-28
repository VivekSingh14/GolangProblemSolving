package main

import (
	"fmt"
	"sort"
)

func main16() {
	nums1 := []int{1, 2}
	nums2 := []int{1, 1}

	fmt.Println(intersect(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) []int {

	sort.Ints(nums1)
	sort.Ints(nums2)

	i := 0
	j := 0

	var res []int

	if len(nums1) > len(nums2) {
		for j < len(nums2) && i < len(nums1) {
			if nums1[i] == nums2[j] {
				res = append(res, nums1[i])
				i++
				j++
			} else if nums1[i] > nums2[j] {
				j++
			} else {
				i++
			}
		}

	} else {
		for i < len(nums1) && j < len(nums2) {
			if nums1[i] == nums2[j] {
				res = append(res, nums1[i])
				i++
				j++
			} else if nums1[i] > nums2[j] {
				j++
			} else {
				i++
			}
		}
	}

	return res
}

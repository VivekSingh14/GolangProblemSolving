package main

import (
	"fmt"
	"sort"
)

func main4() {

	nums1 := []int{9, 4, 9, 8, 4}
	nums2 := []int{4, 9, 5}
	sort.Ints(nums1)
	sort.Ints(nums2)

	fmt.Println(intersect(nums1, nums2))

}

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	i := 0
	j := 0

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

	return res
}

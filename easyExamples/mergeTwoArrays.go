package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	fmt.Println(mergeTwoArrays(nums1, 3, nums2, 3))
}

func mergeTwoArrays(nums1 []int, m int, nums2 []int, n int) []int {
	mergedNums := make([]int, m+n)

	i := 0
	j := 0
	k := 0

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			mergedNums[k] = nums1[i]
			i++
		} else {
			mergedNums[k] = nums2[j]
			j++
		}

		k++
	}

	for i < m {
		mergedNums[k] = nums1[i]
		i++
		k++
	}

	for j < n {
		mergedNums[k] = nums2[j]
		j++
		k++
	}

	copy(nums1, mergedNums)
	return mergedNums
}

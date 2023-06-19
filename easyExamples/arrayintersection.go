package main

import "fmt"

// /- find out intersection of two arrays -/ //

// constraint = if 1 <= nums1.length, nums2.length <= 10

func main() {
	arr1 := []int{4, 9, 5}
	arr2 := []int{9, 4, 9, 8, 4}
	fmt.Println(findintersectionsBruteF(arr1, arr2))

}

func findintersectionsBruteF(nums1 []int, nums2 []int) []int {

	res := [10]int{}
	temp := []int{}
	if len(nums1) > len(nums2) {
		for i := 0; i < len(nums1); i++ {
			for j := 0; j < len(nums2); j++ {
				if nums1[i] == nums2[j] && res[nums2[j]] == 0 {
					res[nums1[i]] = 1
					temp = append(temp, nums1[i])
				}
			}
		}
	} else {
		for i := 0; i < len(nums2); i++ {
			for j := 0; j < len(nums1); j++ {
				if nums2[i] == nums1[j] && res[nums1[j]] == 0 {
					res[nums2[i]] = 1
					temp = append(temp, nums2[i])
				}
			}
		}
	}
	return temp
}

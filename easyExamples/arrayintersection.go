package main

import (
	"fmt"
	"sort"
)

// /- find out intersection of two arrays -/ //

// constraint = if 1 <= nums1.length, nums2.length <= 10

func main5() {
	arr2 := []int{1, 2, 2, 1}
	arr1 := []int{2, 2}
	fmt.Println(findintersectionsBruteF(arr1, arr2))
	fmt.Println(findintersectionsHashTable(arr1, arr2))
	fmt.Println(findintersectionsTwoPointer(arr1, arr2))

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

func findintersectionsHashTable(nums1 []int, nums2 []int) []int {

	hashmap := make(map[int]int)
	res := [1000]int{}
	var temp []int
	if len(nums1) > len(nums2) {
		for i := 0; i < len(nums1); i++ {
			if _, ok := hashmap[nums1[i]]; !ok {
				hashmap[nums1[i]] = 1
			}
		}

		for i := 0; i < len(nums2); i++ {
			if _, ok := hashmap[nums2[i]]; ok {
				res[nums2[i]] = 1
			}
		}
	} else {
		for i := 0; i < len(nums2); i++ {
			if _, ok := hashmap[nums2[i]]; !ok {
				hashmap[nums2[i]] = 1
			}
		}

		for i := 0; i < len(nums1); i++ {
			if _, ok := hashmap[nums1[i]]; ok {
				res[nums1[i]] = 1
			}
		}
	}

	for i := 0; i < len(res); i++ {
		if res[i] == 1 {
			temp = append(temp, i)
		}
	}

	return temp
}

func findintersectionsTwoPointer(nums1 []int, nums2 []int) []int {
	var count []int
	var res []int
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})

	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})

	i := 0
	j := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			count = append(count, i)
			i = i + 1
			j = j + 1
		} else if nums1[i] < nums2[j] {
			i = i + 1
		} else {
			j = j + 1
		}
	}

	for i := 0; i < len(nums1); i++ {

	}
	return res
}

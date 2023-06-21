package main

import "fmt"

func main() {
	nums1 := []int{1, 1, 3, 2}
	nums2 := []int{2, 3}
	nums3 := []int{3}

	fmt.Println(twoOutOfThree(nums1, nums2, nums3))
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	one, two, three := make([]int, 101), make([]int, 101), make([]int, 101)

	for _, num := range nums1 {
		one[num] = one[num] + 1
	}

	for _, num := range nums2 {
		two[num] = two[num] + 1
	}

	for _, num := range nums3 {
		three[num] = three[num] + 1
	}

	var res []int

	for i := 0; i < 101; i++ {
		if one[i] > 0 && two[i] > 0 {
			res = append(res, i)
		} else if two[i] > 0 && three[i] > 0 {
			res = append(res, i)
		} else if three[i] > 0 && one[i] > 0 {
			res = append(res, i)
		}
	}

	return res
}

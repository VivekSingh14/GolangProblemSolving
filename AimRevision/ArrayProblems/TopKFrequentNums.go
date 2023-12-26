package main

import "fmt"

func main() {
	//arr := []int{3, 1, 4, 4, 5, 2, 6, 1}
	arr := []int{7, 10, 11, 5, 2, 5, 5, 7, 11, 8, 9}
	//k := 2
	k := 4

	fmt.Println(topKFrequent(arr, k))
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, val := range nums {
		m[val]++
	}

	bucket := make([][]int, len(nums)+1)

	for key, val := range m {
		bucket[val] = append(bucket[val], key)
	}

	ans := make([]int, 0, k)

	for i := len(bucket) - 1; i >= 0; i-- {
		for _, val := range bucket[i] {
			if k > 0 {
				ans = append(ans, val)
				k--
			}
		}
	}

	return ans
}

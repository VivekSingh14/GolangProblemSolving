package main

import (
	"fmt"
	"math"
	"sort"
)

func main21() {
	arr := []int{7, 10, 4, 3, 20, 15}
	K := 3

	//KthSmallest(arr, K)
	fmt.Println(findKthLargest(arr, K))
}

func KthSmallest(arr []int, k int) {
	sort.Ints(arr)

	fmt.Println(arr[k-1])
}

func findKthLargest(nums []int, k int) int {

	maxValue, minValue := math.MinInt, math.MaxInt
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
		}

		if nums[i] < minValue {
			minValue = nums[i]
		}
	}

	count := make([]int, maxValue-minValue+1)
	for i := 0; i < len(nums); i++ {
		count[nums[i]-minValue]++
	}

	var result int
	for i := 0; i < len(count) && k > 0; i++ {
		if count[i] > 0 {
			result = minValue + i
			k -= count[i]
		}
	}
	return result
}

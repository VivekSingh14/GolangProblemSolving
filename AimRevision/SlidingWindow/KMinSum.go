package main

import (
	"fmt"
	"math"
)

func main1() {
	arr := []int{10, 4, 2, 5, 6, 3, 8, 1}
	k := 3
	minSum(arr, k)
}

func minSum(arr []int, k int) {
	min_window := math.MaxInt

	window_sum := 0

	for i := 0; i < k; i++ {
		window_sum = window_sum + arr[i]
	}
	min_window = window_sum
	for i := k; i < len(arr); i++ {
		//fmt.Println(k - i)
		window_sum = window_sum + arr[i] - arr[i-k]
		if window_sum < min_window {
			min_window = window_sum
		}
	}
	fmt.Println(min_window)
}

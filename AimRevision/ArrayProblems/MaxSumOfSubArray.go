package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	maxSumSubarray(arr)
}

func maxSumSubarray(arr []int) {
	temp_sum := 0
	max_Sum := math.MinInt
	for i := 0; i < len(arr); i++ {
		temp_sum += arr[i]

		if max_Sum < temp_sum {
			max_Sum = temp_sum
		}
		if temp_sum < 0 {
			temp_sum = 0
		}
	}
	fmt.Println(max_Sum)
}

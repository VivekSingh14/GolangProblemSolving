package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{1, 4, 2, 10, 23, 3, 1, 0, 20}
	k := 4

	maxKsum(arr, k)
}

func maxKsum(arr []int, k int) {
	maxSum := math.MinInt
	tempSum := 0
	for i := 0; i < k; i++ {
		tempSum += arr[i]
	}
	if tempSum > maxSum {
		maxSum = tempSum
	}

	for i := k; i < len(arr); i++ {
		if arr[i]+tempSum-arr[i-k] > maxSum {
			tempSum = arr[i] + tempSum - arr[i-k]
			maxSum = tempSum
		}
	}
	fmt.Println(maxSum)
}

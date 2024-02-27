package main

import "fmt"

func main() {
	arr := []int{1, 4, 2, 10, 23, 3, 1, 0, 20}
	k := 4
	fmt.Println(maxSumSubArray(arr, k))
}

func maxSumSubArray(arr []int, k int) int {

	res := 0

	for i := 0; i < k; i++ {
		res += arr[i]
	}

	for i := k; i < len(arr); i++ {
		temp_sum := res + arr[i] - arr[i-k]

		if temp_sum > res {
			res = temp_sum
		}
	}
	return res
}

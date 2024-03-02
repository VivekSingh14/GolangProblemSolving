package main

import "fmt"

func main8() {
	arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}

	fmt.Println(maxSum(arr))

	fmt.Println(maxSumOptimized(arr))
}

func maxSum(arr []int) int {
	res := 0

	for i := 0; i < len(arr)-1; i++ {
		temp_res := 0
		for j := i + 1; j < len(arr); j++ {
			temp_res += arr[j]
			if temp_res > res {
				res = temp_res
			}
		}

	}
	return res
}

func maxSumOptimized(arr []int) int {
	res := 0

	temp_res := 0

	for i := 0; i < len(arr)-1; i++ {
		temp_res += arr[i]
		if temp_res < 0 {
			temp_res = 0
		}
		if temp_res > res {
			res = temp_res
		}

	}
	return res
}

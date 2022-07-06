package main

import "fmt"

func main() {
	arr := make([]int, 5)
	arr[0] = 2
	arr[1] = 3
	arr[2] = 4
	arr[3] = 5
	arr[4] = 6

	k := 2

	fmt.Printf("%d \t\n", sum(arr, k))
	fmt.Println("**************")
	fmt.Printf("%d \t", sumSlidingWindow(arr, k))
}
func sum(arr []int, k int) int {
	max_sum := -100

	for i := 0; i < len(arr)-k+1; i++ {
		curr_sum := 0
		for j := 0; j < k; j++ {
			curr_sum = curr_sum + arr[i+j]
		}
		if curr_sum > max_sum {
			max_sum = curr_sum
		}
	}
	return max_sum
}

func sumSlidingWindow(arr []int, k int) int {
	window_sum := 0
	for i := 0; i < k; i++ {
		window_sum = window_sum + arr[i]
	}
	max_sum := window_sum
	for i := k; i < len(arr)-k+1; i++ {
		if arr[i]+arr[i+1] > max_sum {
			window_sum = arr[i] + arr[i+1]
			max_sum = window_sum
		}
	}
	return max_sum
}

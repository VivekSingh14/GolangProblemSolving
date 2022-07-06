package main

import "fmt"

func main() {
	arr := []int{15, 2, 4, 8, 9, 5, 10, 23}
	sum := 23
	subarraySum(arr, sum)
}

func subarraySum(arr []int, sum int) {
	n := len(arr)
	start := 0

	curr_sum := arr[0]

	for i := 1; i < n; i++ {

		if curr_sum > sum && start < i-1 {
			curr_sum = curr_sum - arr[start]
			start++
		}
		if curr_sum == sum {
			fmt.Printf("From index %d to index %d", start, i-1)
			return
		}
		curr_sum = curr_sum + arr[i]
	}
}

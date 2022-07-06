package main

import "fmt"

//TODO: SEQ NO: 6
//TODO:  Find subarray with given sum from array of integers
//TODO: with N complexity.

func main() {
	arr := []int{1, 4, 20, 3, 10, 5}
	sum := 33
	curr_sum := arr[0]
	start := 0

	for i := 1; i < len(arr); i++ {
		for curr_sum > sum && start < i-1 {
			curr_sum = curr_sum - arr[start]
			start++
		}
		if curr_sum == sum {
			fmt.Println(start, " ", i-1)
		}

		if i < len(arr) {
			curr_sum = curr_sum + arr[i]
		}
	}
}

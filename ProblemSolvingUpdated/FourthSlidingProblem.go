package main

import "fmt"

//TODO SEQ NO 4
//TODO Find Index of 0 to be replaced with 1 to get longest continuous sequence of 1s in a binary array

func main18() {
	arr := []int{1, 1, 0, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1}
	length := len(arr)

	fmt.Println(maxOnesindex2(arr, length))
}

func maxOnesindex2(arr []int, length int) int {
	max_count := 0
	max_index := 0
	prev_zero := -1
	prev_prev_zero := -1

	for curr := 0; curr < length; curr++ {

		if arr[curr] == 0 {
			if curr-prev_prev_zero > max_count {
				max_count = curr - prev_prev_zero
				max_index = prev_zero
			}
			prev_prev_zero = prev_zero
			prev_zero = curr
		}
	}

	if length-prev_prev_zero > max_count {
		max_index = prev_zero
	}
	return max_index
}

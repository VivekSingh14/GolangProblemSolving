package main

//SEQ NO 6
// Find minimum sum subarray of given size k

import "fmt"

func main() {
	arr := []int{10, 4, 2, 5, 6, 3, 8, 1}
	k := 3
	min_sum := 0
	for i := 0; i < k; i++ {
		min_sum += arr[i]
	}
	tempSum := min_sum
	for j := k; j < len(arr); j++ {
		tempSum = tempSum + arr[j] - arr[j-k]
		if tempSum < min_sum {
			min_sum = tempSum
		}
	}
	fmt.Println(min_sum)
}

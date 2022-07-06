package main

import "fmt"

func main() {
	arr := []int{1, 4, 2, 10, 23, 3, 1, 0, 20}
	k := 4
	max_sum := 0
	for i := 0; i < k; i++ {
		max_sum += arr[i]
	}

	for j := k; j < len(arr); j++ {
		tempSum := max_sum + arr[j] - arr[j-k]
		if tempSum > max_sum {
			max_sum = tempSum
		}
	}
	fmt.Println(max_sum)
}

package main

import "fmt"

func main15() {
	nums := []int{-4, -1, 0, 3, 10}
	sortedSquares(nums)
}

func sortedSquares(arr []int) {
	squares := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		if arr[i]*arr[i] > squares[i] {
			squares[i] = arr[i] * arr[i]
		}
	}
	fmt.Println(squares)
}

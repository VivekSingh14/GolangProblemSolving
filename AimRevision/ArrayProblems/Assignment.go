package main

import (
	"fmt"
	"sort"
)

func main5() {
	A := []int{-1, -3}
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	sort.Ints(A)
	fmt.Println(A)

	for i := 0; i < len(A)-1; i++ {
		if A[i] != A[i+1] && A[i+1]-A[i] > 1 && ((A[i+1] > 0) && (A[i] > 0)) {
			return A[i] + 1
		} else if (A[i+1]-A[i] > 1) && ((A[i+1] < 0) && (A[i] < 0)) {
			return 1
		}
	}
	return A[len(A)-1] + 1
}

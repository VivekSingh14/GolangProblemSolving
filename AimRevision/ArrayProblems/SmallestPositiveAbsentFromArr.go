package main

import "fmt"

func main34() {

	arr := []int{-1, -2 - 6, -4, 0, 1, 5, 4, 2}

	fmt.Println(Solution2(arr))
}

func Solution2(A []int) int {
	// Implement your solution here
	var ptr int

	for i := 0; i < len(A); i++ {
		if A[i] == 1 {
			ptr = 1
			break
		}
	}
	if ptr == 0 {
		return 1
	}

	for i := 0; i < len(A); i++ {
		if A[i] <= 0 || A[i] > len(A) {
			A[i] = 1
		}
	}

	for i := 0; i < len(A); i++ {
		A[(A[i]-1)%len(A)] += len(A)
	}

	for i := 0; i < len(A); i++ {
		if A[i] <= len(A) {
			return i + 1
		}
	}
	return len(A) + 1

}

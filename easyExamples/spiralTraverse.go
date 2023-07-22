package main

import "fmt"

func main32() {
	arr := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	res := traverse(arr)
	fmt.Println(res)

}

func traverse(matrix [][]int) []int {

	top := 0
	left := 0
	right := len(matrix[0]) - 1
	down := len(matrix) - 1

	elements := make([]int, 0)

	for left <= right && top <= down {
		// matrix[top][left -> right]
		for i := left; i <= right; i++ {
			elements = append(elements, matrix[top][i])
		}
		top++

		// matrix[top -> down][right]
		for i := top; i <= down; i++ {
			elements = append(elements, matrix[i][right])
		}
		right--

		// matrix[down][right -> left]
		for i := right; left <= i; i-- {
			elements = append(elements, matrix[down][i])
		}
		down--

		// matrix[down -> top][left]
		for i := down; top <= i; i-- {
			elements = append(elements, matrix[i][left])
		}
		left++
	}

	area := len(matrix) * len(matrix[0])
	return elements[:area]
}

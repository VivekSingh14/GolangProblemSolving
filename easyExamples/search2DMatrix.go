package main

import "fmt"

func main18() {

	//matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}

	//matrix := [][]int{{1}, {3}}

	//matrix := [][]int{{1}}

	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	target := 23

	fmt.Println(searchMatrix(matrix, target))

}

func searchMatrix(matrix [][]int, target int) bool {

	length := len(matrix)

	if length == 1 {
		i := 0

		min := 0
		max := len(matrix[i]) - 1

		for max >= min {
			mid := min + (max-min)/2
			if target == matrix[i][mid] {
				return true
			} else if matrix[i][mid] > target {
				max = mid - 1
			} else {
				min = mid + 1
			}
		}
	}

	for i := 0; i < length; i++ {
		if target == matrix[i][0] || target == matrix[i][len(matrix[i])-1] {
			return true
		}

		if target > matrix[i][0] && target < matrix[i][len(matrix[i])-1] {
			min := 0
			max := len(matrix[i]) - 1

			for max >= min {
				mid := min + (max-min)/2
				if target == matrix[i][mid] {
					return true
				} else if matrix[i][mid] > target {
					max = mid - 1
				} else {
					min = mid + 1
				}
			}
		}
	}

	return false
}

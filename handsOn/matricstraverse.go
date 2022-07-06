package main

import "fmt"

func main() {
	var mat [3][3]int
	var count int = 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			mat[i][j] = count
			count = count + 2
		}
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if i == j {
				fmt.Print(mat[i][j], " ")
			}
		}
		fmt.Println()
	}

}

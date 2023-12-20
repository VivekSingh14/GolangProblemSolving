package main

import "fmt"

func main16() {
	arr := [][]int{
		{1, 1, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{1, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 0, 0}}

	ShowNoOfIslands(arr)
}

func ShowNoOfIslands(arr [][]int) {
	//islands := 0

	fmt.Println(len(arr[0]))

	for i := 0; i < len(arr); i++ {

	}
}

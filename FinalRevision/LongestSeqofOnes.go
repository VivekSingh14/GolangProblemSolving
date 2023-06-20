package main

import "fmt"

func main3() {
	arr := []int{0, 0, 1, 0, 1, 1, 1, 0, 1, 1}
	var templen int
	var max int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			templen = j - i + 1
			temp := checkMaxlength(arr, i, j, templen)
			if temp > max {
				max = temp
			}
		}
	}
	fmt.Println(max)
}

func checkMaxlength(arr []int, i int, j int, templen int) int {
	var numofones int
	for k := i; k <= j; k++ {
		if arr[k] == 1 {
			numofones++
		}
	}
	if templen-numofones == 1 {
		return templen
	}

	return 0
}

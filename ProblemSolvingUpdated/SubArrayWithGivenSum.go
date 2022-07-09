package main

import "fmt"

//TODO: SEQ NO: 7
//TODO:  Find subarray with given sum from array of integers
//TODO: with N^2 complexity.

func main17() {
	arr := []int{1, 4, 20, 3, 10, 5}
	sum := 33
	temp_sum := 0
	for i := 0; i < len(arr)-1; i++ {
		temp_sum = arr[i]
		for j := i + 1; j < len(arr); j++ {
			if temp_sum == sum {
				fmt.Println(i, " ", j-1)
				break
			}
			if temp_sum > sum || j == len(arr)-1 {
				break
			}
			temp_sum = temp_sum + arr[j]
		}
	}

}

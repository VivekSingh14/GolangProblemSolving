package main

import "fmt"

func main1() {
	var arr []int = []int{2, 4, 6, 8, 10, 12}
	fmt.Println("Sum of an array elements is: ", sum(arr))
}

func sum(arr []int) int {
	var sum int = 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}

	return sum
}

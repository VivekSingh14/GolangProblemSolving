package main

import "fmt"

func main() {
	//sequentially
	arr := []int{0, -1, 2, -3, 1}

	tripletSumZero(arr)
}

func tripletSumZero(arr []int) {
	sum := arr[0] + arr[1] + arr[2]

	if sum == 0 {
		fmt.Println(arr[0], " ", arr[1], " ", arr[2])
	}
	for i := 3; i < len(arr); i++ {
		sum = sum - arr[i-3] + arr[i]
		if sum == 0 {
			fmt.Println(arr[i-2], " ", arr[i-1], " ", arr[i])
		}
	}
}

package main

import "fmt"

func main() {
	var arr []int = []int{6, 1, 8, 3, 10, 5, 33, 19}

	var max int = 0
	for i := 0; i < 8; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Printf("Maximum number is: %d \n", max)

}

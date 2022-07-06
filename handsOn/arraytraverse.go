package main

import "fmt"

func main() {
	var arrsize int = 5

	arr := make([]int, arrsize)

	fmt.Println("Please insert the array elements:  ")
	for i := 0; i < arrsize; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < arrsize; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

}

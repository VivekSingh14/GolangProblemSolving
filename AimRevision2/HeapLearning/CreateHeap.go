package main

import "fmt"

func main() {
	arr := []int{0, 10, 20, 30, 25, 5, 40, 35}
	CreateMaxHeap(arr)
}

func CreateMaxHeap(arr []int) {

	for i := 2; i < len(arr); i++ {

		fmt.Println(InsertionIntoMaxHeap2(arr, i))
	}
}

func InsertionIntoMaxHeap2(arr []int, n int) []int {

	i := n
	temp := arr[n]

	for i > 1 && temp > arr[i/2] {
		arr[i] = arr[i/2]
		i = i / 2
	}
	arr[i] = temp
	return arr
}

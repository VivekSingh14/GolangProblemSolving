package main

import "fmt"

func main() {
	heap := []int{0, 30, 20, 15, 5, 10, 12, 6}
	element := 40
	//res := InsertionIntoMaxHeap(heap, element)
	fmt.Println(InsertionIntoMaxHeap(heap, element)[1:])

}

func InsertionIntoMaxHeap(arr []int, element int) []int {

	arr = append(arr, element)
	i := len(arr) - 1
	for i > 1 {
		if arr[i/2] < element {
			arr[i] = arr[i/2]
			arr[i/2] = element
		}
		i = i / 2
	}

	return arr
}

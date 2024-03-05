package main

import "fmt"

func main() {
	heap := []int{0, 5, 8, 10, 11, 12, 14, 16}
	element := 6
	//res := InsertionIntoMaxHeap(heap, element)
	fmt.Println(InsertionIntoMinHeap(heap, element)[1:])
}

func InsertionIntoMinHeap(heap []int, element int) []int {
	heap = append(heap, element)
	i := len(heap) - 1
	for i > 1 && heap[i/2] > element {
		temp := heap[i/2]
		heap[i/2] = heap[i]
		heap[i] = temp
		i = i / 2
	}
	return heap
}

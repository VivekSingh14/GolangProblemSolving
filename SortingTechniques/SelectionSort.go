package main

import "fmt"

func main() {
	arr := []int{3, 7, 9, 10, 6, 5, 12, 4, 11, 2}
	sorted := SelectionSort(arr)
	fmt.Println(sorted)
}

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		k := i
		for j := i; j < len(arr); j++ {
			if arr[k] > arr[j] {
				k = j
			}
		}
		temp := arr[k]
		arr[k] = arr[i]
		arr[i] = temp
	}
	return arr
}

package main

import "fmt"

func main() {
	arr := []int{3, 7, 9, 10, 6, 5, 12, 4, 11, 2}
	sorted := InsertionSort(arr)
	fmt.Println(sorted)
}

func InsertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	return arr
}

package main

import "fmt"

func main() {

	arr := []int{3, 7, 9, 10, 6, 5, 12, 4, 11, 2}

	fmt.Println(arr)

	quickSort(arr, 0, len(arr)-1)

	fmt.Println(arr)

}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low int, high int) int {

	pivot := arr[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
	}
	swap(arr, i+1, high)
	return i + 1

}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

package main

import "fmt"

func main3() {
	arr := []int{2, 4, 6, 8, 10}
	num := 4
	place := binarySearch(arr, 0, len(arr)-1, num)
	fmt.Println("Number is found at", place)
}

func binarySearch(arr []int, first int, length int, num int) int {

	midTemp := first + (length-first)/2
	if arr[midTemp] == num {
		return midTemp
	}
	if num < arr[(first+(length-first))/2] {
		mid := (first + (length - first)) / 2
		return binarySearch(arr, first, mid-1, num)
	} else {
		mid := (first + (length - first)) / 2
		return binarySearch(arr, mid+1, length-1, num)
	}
}

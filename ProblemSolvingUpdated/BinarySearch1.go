package main

import "fmt"

func main() {
	arr := []int{-1, 0, 3, 5, 9, 12}
	target := 12
	start := 0
	end := len(arr)
	binarySearch1(arr, start, end, target)
}

func binarySearch1(arr []int, start int, end int, target int) {
	if target == arr[(start+end)/2] {
		fmt.Println("found at: ", (start+end)/2+1)
		return
	}

	if target > arr[(start+end)/2] {
		start = ((start + end) / 2) + 1
		binarySearch1(arr, start, end, target)
	} else {
		end = ((start + end) / 2) - 1
		binarySearch1(arr, start, end, target)
	}

}

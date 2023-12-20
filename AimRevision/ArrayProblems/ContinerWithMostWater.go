package main

import "fmt"

func main15() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	containerWithMostWater(arr)

}

func containerWithMostWater(arr []int) {
	l := 0

	r := len(arr) - 1
	max1 := 0
	for l < r {
		length := min1(arr[l], arr[r])
		if (r-l)*length > max1 && length == arr[l] {
			max1 = (r - l) * length
			l++
		} else if (r-l)*length > max1 && length == arr[r] {
			max1 = (r - l) * length
			r--
		} else {
			if length == arr[l] {
				l++
			} else {
				r--
			}
		}
	}
	fmt.Println(max1)
}

func min1(a, b int) int {
	if a > b {
		return b
	}
	return a
}

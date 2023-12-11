package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{0, -1, 2, -3, 1}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	tripletSumZeroUpdated(arr)
}

func tripletSumZeroUpdated(arr []int) {

	for j := 0; j < len(arr); j++ {
		l := j + 1
		r := len(arr) - 1
		x := arr[j]
		for l < r {
			if x+arr[l]+arr[r] == 0 {
				fmt.Println(x, " ", arr[l], " ", arr[r])

				l++
				r--
			} else if x+arr[l]+arr[r] < 0 {

				l++
			} else {
				r--
			}
		}
	}
	fmt.Println("")
}

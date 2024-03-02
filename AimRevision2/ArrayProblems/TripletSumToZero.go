package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{0, -1, 2, -3, 1}

	fmt.Println(tripletSumZero(arr))
	tripletSumZeroOptimized(arr)
}

func tripletSumZero(arr []int) [][]int {

	var res [][]int

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			for k := j; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 0 && i != j && i != k && j != k {
					var temp []int
					temp = append(temp, arr[i])
					temp = append(temp, arr[j])
					temp = append(temp, arr[k])
					res = append(res, temp)
				}
			}
		}
	}
	return res
}

func tripletSumZeroOptimized(arr []int) {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

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

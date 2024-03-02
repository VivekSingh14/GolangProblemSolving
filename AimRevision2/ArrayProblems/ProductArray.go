package main

import "fmt"

func main9() {

	arr := []int{1, 2, 3, 4}
	fmt.Println(productArray(arr))
	fmt.Println(productArrayOptimized(arr))
}

func productArray(arr []int) []int {

	var res []int

	for i := 0; i < len(arr); i++ {
		temp := 1
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}
			temp *= arr[j]
		}
		res = append(res, temp)
	}

	return res
}

func productArrayOptimized(arr []int) []int {

	res := make([]int, len(arr))

	res[0], res[len(arr)-1] = 1, 1

	for i := 1; i < len(arr); i++ {
		res[i] = res[i-1] * arr[i-1]
	}

	tempProduct := 1

	for i := len(arr) - 2; i >= 0; i-- {
		tempProduct *= arr[i+1]
		res[i] *= tempProduct
	}

	return res
}

package main

import "fmt"

func main() {
	arr := []int{7, 6, 4, 3, 1}
	maxProfit(arr)
}

func maxProfit(arr []int) {
	minPrice := arr[0]
	maxProf := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < minPrice {
			minPrice = arr[i]
		} else if arr[i]-minPrice > maxProf {
			maxProf = arr[i] - minPrice
		}
	}
	fmt.Println(maxProf)
}

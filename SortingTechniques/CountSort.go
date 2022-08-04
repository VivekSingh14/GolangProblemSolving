package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{3, 7, 9, 10, 6, 5, 12, 4, 11, 2}

	max := findMax(arr)

	hashTable := make([]int, max+1)

	for i := 0; i < len(arr); i++ {
		hashTable[arr[i]]++
	}

	fmt.Printf("%d \n", hashTable)
	k := 0
	for i := 0; i < len(hashTable); i++ {
		j := hashTable[i]
		for j > 0 {
			arr[k] = i
			k++
			j--
		}
	}
	fmt.Printf("%d \n", arr)
}

func findMax(arr []int) int {
	max := math.MinInt

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

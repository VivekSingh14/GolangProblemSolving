//given an array A[] of n numbers and another number k, determines whether or not there exist two elements in A[] whose sum is exactly k.

// two pointer approach
package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{35, 20, 10, 80, 75, 50}

	k := 70

	sort.Ints(arr)
	//10, 20, 35, 50, 75, 80

	fmt.Println(checkPair(arr, k))

}

func checkPair(arr []int, k int) bool {
	start := 0
	end := len(arr) - 1

	for start < end {
		if arr[start]+arr[end] > k {
			end--
		} else if arr[start]+arr[end] < k {
			start++
		} else {
			return true
		}

	}

	return false
}

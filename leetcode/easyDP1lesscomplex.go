package main

//largest sum subarray using kadane's algo with O(n) complexity

import (
	"fmt"
	"math"
)

func main2() {

	var arr []int = []int{5, 4, -1, 7, 8}
	//
	//-2, 1, -3, 4, -1, 2, 1, -5, 4

	max_so_far := math.MinInt
	max_end_here := 0

	for i := 0; i < len(arr); i++ {
		max_end_here = max_end_here + arr[i]
		if max_so_far < max_end_here {
			max_so_far = max_end_here
		}
		if max_end_here < 0 {
			max_end_here = 0
		}
	}

	fmt.Println(max_so_far)

}

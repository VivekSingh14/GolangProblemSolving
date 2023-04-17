//n^2 complexity
package main

import (
	"fmt"
	"math"
)

func main1() {
	var arr []int = []int{5, 4, -1, 7, 8}
	//
	//-2, 1, -3, 4, -1, 2, 1, -5, 4

	maxsum := math.MinInt

	for i := 0; i < len(arr); i++ {
		tempsum := 0
		for j := i; j < len(arr); j++ {
			tempsum = tempsum + arr[j]
			if maxsum < tempsum {
				maxsum = tempsum
			}
		}

	}

	fmt.Println(maxsum)

}

package main

import (
	"fmt"
	"math"
)

func main() {
	strs := []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {

	res := ""

	minLength := minLengthOfPattern(strs)

	low := 0
	high := minLength - 1

	for low <= high {
		mid := low + (high-low)/2
		if allContains(strs, strs[0], low, mid) {
			re := strs[0]
			res += re[low : mid+1]
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return res
}

func minLengthOfPattern(strs []string) int {
	minLength := math.MaxInt
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < minLength {
			minLength = len(strs[i])

		}
	}
	return minLength
}

func allContains(strs []string, com string, min int, mid int) bool {

	for i := 0; i < len(strs); i++ {
		tempstr := strs[i]
		for j := min; j <= mid; j++ {
			if tempstr[j] != com[j] {
				return false
			}
		}
	}
	return true
}

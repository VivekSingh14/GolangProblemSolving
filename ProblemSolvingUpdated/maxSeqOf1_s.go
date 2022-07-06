package main

import "fmt"

//TODO: SEQ NO: 5
//TODO:  Find the maximum sequence of continuous 1’s that can be formed by replacing at-most k zeroes by ones

func main11() {
	arr := []bool{true, true, true, true, false}
	fmt.Println(maxOnesindex(arr))
}

func maxOnesindex(arr []bool) int {
	n := len(arr)
	start := 0
	end := 0
	maxIndex := -1
	lastInd := -1
	maxCnt := 0

	for end < n {
		for end < n && arr[end] {
			//fmt.Println("Inner loop ", end)
			end++
		}
		if maxCnt < end-start && lastInd != -1 {
			maxCnt = end - start
			maxIndex = lastInd
		}
		start = lastInd + 1
		lastInd = end
		//fmt.Println("Outer loop ", end)
		end++
	}
	if maxCnt < end-start && lastInd != -1 {
		maxCnt = end - start
		maxIndex = lastInd
	}
	return maxIndex
}

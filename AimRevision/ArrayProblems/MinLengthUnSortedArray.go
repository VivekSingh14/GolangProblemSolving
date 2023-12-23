package main

import "fmt"

func main18() {
	arr := []int{10, 12, 20, 30, 25, 40, 32, 31, 35, 50, 60}
	minLengthUnsortedArray(arr)
}

func minLengthUnsortedArray(nums []int) {
	endIndex := -1
	maximum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < maximum {
			endIndex = i
		} else {
			maximum = nums[i]
		}
	}

	startIndex := len(nums)
	minimum := nums[len(nums)-1]
	for i := len(nums) - 2; i > -1; i-- {
		if nums[i] > minimum {
			startIndex = i
		} else {
			minimum = nums[i]
		}
	}
	if endIndex-startIndex > 0 {
		fmt.Println(endIndex, " ", startIndex+1)
		return
	}

	return
}

package main

import "fmt"

func main() {
	nums := []int{2, 3, -1, 8, 4}
	// o/p = 3
	// because leftmax = 2 + 3 + (-1)
	// 8 is on 3 index and result is 3
	// rightmax = 4
	fmt.Println(findMiddleIndex(nums))

}

func findMiddleIndex(nums []int) int {
	leftMax := 0
	rightMax := 0

	for _, num := range nums {
		leftMax += num
	}

	for i := 0; i < len(nums); i++ {
		rightMax += nums[i]
		if leftMax == rightMax {
			return i
		}
		leftMax -= nums[i]
	}
	return -1
}

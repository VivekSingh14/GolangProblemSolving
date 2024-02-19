package main

import "fmt"

func main25() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(robT(nums))

}

func robT(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	r1 := robber(nums[1:])
	r2 := robber(nums[:len(nums)-1])
	if r1 > r2 {
		return r1
	}
	return r2
}

func robber(nums []int) int {
	rob1, rob2 := 0, 0
	for _, n := range nums {
		temp := maxsss(n+rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}
	return rob2
}

func maxsss(a, b int) int {
	if a > b {
		return a
	}

	return b
}

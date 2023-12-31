package main

import "fmt"

func main28() {
	nums := []int{2, 3, 2}
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
	r1, r2, temp := 0, 0, 0
	for _, n := range nums {
		if r1+n > r2 {
			temp = r1 + n
		} else {
			temp = r2
		}
		r1 = r2
		r2 = temp

	}
	return r2
}

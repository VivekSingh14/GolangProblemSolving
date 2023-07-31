package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotate(nums, k)

}

func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		temp := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = temp
	}
	fmt.Println(nums)
}

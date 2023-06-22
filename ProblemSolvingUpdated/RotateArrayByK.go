package main

import "fmt"

func main29() {
	//nums = [1,2,3,4,5,6,7]

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotate(nums, k)

}

func rotate(nums []int, k int) {
	max := len(nums) - 1
	for i := 1; i <= k; i++ {
		temp := nums[len(nums)-1]
		for j := 0; j < len(nums)-1; j++ {
			nums[max-j] = nums[max-(j+1)]
		}
		nums[0] = temp
	}

	fmt.Println(nums)
}

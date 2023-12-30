package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 3, 2, 3}

	longestIncreasing(arr)
}

func longestIncreasing(nums []int) {

	var res []int
	res = append(res, nums[0])

	for i := 0; i < len(nums); i++ {
		if nums[i] > res[len(res)-1] {
			res = append(res, nums[i])
		} else {
			for j := 0; j < len(res); j++ {
				if res[j] >= nums[i] {
					res[j] = nums[i]
					break
				}
			}
		}
	}
	fmt.Println(len(res))
}

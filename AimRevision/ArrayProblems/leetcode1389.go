package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}

	fmt.Println(createTargetArray(nums, index))

}

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(nums))

	for i := 0; i < len(index); i++ {
		if res[index[i]] != -1 {
			for j := len(res) - 2; j >= index[i]; j-- {
				res[j+1] = res[j]
			}
			res[index[i]] = nums[i]
		}
		res[index[i]] = nums[i]
	}
	return res
}

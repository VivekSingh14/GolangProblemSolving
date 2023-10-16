package main

import "fmt"

func main4() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	//comment
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

func rotate1(nums []int, k int) {
	ln := len(nums)
	if ln == 1 {
		return
	}
	if ln < k {
		k = k % ln
	}
	sl1 := nums[ln-k:]
	sl2 := nums[:ln-k]
	sl1 = append(sl1, sl2...)
	for i := 0; i < ln; i++ {
		nums[i] = sl1[i]
	}
}

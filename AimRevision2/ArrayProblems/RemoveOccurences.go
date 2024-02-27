package main

import "fmt"

func main3() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElements(nums, val))
}

func removeElements(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}

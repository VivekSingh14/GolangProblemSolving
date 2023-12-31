package main

import "fmt"

func main8() {
	arr1 := []int{10, 3, 5, 6, 2}

	// o/p := {180, 600, 360, 300, 900}
	var res []int

	for i := 0; i < len(arr1); i++ {
		temp := 1
		for j := 0; j < len(arr1); j++ {
			if i != j {
				temp *= arr1[j]
			}
		}
		res = append(res, temp)
	}

	fmt.Println(res)

}

//other method

// func productExceptSelf(nums []int) []int {
//     size := len(nums)
//     res := make([]int, size)
//     res[0], res[size-1] = 1, 1

//     // Calculate the products of elements to the left of each index
//     for i := 1; i < size; i++ {
//         res[i] = res[i-1] * nums[i-1]
//     }

//     // Multiply by the products of elements to the right and accumulate in res
//     rightProduct := 1
//     for i := size - 2; i >= 0; i-- {
//         rightProduct *= nums[i+1]
//         res[i] *= rightProduct
//     }

//     return res
// }

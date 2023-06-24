package main

// /-- Two sum using hash map(optimised)  --/

import "fmt"

func main8() {
	arr := []int{2, 7, 11, 15}

	target := 9

	// hashmap := make(map[int]int)

	// for i := 0; i < len(arr); i++ {
	// 	if value, ok := hashmap[target-arr[i]]; ok {
	// 		fmt.Println(value, i)
	// 	}
	// 	hashmap[arr[i]] = i
	// }

	fmt.Println(twoSumTwoPointer(arr, target))

}

func twoSumTwoPointer(arr []int, target int) []int {
	var res []int
	i := 0
	j := len(arr) - 1

	for i < len(arr)-2 {
		if arr[i]+arr[j] == target {
			res = append(res, arr[i])
			res = append(res, arr[j])
			return res
		} else if arr[j] > target {
			j = j - 1
		} else {
			i = i + 1
		}
	}

	return res
}

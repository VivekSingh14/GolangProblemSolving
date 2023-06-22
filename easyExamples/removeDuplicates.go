package main

import "fmt"

// /- Remove Duplicates from Sorted Array -/ //

// two pointer algo because array of nums is already sorted.

func main9() {
	nums := []int{1, 1, 2}

	fmt.Println(removeDuplicates(nums))
	fmt.Println(removeDuplicatesHashTable(nums))
	fmt.Println(removeDuplicatesTwoPointer(nums))
}

func removeDuplicates(nums []int) int {

	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	//var res []int
	count := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			//res[count] = nums[i]
			count = count + 1
		}
	}
	//res[count] = nums[len(nums)-1]

	return count + 1
}

func removeDuplicatesHashTable(nums []int) int {

	hashmap := make(map[int]int)
	var res []int

	for i := 0; i < len(nums); i++ {
		if _, ok := hashmap[nums[i]]; !ok {
			hashmap[nums[i]] = 1
		} else {
			temp := hashmap[nums[i]] + 1
			hashmap[nums[i]] = temp
		}
	}

	for k, v := range hashmap {
		if v >= 1 {
			res = append(res, k)
		}
	}

	fmt.Println(res)
	return len(res)
}

func removeDuplicatesTwoPointer(nums []int) int {

	slow := 0
	fast := 1

	for fast < len(nums) {
		if nums[slow] == nums[fast] {
			fast = fast + 1
		} else {
			slow = slow + 1
			nums[slow] = nums[fast]
			fast = fast + 1
		}
	}
	return slow + 1
}

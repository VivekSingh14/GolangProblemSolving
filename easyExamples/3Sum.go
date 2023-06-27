package main

import (
	"fmt"
	"sort"
)

func main15() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	for first := 0; first < n-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -nums[first]
		for second := first + 1; second < n-1; second++ {
			// if second > first+1 && nums[second] == nums[second-1] {
			// 	continue
			// }
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func ThreeSumOptimized(nums []int) [][]int {
	o := [][]int{}
	// sort array
	sort.Ints(nums)
	for i := range nums {
		// skip duplicates
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		// two pointer, from next (j) and end (k)
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			// if sum > 0, decrease k
			if sum > 0 {
				k--
				// if sum < 0, increase j
			} else if sum < 0 {
				j++
				// if sum == 0
			} else {
				o = append(o, []int{nums[i], nums[j], nums[k]})
				j++
				// move j until new value found
				for nums[j-1] == nums[j] && j < k {
					j++
				}

			}
		}
	}
	return o
}

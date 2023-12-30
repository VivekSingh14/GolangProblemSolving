package main

import "fmt"

func main27() {
	nums := []int{1, 2, 3, 1}

	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		dp[i] = maxr(dp[i+1], dp[i+2]+nums[i])
	}
	return dp[0]
}

func maxr(a, b int) int {
	if a > b {
		return a
	}

	return b
}

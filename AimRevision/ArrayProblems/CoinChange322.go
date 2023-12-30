package main

import (
	"fmt"
)

func main24() {
	coins := []int{1, 2, 5}
	sum := 11
	fmt.Println(minCoins(coins, sum))

}

func minCoins(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = minCoin(dp[i], 1+dp[i-coin])
		}
	}

	if dp[amount] != amount+1 {
		return dp[amount]
	} else {
		return -1
	}
}
func minCoin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

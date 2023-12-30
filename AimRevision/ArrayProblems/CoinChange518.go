package main

import (
	"fmt"
)

func main25() {
	coins := []int{1, 2, 5}
	sum := 5
	fmt.Println(MinCoins(coins, sum))

}

func MinCoins(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

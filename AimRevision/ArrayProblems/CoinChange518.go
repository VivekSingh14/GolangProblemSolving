package main

import (
	"fmt"
)

func main23() {
	coins := []int{1, 2, 3}
	sum := 5
	fmt.Println(CoinsCombination(coins, sum))

}

func CoinsCombination(coins []int, amount int) int {
	dp := make([]int, amount+1)

	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

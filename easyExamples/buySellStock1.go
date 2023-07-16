package main

import (
	"fmt"
)

func main26() {
	prices := []int{3, 2, 6, 5, 0, 3}

	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	buy := prices[0]
	max_profit := 0

	for i := 1; i < len(prices); i++ {
		if buy > prices[i] {
			buy = prices[i]
		} else if prices[i]-buy > max_profit {
			max_profit = prices[i] - buy
		}
	}

	return max_profit
}

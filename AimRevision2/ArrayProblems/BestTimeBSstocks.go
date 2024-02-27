package main

import "fmt"

func main() {
	prices := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfit(prices))

}

func maxProfit(prices []int) int {
	min_price := prices[0]

	max_profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < min_price {
			min_price = prices[i]
		}
		if prices[i]-min_price > max_profit {
			max_profit = prices[i] - min_price
		}
	}

	return max_profit
}

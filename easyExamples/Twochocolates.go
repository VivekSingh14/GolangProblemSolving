package main

import (
	"fmt"
	"sort"
)

// -- buy two chocolates from given chocolate options and initial amount will be given,
// -- such that at the end of buying chocolates  you should have some amount or 0 debt.

func main() {
	price := []int{3, 2, 3}
	money := 3

	fmt.Println(buyChoco(price, money))

}

func buyChoco(prices []int, money int) int {

	sort.Slice(prices, func(i, j int) bool {
		return prices[i] < prices[j]
	})

	i := 0
	j := i + 1

	for i < len(prices) && j < len(prices) {
		if prices[i]+prices[j] == money || prices[i]+prices[j] <= money {
			res := money - (prices[i] + prices[j])
			return res
		}
		j = j + 1
		i = i + 1
	}

	return money
}

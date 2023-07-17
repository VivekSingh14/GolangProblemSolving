package main

import "fmt"

func main() {
	n := 5

	fmt.Println(climbStairs(n + 1))
}

func climbStairs(n int) int {
	m := 2
	ways := make([]int, n+1)

	ways[0] = 1
	ways[1] = 1

	for i := 2; i <= n; i++ {
		ways[i] = 0
		for j := 1; j <= m && j <= i; j++ {
			ways[i] += ways[i-j]
		}
	}
	return ways[n]

}

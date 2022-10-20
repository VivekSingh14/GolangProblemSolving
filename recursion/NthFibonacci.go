package main

import "fmt"

func main() {
	num := 5
	fmt.Println(NthFibonacci(num, 1))
}

func NthFibonacci(num int, fib int) int {

	if num > 0 {
		fib = fib + 1
		return NthFibonacci(num-1, fib)
	}
	return 1
}

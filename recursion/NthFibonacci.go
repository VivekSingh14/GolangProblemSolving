package main

import "fmt"

func main() {
	num := 85
	fmt.Println(NthFibonacci(num))
}

func NthFibonacci(num int) int {

	if num <= 1 {
		return num
	}
	return NthFibonacci(num-1) + NthFibonacci(num-2)
}

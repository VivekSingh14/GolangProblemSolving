package main

import "fmt"

func main() {
	num := 5
	fmt.Println(Factorial(num))
}

//adding a comment to continue the streak
func Factorial(num int) int {
	if num > 0 {
		return Factorial(num-1) * num
	}
	return 1
}

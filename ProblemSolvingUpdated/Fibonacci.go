package main

import "fmt"

func main7() {
	num1 := 0
	num2 := 1
	fmt.Print(num1, "\t")
	fmt.Print(num2, "\t")
	for i := 3; i <= 10; i++ {

		num3 := num1 + num2
		fmt.Print(num3, "\t")
		num1 = num2
		num2 = num3
	}
}

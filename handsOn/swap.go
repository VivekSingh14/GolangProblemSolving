package main

import "fmt"

func main28() {
	var a int = 10
	var b int = 20
	fmt.Printf("value of a: %d and value of b:  %d before swapping. \n ", a, b)
	var c int = a
	a = b
	b = c

	fmt.Printf("value of a: %d and value of b:  %d after swapping. \n", a, b)
}

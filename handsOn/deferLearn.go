package main

import "fmt"

func main30() {

	for i := 0; i < 4; i++ {
		fmt.Print(i)
		defer temp(i)
	}
}
func temp(i int) {
	fmt.Print(i)
}

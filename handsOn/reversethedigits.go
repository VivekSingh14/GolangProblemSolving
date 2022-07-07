package main

import "fmt"

func main20() {
	var num int = 327
	var rev int = 0
	for num > 0 {
		remainder := num % 10
		rev = (rev * 10) + remainder
		num = num / 10
	}

	fmt.Printf("number after getting reversed: %d \n", rev)
}

package main

import "fmt"

func main5() {
	var num int
	fmt.Println("Enter the number....")
	fmt.Scanln(&num)
	var digits int = 0
	for num > 0 {
		digits++
		num = num / 10
	}

	fmt.Println(digits)

}

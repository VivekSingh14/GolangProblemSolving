package main

import "fmt"

func main23() {
	var name string = "Hello Vivek!! this is golang terminal."

	for _, ch := range name {
		if ch == '!' {
			fmt.Println()
		}
		fmt.Printf("%c", ch)
	}

	fmt.Println()

	var str string = "HelloJaipur"

	fmt.Println(len(str))

	for i, i2 := range str {
		fmt.Printf("%c  %d \n", i2, i)
	}
}

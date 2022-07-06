package main

import "fmt"

//TODO:  Numbers are called if their digit's cube's sum is equal to the number such as 153 ==>  1*1*1 + 5*5*5 + 3*3*3

func main2() {
	num := 154
	temp := num
	newNum := 0
	for temp > 0 {
		rem := temp % 10
		newNum = newNum + (rem * rem * rem)
		temp = temp / 10
	}
	if num == newNum {
		fmt.Println("Number is armstrong.")
	} else {
		fmt.Println("Number is not armstrong")
	}
}

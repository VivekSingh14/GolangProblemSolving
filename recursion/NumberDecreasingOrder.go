package main

import "fmt"

func main2() {
	num := 9

	PrintInReverse(num)

}

func PrintInReverse(num int) {
	if num > 0 {
		fmt.Println(num)
		PrintInReverse(num - 1)
	}
}

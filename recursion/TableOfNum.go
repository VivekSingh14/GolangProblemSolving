package main

import "fmt"

func main3() {
	count := 10
	num := 2

	TableOfNum(count, num)

}

func TableOfNum(n int, num int) {
	if n > 0 {
		TableOfNum(n-1, num)
		fmt.Printf("%d * %d = %d \n", num, n, num*n)
	}
}

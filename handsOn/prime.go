package main

import "fmt"

func main17() {
	var count int
	for i := 2; i <= 30; i++ {
		count = 0
		for j := 2; j < i/2; j++ {
			if i%j == 0 {
				count++

			}
		}
		if count == 0 {
			fmt.Println(i)
		}
	}
}

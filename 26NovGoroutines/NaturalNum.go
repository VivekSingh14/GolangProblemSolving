package main

import "fmt"

func main1() {

	ch := make(chan int)

	go printNumber(ch)

	for i := 1; i <= 10; i++ {
		ch <- i
	}

}

func printNumber(ch chan int) {

	for {
		select {
		case num := <-ch:
			{
				fmt.Println(num)
			}
		}
	}
}

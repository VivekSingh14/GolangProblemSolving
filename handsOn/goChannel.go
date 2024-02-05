package main

import "fmt"

func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
	ch <- 1
}
func main32() {
	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23
	num := <-ch
	fmt.Println("End Main method ", num)
}

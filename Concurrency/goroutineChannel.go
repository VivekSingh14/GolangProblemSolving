package main

import "fmt"

func sumRoutine(num chan int) {
	fmt.Println(23 + <-num)
	//num <- 23 + <-num
	// temp := 23 + <-num
	// fmt.Println(temp)
	close(num)
}

func main() {

	up := make(chan int)

	go sumRoutine(up)

	up <- 3

	res, ok := <-up

	fmt.Println("End of the main function.", res, " ", ok)
}

package main

import (
	"fmt"
	"sync"
)

func main10() {

	fmt.Println("Before Goroutine.")
	var wg sync.WaitGroup
	wg.Add(1)
	temp := make(chan int)

	//time.Sleep(time.Second)
	go func(wg *sync.WaitGroup, temp chan int) {
		defer wg.Done()
		fmt.Println("Inside Goroutine", " : ", <-temp)
		temp <- -1
	}(&wg, temp)
	temp <- 1
	res := <-temp
	fmt.Println(res)

	fmt.Println("After Goroutine.")
	wg.Wait()

}

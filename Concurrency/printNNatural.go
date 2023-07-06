package main

import (
	"fmt"
	"sync"
)

var resNum = 0

func main15() {

	numtemp := make(chan int)
	//var numtemp chan int

	var wg sync.WaitGroup

	wg.Add(1)

	go goroutineresNum(&wg, numtemp)

	numtemp <- 1

	val := <-numtemp

	//close(numtemp)

	fmt.Println(val)

	wg.Wait()

}

func goroutineresNum(wg *sync.WaitGroup, numtemp chan int) {
	defer wg.Done()
	random := <-numtemp
	fmt.Println("from channel: ", random)

	for i := 0; i < 10; i++ {
		resNum = resNum + 1
	}
	numtemp <- resNum

}

package main

import (
	"fmt"
	"sync"
)

var n1 = 0
var n2 = 1

func main() {

	num := make(chan int)
	var wg sync.WaitGroup
	var sy sync.Mutex

	fmt.Print(n1, " ", n2, " ")
	for i := 3; i <= 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, sy *sync.Mutex, temp chan int) {
			defer wg.Done()
			sy.Lock()
			res := <-num
			n1 = n2
			n2 = res
			fmt.Print(n2, " ")
			sy.Unlock()
			num <- 0

		}(&wg, &sy, num)

	}

	for i := 3; i <= 5; i++ {
		num <- n1 + n2
		<-num
	}

	fmt.Println()

	wg.Wait()

}

package main

import (
	"fmt"
	"sync"
)

var count1 int

func main9() {

	var wg sync.WaitGroup

	wg.Add(2)

	go goroutine1(&wg)

	go goroutine2(&wg)
	fmt.Println("before wait")
	wg.Wait()

	fmt.Println("after wait")

}

func goroutine1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		count1 = count1 + 1
		fmt.Println(count1, " : gorountine1 ")

	}
}

func goroutine2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		count1 = count1 + 2
		fmt.Println(count1, " : gorountine2 ")

	}
}

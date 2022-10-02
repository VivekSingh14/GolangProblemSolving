package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool)

	wg.Add(1) // 2
	go print_even(ch, &wg)
	go print_odd(ch, &wg)
	ch <- true

	wg.Wait()
}

func print_even(ch chan bool, wg *sync.WaitGroup) {

	for i := 2; i < 11; i = i + 2 {
		<-ch
		//time.Sleep(1 * time.Second)
		fmt.Println(i, " from print_even")
		ch <- true
	}
	defer wg.Done()
}

func print_odd(ch chan bool, wg *sync.WaitGroup) {
	for i := 1; i < 11; i = i + 2 {
		<-ch
		//time.Sleep(1 * time.Second)
		fmt.Println(i, " from print_odd")
		ch <- true
	}
	defer wg.Done()
}

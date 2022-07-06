package main

import (
	"fmt"
	"sync"
)

var GFG int = 0

func worker(wg *sync.WaitGroup, wq *sync.Mutex) {
	wq.Lock()
	GFG = GFG + 1
	wq.Unlock()
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	var wq sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &wq)
	}
	wg.Wait()
	fmt.Println("Value of GFG: ", GFG)
}

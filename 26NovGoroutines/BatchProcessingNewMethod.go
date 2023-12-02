package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make([]int, 0, 100)
	for n := 0; n < 100; n++ {
		data = append(data, n)
	}
	batch(data)
}

const batchSize1 = 10

func batch(data []int) {
	ch := make(chan struct{}, batchSize1)
	var wg sync.WaitGroup
	for _, i := range data {
		wg.Add(1)
		ch <- struct{}{}
		x := i
		go func() {
			defer wg.Done()
			// do more complex things here
			fmt.Println(x)
			<-ch
		}()
	}
	wg.Wait()
	fmt.Println("done processing all data")
}

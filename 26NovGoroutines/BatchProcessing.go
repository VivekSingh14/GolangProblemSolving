package main

import (
	"fmt"
	"sync"
)

const batchSize = 10

func main3() {
	data := make([]int, 0, 100)
	for n := 0; n < 100; n++ {
		data = append(data, n)
	}
	process(data)
}

func process(data []int) {
	for start, end := 0, 0; start <= len(data)-1; start = end {
		end = start + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[start:end]
		processBatch(batch)
	}
	fmt.Println("done processing all data")
}

func processBatch(batch []int) {
	var wg sync.WaitGroup

	for _, i := range batch {
		x := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(x)
		}()

	}
}

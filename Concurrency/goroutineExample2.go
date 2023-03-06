package main

import (
	"fmt"
	"sync"
)

func main2() {
	var wg sync.WaitGroup
	wg.Add(1)
	go count("vivek", &wg)
	wg.Wait()

}

func count(value string, wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, value)
	}
	wg.Done()
}

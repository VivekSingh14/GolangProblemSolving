package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func main9() {

	// resCh := make(chan string, 2)

	// var waitGroup sync.WaitGroup
	// waitGroup.Add(2)

	// go func(ch chan<- string, wg *sync.WaitGroup) {
	// 	defer wg.Done()
	// 	time.Sleep(1 * time.Second)
	// 	ch <- "Hello"
	// }(resCh, &waitGroup)

	// go func(ch chan<- string, wg *sync.WaitGroup) {
	// 	defer wg.Done()
	// 	time.Sleep(1 * time.Second)
	// 	ch <- "World"
	// }(resCh, &waitGroup)

	// waitGroup.Wait()

	// for len(resCh) > 0 {
	// 	res := <-resCh
	// 	fmt.Println(res)
	// }

	bufferSize := 5
	wordCh := make(chan string, bufferSize)
	words := strings.Split("the quick brown fox jumped over the lazy dog", " ")

	var waitGroup sync.WaitGroup

	// what is wrong here?
	for _, word := range words {
		waitGroup.Add(1)
		go func(ch chan<- string, wg *sync.WaitGroup, someword string) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			ch <- someword
		}(wordCh, &waitGroup, word)
	}

	done := make(chan bool)
	go func(d chan bool) {
		defer close(d)
		for res := range wordCh {
			fmt.Print(fmt.Sprintf("%s ", res))
		}
	}(done)

	waitGroup.Wait()
	close(wordCh)
	<-done

}

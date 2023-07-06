package main

import (
	"fmt"
	"sync"
	"time"
)

var count2 int

func main23() {
	var mut sync.Mutex

	go goroutine3(&mut)

	go goroutine4(&mut)
	time.Sleep(10 * time.Second)

}

func goroutine3(mut *sync.Mutex) {

	for i := 0; i < 10; i++ {
		mut.Lock()
		count2 = count2 + 1
		fmt.Println(count2, " : gorountine3 ")
		mut.Unlock()
	}

	//time.Sleep(2 * time.Second)
}

func goroutine4(mut *sync.Mutex) {

	for i := 0; i < 10; i++ {
		mut.Lock()
		count2 = count2 + 2
		fmt.Println(count2, " : gorountine4 ")
		mut.Unlock()
	}

	//time.Sleep(2 * time.Second)
}

package main

import (
	"fmt"
	"time"
)

func main7() {
	go sum(3, 2)
	time.Sleep(1 * time.Second)
	sum(4, 5)

}

func sum(a int, b int) {
	//time.Sleep(1 * time.Second)
	fmt.Println(a + b)
}

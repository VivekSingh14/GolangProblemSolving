package main

import (
	"fmt"
	"time"
)

func main() {
	go sum(3, 2)
	sum(4, 5)

}

func sum(a int, b int) {
	time.Sleep(1 * time.Second)
	fmt.Println(a + b)
}

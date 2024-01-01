package main

import (
	"fmt"
	"time"
)

func execute1(some string) {
	for i := 1; true; i++ {
		fmt.Println(some, i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	go execute1("Vivek")

	go execute1("Singh")

	fmt.Println("Program ends successfully.")
	time.Sleep(time.Millisecond * 1000)
}

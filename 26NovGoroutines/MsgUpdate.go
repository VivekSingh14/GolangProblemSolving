package main

import "fmt"

func main() {

	first := make(chan string)
	sec := make(chan string)

	go ping(first, sec)
	go pingHacked(first, sec)

	fmt.Println(<-sec)

}

func ping(first chan string, sec chan string) {
	first <- "Hello"
}

func pingHacked(first chan string, sec chan string) {

	temp := <-first

	temp = temp + " hacked msg"
	sec <- temp
}

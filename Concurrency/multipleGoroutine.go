package main

import (
	"fmt"
	"time"
)

func Aname() {
	arr1 := [4]string{"Vivek", "Singh", "Jaipur", "Pune"}

	for i := 0; i < len(arr1); i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(arr1[i])
	}
}

func Aid() {
	arr1 := [4]int{101, 102, 103, 104}

	for i := 0; i < len(arr1); i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(arr1[i])
	}
}

func main() {
	fmt.Println(".....Main goroutine starts......")

	go Aname()

	go Aid()

	time.Sleep(2003 * time.Millisecond)

	fmt.Println(".....Main goroutine stops.......")
}

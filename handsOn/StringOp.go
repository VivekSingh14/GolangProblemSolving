package main

import (
	"fmt"
	"strings"
)

func main24() {
	str := "My  name is  Vivek"

	fmt.Println(str)

	fmt.Println("********************")

	str = strings.Join(strings.Fields(strings.TrimSpace(str)), " ")

	arr := strings.SplitAfter(str, " ")

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
	fmt.Println(" ")

	fmt.Println(str)
}

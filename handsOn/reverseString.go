package main

import "fmt"

func main() {
	var str string = "Hello Golang"

	var arr []rune = []rune(str)
	var j int = len(arr) - 1
	for i := 0; i < (len(arr)-1)/2; i++ {
		for j > len(arr)/2 {
			var temp rune = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			j--
		}

	}
	//fmt.Print("Hello debugging")
	fmt.Println(string(arr))

}

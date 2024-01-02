package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}

	var any interface{}

	any = arr

	fmt.Println(any)

	//this doesn't work because any is an empty interface value, it is not an int slice.
	//fmt.Println(len(any))

	//this will work because first we are asserting a type and then checking a length
	fmt.Println(len(any.([]int)))
}

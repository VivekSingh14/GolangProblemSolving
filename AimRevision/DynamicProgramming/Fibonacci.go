package main

import "fmt"

func main2() {

	finalTerm := 5

	arr := make([]int, finalTerm)

	//base cases are created
	arr[0] = 0
	arr[1] = 1

	for i := 2; i < finalTerm; i++ {
		//using the previous result to get the new result
		arr[i] = arr[i-1] + arr[i-2]
	}

	fmt.Println(arr[finalTerm-1])

}

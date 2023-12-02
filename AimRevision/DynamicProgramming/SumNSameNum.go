package main

import "fmt"

func main() {

	N := 3
	term := 5

	arr := make([]int, term+1)
	//base cases are created.
	arr[0] = 0
	arr[1] = N

	for i := 2; i <= term; i++ {
		//using the previous result to get the new result
		// where previous result := arr[i-1]
		arr[i] = arr[i-1] + N
	}

	fmt.Println(arr[term])
}

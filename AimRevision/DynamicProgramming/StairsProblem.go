package main

import "fmt"

func main() {

	stairsTarget := 3

	//NumOfWays := 2

	arr := make([]int, stairsTarget+1)
	//base cases
	arr[0] = 1
	arr[1] = 1
	arr[2] = 2

	for i := 3; i <= stairsTarget; i++ {

		arr[i] = arr[i-1] + arr[i-2]
	}

	fmt.Println(arr[stairsTarget])

}

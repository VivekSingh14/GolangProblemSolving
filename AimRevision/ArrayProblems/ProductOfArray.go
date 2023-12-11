package main

import "fmt"

func main8() {
	arr1 := []int{10, 3, 5, 6, 2}

	// o/p := {180, 600, 360, 300, 900}
	var res []int

	for i := 0; i < len(arr1); i++ {
		temp := 1
		for j := 0; j < len(arr1); j++ {
			if i != j {
				temp *= arr1[j]
			}
		}
		res = append(res, temp)
	}

	fmt.Println(res)

}

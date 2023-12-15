package main

import "fmt"

type data struct {
	start, end int
}

func main11() {
	arr := []data{{1, 3}, {2, 4}, {6, 8}, {9, 10}}

	//o/p := {1,4} {6,8}, {9,10}

	index := 0

	for i := 1; i < len(arr); i++ {
		if arr[index].end > arr[i].start {
			if arr[index].end < arr[i].end {
				arr[index].end = arr[i].end
			}
		} else {
			index++
			arr[index] = arr[i]
		}
	}

	for i := 0; i <= index; i++ {
		fmt.Println("{", arr[i].start, ",", arr[i].end, "}")
	}
}

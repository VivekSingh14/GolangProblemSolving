package main

import "fmt"

func main() {
	arr := []int{2, 3, 6, 3, 9, 4}

	for i := 0; i < 5; i++ {
		for j := i; j < 6; j++ {
			if arr[i] == arr[j] && i != j {
				fmt.Println("Array is having duplicate elements and the duplicate element is: ", arr[i])
				break
			}
		}
	}
}

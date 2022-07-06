package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8, 10, 3, 5, 7, 1, 0}
	sum := 4
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				fmt.Println(arr[i], " ", arr[j])
				count += 1
			}
		}
	}
	fmt.Println("Total number of pairs: ", count)
}

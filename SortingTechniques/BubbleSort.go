package main

import "fmt"

func main1() {
	arr := []int{3, 7, 9, 10, 6, 5, 12, 4, 11, 2}
	sorted := bubbleSort(arr)
	fmt.Println(sorted)
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				//swap(&arr[j], &arr[j+1])
			}
		}
	}
	return arr
}

// func swap(num1 *int, num2 *int) {
// 	temp := *num1
// 	*num1 = *num2
// 	num2 = &temp
// }

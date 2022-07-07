package main

import "fmt"

func main22() {

	var slice []int

	fmt.Println(len(slice))

	slice = append(slice, 4)

	fmt.Println(len(slice))

	var slice2 []int

	slice2 = append(slice2, 6)
	slice2 = append(slice2, 8)
	slice2 = append(slice2, 10)

	slice = append(slice, slice2...)

	fmt.Println(len(slice))

	fmt.Println(slice[0:2])

	fmt.Println("******************")

	var testArr [3]int = [3]int{2, 4, 6}

	fmt.Println(len(testArr), " ", cap(testArr))
	fmt.Println("*******************")

	arr1 := make([]int, 5)

	fmt.Println(arr1)

	arr1 = append(arr1, 3)
	arr1[1] = 4

	fmt.Println(arr1)

	arr1 = append(arr1, 4, 1)

	fmt.Println(arr1)

	var even []int = []int{2, 4}
	var odd []int = []int{1, 3, 5}
	fmt.Println("*****************")
	fmt.Println(even)
	fmt.Println(odd)
	copy(odd, even)

	fmt.Println(odd)
}

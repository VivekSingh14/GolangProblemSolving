package main

import "fmt"

func main4() {
	var number_slice []int

	fmt.Println("slice now: ", number_slice, " length of a slice: ", len(number_slice), " and capacity of an array: ", cap(number_slice))

	number_slice = append(number_slice, 1)

	fmt.Println("slice now: ", number_slice, " length of a slice: ", len(number_slice), " and capacity of an array: ", cap(number_slice))

	number_slice = append(number_slice, 2)

	fmt.Println("slice now: ", number_slice, " length of a slice: ", len(number_slice), " and capacity of an array: ", cap(number_slice))

	number_slice = append(number_slice, 3)
	fmt.Println("slice now: ", number_slice, " length of a slice: ", len(number_slice), " and capacity of an array: ", cap(number_slice))

	number_slice = append(number_slice, 4)

	fmt.Println("slice now: ", number_slice, " length of a slice: ", len(number_slice), " and capacity of an array: ", cap(number_slice))

	number_slice = append(number_slice, 5)

	fmt.Println("slice now: ", number_slice, " length of a slice: ", len(number_slice), " and capacity of an array: ", cap(number_slice))

}

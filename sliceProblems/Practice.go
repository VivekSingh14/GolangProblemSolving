package main

import "fmt"

func main3() {

	a := [5]int{1, 2, 3, 4, 5}

	fmt.Println(a)

	fmt.Println("length of a slice: ", len(a), " and capacity of an array: ", cap(a))

	var b []int = a[1:4]
	fmt.Println(b)

	fmt.Println("length of a slice: ", len(b), " and capacity of an array: ", cap(b))

	fmt.Println("------------------------")

	fmt.Println("Now appending value to b")
	b = append(b, 9)
	fmt.Println(b)
	fmt.Println("length of a slice: ", len(b), " and capacity of an array: ", cap(b))

	fmt.Println("------------------------")
	fmt.Println("Now appending one more value to b")
	b = append(b, 8)
	fmt.Println(b)
	fmt.Println("length of a slice: ", len(b), " and capacity of an array: ", cap(b))

	fmt.Println("------------------------")
	var s []int
	s = append(s, 1)
	fmt.Println("s printing: ", s)
}

package main

import "fmt"

func main() {
	//when you declare the array with size such as below it is called an array
	//And if you declare the array without size such as var arr []string(whether it is initialized or not ) it is called slice.
	var arr [5]string = [5]string{"Scala", "Perl", "Java", "Python", "Go"}

	arr2 := arr

	fmt.Println(arr)

	fmt.Println(arr2)

	arr[2] = "CPP"

	fmt.Println("******************")

	fmt.Println(arr)

	fmt.Println(arr2)

	arr3 := &arr

	fmt.Println("******************")

	fmt.Println(arr)

	fmt.Println(*arr3)

	arr[1] = "Test"

	fmt.Println("******************")

	fmt.Println(arr)

	fmt.Println(*arr3)

}

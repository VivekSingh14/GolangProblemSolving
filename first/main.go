package main

import "fmt"

//adding comment.
func main() {
	// f, err := os.Create("/Users/viveksingh/Documents/WorkDir/repositories/repo2/demo1/first/test.txt")

	// if err != nil {
	// 	panic("cannot create a file")
	// } else {
	// 	fmt.Printf("err = %+v \n", err)
	// }
	// defer f.Close()
	// fmt.Fprintf(f, "Hello")

	arr := []int{1, 2, 3, 4, 5}

	fmt.Println("capacity of an array arr: ", cap(arr))
	fmt.Println("length of an array arr: ", len(arr))

	Temp(&arr)
}

func Temp(arr *[]int) {
	*arr = append(*arr, 4)
	fmt.Println("capacity of an array arr inside Temp: ", cap(*arr))
	fmt.Println("length of an array arr inside Temp: ", len(*arr))
}

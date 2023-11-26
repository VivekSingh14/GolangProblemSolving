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

	// arr := []int{1, 2, 3, 4, 5}

	// fmt.Println("capacity of an array arr: ", cap(arr))
	// fmt.Println("length of an array arr: ", len(arr))

	// Temp(&arr)

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// fmt.Println("before enqueue")

	// fmt.Println(a)

	// a = enQueue(a, 9)
	// a = deQueue(a)
	// fmt.Println("after enqueue and dequeue")
	// fmt.Println(a)

	// fmt.Println("before push")

	// fmt.Println(a)
	// a = push(a, 1)
	// fmt.Println("after push")
	// fmt.Println(a)
	// a = pop(a)

	// fmt.Println("after pop")
	// fmt.Println(a)

	fmt.Println(removeMidnum(a))

	// var b []int = a[1:4]

	// fmt.Println(b)
	// fmt.Println("-------------")
	// fmt.Println("len(b): ", len(b), "cap(b): ", cap(b))

	// b = append(b, 9)

	// fmt.Println("-------------")
	// fmt.Println("-------------")

	// fmt.Println(b)
	// fmt.Println("-------------")
	// fmt.Println("len(b): ", len(b), "cap(b): ", cap(b))

	// b = append(b, 11)

	// fmt.Println("-------------")
	// fmt.Println("-------------")

	// fmt.Println(b)
	// fmt.Println("-------------")
	// fmt.Println("len(b): ", len(b), "cap(b): ", cap(b))

}

func Temp(arr *[]int) {
	*arr = append(*arr, 4)
	fmt.Println("capacity of an array arr inside Temp: ", cap(*arr))
	fmt.Println("length of an array arr inside Temp: ", len(*arr))
}

func deQueue(arr []int) []int {
	return arr[1:]
}

func enQueue(arr []int, data int) []int {
	return append(arr, data)
}

func push(arr []int, data int) []int {
	return append(arr, data)
}

func pop(arr []int) []int {
	return arr[:len(arr)-1]
}

func removeMidnum(arr []int) []int {

	fmt.Println("before operations: len and cap ", len(arr), cap(arr))
	//first method
	// mid := len(arr) / 2
	// tail := arr[mid+1:]
	// arr = arr[:mid]
	// arr = append(arr, tail...)

	//second method, if you want to remove the memory leak

	mid := len(arr) / 2
	arr = append(arr[:mid], arr[mid+1:]...)

	ar := make([]int, len(arr))

	copy(ar, arr)

	//second method ended here

	fmt.Println("after operations: len and cap ", len(ar), cap(ar))
	return ar
}

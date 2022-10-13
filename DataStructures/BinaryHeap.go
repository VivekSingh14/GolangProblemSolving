package main

import "fmt"

func main() {
	heapArr := make([]int, 0)
	heapArr = append(heapArr, 0)
	heapArr = append(heapArr, 40)
	heapArr = append(heapArr, 35)
	heapArr = append(heapArr, 15)
	heapArr = append(heapArr, 30)
	heapArr = append(heapArr, 10)
	heapArr = append(heapArr, 12)
	heapArr = append(heapArr, 6)
	heapArr = append(heapArr, 5)
	heapArr = append(heapArr, 20)

	//to be inserted in heap
	heapArr = append(heapArr, 50)

	fmt.Println(heapArr)
	fmt.Println(len(heapArr))
	fmt.Println(cap(heapArr))
}

func InsertInHeap(arr []int, key int) []int {

	return nil
}

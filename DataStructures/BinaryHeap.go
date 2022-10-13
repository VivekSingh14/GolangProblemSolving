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
	heapArr = append(heapArr, 45)

	//fmt.Println(heapArr[10])
	//fmt.Println(len(heapArr))
	//fmt.Println(cap(heapArr))

	var res []int

	for i := 2; i < len(heapArr); i++ {
		res = InsertInHeap(heapArr, i)
	}

	fmt.Println(res)
}

func InsertInHeap(arr []int, n int) []int {

	i := n
	temp := arr[n]

	for i > 1 && temp > arr[i/2] {

		arr[i] = arr[i/2]

		i = i / 2
	}
	arr[i] = temp
	return arr
}

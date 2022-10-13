package main

import "fmt"

func main() {
	heapArr := make([]int, 0)
	heapArr = append(heapArr, 0)
	heapArr = append(heapArr, 14)
	heapArr = append(heapArr, 15)
	heapArr = append(heapArr, 5)
	heapArr = append(heapArr, 20)
	heapArr = append(heapArr, 30)
	heapArr = append(heapArr, 8)

	//to be inserted in heap
	heapArr = append(heapArr, 40)

	//fmt.Println(heapArr[10])
	//fmt.Println(len(heapArr))
	//fmt.Println(cap(heapArr))

	var res []int

	for i := 2; i < len(heapArr); i++ {
		res = InsertInMaxHeap(heapArr, i)
	}
	fmt.Println("after insertion.")
	fmt.Println(res)

	fmt.Println("after deletion")
	resu := DeleteFromMaxHeap(heapArr, 7)
	fmt.Println(resu)

	resu = DeleteFromMaxHeap(heapArr, 6)
	fmt.Println(resu)

	resu = DeleteFromMaxHeap(heapArr, 5)
	fmt.Println(resu)

	// res = DeleteFromMaxHeap(heapArr, 7)
	// fmt.Println(res)

	// res = DeleteFromMaxHeap(heapArr, 6)
	// fmt.Println(res)

	// res = DeleteFromMaxHeap(heapArr, 5)
	// fmt.Println(res)
}

func InsertInMaxHeap(arr []int, n int) []int {

	i := n
	temp := arr[n]

	for i > 1 && temp > arr[i/2] {

		arr[i] = arr[i/2]

		i = i / 2
	}
	arr[i] = temp
	return arr
}

func DeleteFromMaxHeap(arr []int, n int) int {
	val := arr[1]
	x := arr[n]
	arr[1] = arr[n]
	arr[n] = val
	i := 1
	j := 2 * i

	for j <= n-1 {
		if j < n-1 && arr[j+1] > arr[j] {
			j = j + 1
		}

		if arr[i] < arr[j] {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp

			i = j
			j = 2 * j
		} else {
			break
		}
	}
	return x
}

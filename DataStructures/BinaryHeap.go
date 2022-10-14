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

func DeleteFromMaxHeap(A []int, n int) int {
	var i int
	var j int
	//var x int
	var temp int
	var val int
	val = A[1]
	//x = A[n]
	A[1] = A[n]
	A[n] = val
	i = 1
	j = i * 2
	for j <= n-1 {
		if j < n-1 && A[j+1] > A[j] {
			j = j + 1
			if A[i] < A[j] {
				temp = A[i]
				A[i] = A[j]
				A[j] = temp
				i = j
				j = 2 * j
			} else {
				break
			}
		}
	}
	return val
}

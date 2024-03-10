package main

import "fmt"

func main() {
	//arr := []int{0, 10, 20, 30, 25, 5, 40, 35}

	//newHeap := CreateMaxHeap(arr)

	//fmt.Println(newHeap[1:])

	arr := []int{0, 30, 20, 15, 5, 10, 12, 6}
	ele := 40
	fmt.Println(InsertionIntoMaxHeap2(arr, ele))
	fmt.Println("---------------------")
	arr2 := []int{0, 30, 20, 15, 5, 10, 12, 6}
	fmt.Println(RemoveFromHeap(arr2))
	fmt.Println("---------------------")
	arr3 := []int{0, 5, 10, 30, 20, 35, 40, 15}
	fmt.Println(Heapify(arr3))

}

func CreateMaxHeap(arr []int) []int {
	var res []int
	for i := 2; i < len(arr); i++ {

		res = InsertionIntoMaxHeap2(arr, i)
	}
	return res
}

func InsertionIntoMaxHeap2(arr []int, n int) []int {
	arr = append(arr, n)
	i := len(arr) - 1
	fmt.Println(i)
	temp := arr[len(arr)-1]

	for i > 1 && temp > arr[i/2] {
		arr[i] = arr[i/2]
		i = i / 2
	}
	arr[i] = temp
	return arr
}

// we cannot directly remove the element from heap.
// step1: remove the root element
// step2: put the last element of an array into root place
// step3: now start comparing it with its children node whichever in both, will be compared with it and replace it.
// and so on untill we reached at end or node without children
func RemoveFromHeap(arr []int) []int {

	//temp := arr[1]

	arr[1] = arr[len(arr)-1]
	i := 1
	j := 2 * i
	for j < len(arr)-1 {
		if arr[j+1] > arr[j] {
			j = j + 1
		}
		if arr[i] < arr[j] {
			tempo := arr[i]
			arr[i] = arr[j]
			arr[j] = tempo
			i = j
			j = 2 * j
		} else {
			break
		}
	}
	//arr[len(arr)-1] = temp

	return arr[1 : len(arr)-1]
}

// heapify is related to creation of heap, not insertion into heap
func Heapify(arr []int) []int {
	i := len(arr) - 1
	r := 1
	for i >= 1 && r >= 1 {
		r = i / 2
		if arr[r] < arr[2*r] || arr[r] < arr[2*r+1] {
			if arr[2*r] >= arr[2*r+1] {
				temp := arr[2*r]
				arr[2*r] = arr[r]
				arr[r] = temp
			} else {
				temp := arr[2*r+1]
				arr[2*r+1] = arr[r]
				arr[r] = temp
			}

		}
		i--
	}
	return arr
}

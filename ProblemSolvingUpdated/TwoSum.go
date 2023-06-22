package main

import "fmt"

func main() {
	arr := []int{11, 1, 2, 7}
	target := 9

	res := twoSum(arr, target)

	fmt.Println(res)

}

func twoSum(arr []int, target int) []int {
	var temp []int
	for i := 0; i < len(arr)-1; i++ {
		if target == arr[i]+arr[i+1] {
			temp = append(temp, i)
			temp = append(temp, i+1)
		}
	}

	return temp
}

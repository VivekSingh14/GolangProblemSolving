package main

import "fmt"

func main38() {
	arr := []int{1, 4, 2, 5, 3}
	fmt.Println(sumOddLengthSubarrays(arr))

}

func sumOddLengthSubarrays(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	var count int

	for i := 0; i < len(arr); i++ {
		count += arr[i]
	}

	if len(arr)%2 == 1 {
		count += count
	}
	odd := 0
	i := 0
	for odd < len(arr) {

		// for j := 0; j < odd; j++ {
		// 	count = count+arr[j]
		// }

		for j := 0; j < len(arr); j++ {
			count = count + arr[i+odd] + arr[i+1+odd] + arr[i+2+odd]
		}
		odd = odd + 1
	}

	return count
}

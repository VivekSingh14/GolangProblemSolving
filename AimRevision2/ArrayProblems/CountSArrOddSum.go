package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(numOfSubarrays(arr))
	fmt.Println(numOfSubarraysOpt(arr))

}
func numOfSubarrays(arr []int) int {
	count := 0

	temp_sum := 0
	start := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			count++
		}
		temp_sum += arr[i]
		if temp_sum%2 != 0 {
			count++
		} else {
			temp_sum -= arr[start]
			start++
		}
	}
	return count
}

func numOfSubarraysOpt(arr []int) int {

	summary := make([]int, 0, len(arr))
	current := 0
	for _, v := range arr {
		current = (current + v) % 2
		summary = append(summary, current)
	}
	odd, even := 0, 0
	number := 0
	for _, v := range summary {
		if v%2 == 0 {
			even++
			number += odd
		} else {
			odd++
			number += even + 1
		}
		number %= 1000000007
	}
	return number
}

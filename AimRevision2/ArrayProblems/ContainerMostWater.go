package main

import "fmt"

func main2() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(containerWithMostWater(height))
}

func containerWithMostWater(height []int) int {

	start := 0
	end := len(height) - 1

	max := 0

	for start < end {
		length := min1(height[start], height[end])
		if (end-start)*length > max && length == height[start] {
			max = (end - start) * length
			start++
		} else if (end-start)*length > max && length == height[end] {
			max = (end - start) * length
			end--
		} else {
			if length == height[start] {
				start++
			} else {
				end--
			}
		}
	}
	return max
}

func min1(a, b int) int {
	if a > b {
		return b
	}
	return a
}

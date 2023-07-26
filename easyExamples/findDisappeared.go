package main

import (
	"fmt"
)

func main36() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}

	fmt.Println(findDisappearedNumbers(nums))
}

func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]int, len(nums))

	for _, v := range nums {
		m[v]++
	}

	var missing []int
	for i, _ := range nums {
		if _, ok := m[i+1]; !ok {
			missing = append(missing, (i + 1))

		}
	}
	return missing

}

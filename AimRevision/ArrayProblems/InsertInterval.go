package main

import "fmt"

func main() {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}
	fmt.Println(insert(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	for i, interval := range intervals {
		if newInterval[1] < interval[0] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			return res
		}
		if newInterval[0] > interval[1] {
			res = append(res, interval)
			continue
		}
		newInterval[0] = Min(newInterval[0], interval[0])
		newInterval[1] = Max(newInterval[1], interval[1])
	}
	return append(res, newInterval)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"fmt"
	"sort"
)

func main20() {
	arr := [][]int{{0, 30}, {5, 10}, {15, 20}}
	minMeetingRooms(arr)
}
func minMeetingRooms(arr [][]int) {
	var st []int
	var end []int
	for i := 0; i < len(arr); i++ {
		st = append(st, arr[i][0])
		end = append(end, arr[i][1])
	}

	first := 0
	last := 0
	sort.Ints(st)
	sort.Ints(end)
	//var res int
	var count int

	for first < len(st) && last < len(end) {
		if st[first] < end[last] {
			first++
			count++
		} else if st[first] >= end[last] {
			last++
			count--
		} else {
			first++
			last++
		}
		//res = Maxa(res, count)
	}
	fmt.Println(count)

}

func Maxa(a, b int) int {
	if a > b {
		return a
	}
	return b
}

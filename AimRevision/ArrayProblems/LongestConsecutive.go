package main

import "fmt"

func main22() {
	//arr := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	arr := []int{100, 4, 200, 1, 3, 2}

	//fmt.Println(longestSequence(arr))
	fmt.Println(longestConsecutive(arr))
}

func longestSequence(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		m[n] = struct{}{}
	}
	var longest int
	for n := range m {
		if _, ok := m[n-1]; !ok {
			length := 1
			for _, ok := m[n+length]; ok; _, ok = m[n+length] {
				length++
			}
			longest = maxim(longest, length)
		}
	}
	return longest
}

func maxim(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, num := range nums {
		set[num] = true
	}

	res := 0
	for _, num := range nums {
		if set[num-1] {
			continue
		}

		sequence, temp := 1, num+1
		for set[temp] {
			sequence++
			temp++
		}

		res = maxim(res, sequence)
		if sequence < len(nums)-1 {
			break
		}
	}
	return res
}

package main

import (
	"fmt"
)

func main4() {
	str := "abcbdbdbbdcdabd"
	k := 2

	longestSubstring(str, k)
}

func longestSubstring(str string, k int) {
	curr_min := 0
	curr_end := 0
	max_window_size := 1
	max_window_start := 0
	count := make([]int, 26)
	count[str[0]-'a']++
	for i := 1; i < len(str); i++ {
		count[str[i]-'a']++
		curr_end++
		for !checK(count, k) {
			count[str[curr_min]-'a']--
			curr_min++
		}
		if curr_end-curr_min+1 > max_window_size {
			max_window_size = curr_end - curr_min + 1
			max_window_start = curr_min
		}

	}

	fmt.Println(max_window_size)
	fmt.Println(max_window_start)
	fmt.Println(str[max_window_start : max_window_size+max_window_start])
}

func checK(count []int, k int) bool {
	val := 0
	for i := 0; i < 26; i++ {
		if count[i] > 0 {
			val++
		}
	}
	return k >= val
}

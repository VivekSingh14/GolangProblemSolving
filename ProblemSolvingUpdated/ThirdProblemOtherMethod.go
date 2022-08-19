package main

import "fmt"

const MAX_CHARS1 int = 26

func main18() {
	str := "geeksforgeeks"
	//str := "abacb"
	//k := 3
	Uniques(str)
}

func Uniques(str string) {

	fmt.Println("******************")
	curr_start := 0
	curr_end := 0

	max_window_size := 1
	max_window_start := 0
	count := make([]int, MAX_CHARS1)
	count[str[0]-'a']++

	for i := 1; i < len(str); i++ {
		count[str[i]-'a']++
		curr_end++

		for !isValid1(count) {
			count[str[curr_start]-'a']--
			curr_start++
		}

		if curr_end-curr_start+1 > max_window_size {
			max_window_size = curr_end - curr_start + 1
			max_window_start = curr_start
		}
	}
	fmt.Println(max_window_size)
	fmt.Println(max_window_start)
	fmt.Println(str[max_window_start : max_window_size+max_window_start])
}

func isValid1(count []int) bool {
	//val := 0
	for i := 0; i < MAX_CHARS1; i++ {
		if count[i] > 1 {
			return false
		}
	}
	return true
}

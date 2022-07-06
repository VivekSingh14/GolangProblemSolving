package main

//TODO: SEQ NO: 1
import "fmt"

const MAX_CHARS int = 26

func main16() {
	str := "aabacbebebe"
	//str := "abacb"
	k := 3
	kUniques(str, k)
}

func kUniques(str string, k int) {
	//u := 0
	//count1 := make([]int, MAX_CHARS)
	//for i := 0; i < len(str); i++ {
	//	if count1[str[i]-'a'] == 0 {
	//		u++
	//	}
	//	count1[str[i]-'a']++
	//}
	//if u < k {
	//	fmt.Println("Not enough chars")
	//	return
	//}

	fmt.Println("******************")
	curr_start := 0
	curr_end := 0

	max_window_size := 1
	max_window_start := 0
	count := make([]int, MAX_CHARS)
	count[str[0]-'a']++

	for i := 1; i < len(str); i++ {
		count[str[i]-'a']++
		curr_end++

		for !isValid(count, k) {
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

func isValid(count []int, k int) bool {
	val := 0
	for i := 0; i < MAX_CHARS; i++ {
		if count[i] > 0 {
			val++
		}
	}
	return k >= val
}

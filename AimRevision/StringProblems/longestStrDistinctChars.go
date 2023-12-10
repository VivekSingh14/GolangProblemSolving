package main

import "fmt"

func main5() {
	str := "longestsubstr"

	uniqueString(str)
}

func uniqueString(str string) {

	loc_min := 0
	loc_max := 0
	max_window_size1 := 0
	max_window_star := 0

	countChar := make([]int, 26)
	countChar[str[0]-'a']++

	for i := 1; i < len(str); i++ {
		countChar[str[i]-'a']++
		loc_max++
		for !isUnique(countChar) {
			countChar[str[loc_min]-'a']--
			loc_min++
		}
		if loc_max-loc_min+1 > max_window_size1 {
			max_window_size1 = loc_max - loc_min + 1
			max_window_star = loc_min
		}
	}

	fmt.Println(max_window_size1)
	fmt.Println(max_window_star)
	fmt.Println(str[max_window_star : max_window_size1+max_window_star])
}

func isUnique(countchars []int) bool {
	for i := 0; i < len(countchars); i++ {
		if countchars[i] > 1 {
			return false
		}
	}
	return true
}

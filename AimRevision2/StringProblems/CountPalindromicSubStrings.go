package main

import "fmt"

func main() {
	s := "aaa"
	palindromicSubstrings(s)
	palindromicSubstringsOptimized(s)
}

func palindromicSubstrings(s string) {

	if len(s) <= 1 {
		fmt.Println(len(s))
	}
	count := len(s)

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if isPalindrom(s, i, j) {
				count++
			}
		}
	}

	fmt.Println(count)

}

func isPalindrom(s string, st, end int) bool {
	for st <= end {
		if s[st] != s[end] {
			return false
		}
		st++
		end--
	}
	return true
}

func palindromicSubstringsOptimized(str string) {

	res := 0

	for i := 0; i < len(str); i++ {

		l, r := i, i

		for l >= 0 && r < len(str) && str[l] == str[r] {
			res++
			l--
			r++
		}

		l, r = i, i+1

		for l >= 0 && r < len(str) && str[l] == str[r] {
			res++
			l--
			r++
		}
	}

	fmt.Println("----------------------------")
	fmt.Println(res)
}

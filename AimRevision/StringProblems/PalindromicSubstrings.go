package main

import "fmt"

func main() {
	s := "aaa"
	palindromicSubstrings(s)
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

package main

import "fmt"

func main11() {
	s := "babad"
	longestPalindromicString(s)

}

func longestPalindromicString(s string) {
	maxlen := 0
	start := 0
	end := 0
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if IsPalindrom(s, i, j) && maxlen < j-i+1 {
				maxlen = j - i + 1
				start = i
				end = j
			}
		}
	}
	res := []rune(s)
	res = res[start : end+1]
	fmt.Println(string(res))
}

func IsPalindrom(s string, st, end int) bool {
	for st <= end {
		if s[st] != s[end] {
			return false
		}
		st++
		end--
	}
	return true
}

// func longestPalindrome(s string) string {
// 	len := len(s)
// 	if len == 0 {
// 		return ""
// 	}

// 	var l, r, pl, pr int
// 	for r < len {
// 		// gobble up dup chars
// 		for r+1 < len && s[l] == s[r+1] {
// 			r++
// 		}
// 		// find size of this palindrome
// 		for l-1 >= 0 && r+1 < len && s[l-1] == s[r+1] {
// 			l--
// 			r++
// 		}
// 		if r-l > pr-pl {
// 			pl, pr = l, r
// 		}
// 		// reset to next mid point
// 		l = (l+r)/2 + 1
// 		r = l
// 	}
// 	return s[pl : pr+1]
// }

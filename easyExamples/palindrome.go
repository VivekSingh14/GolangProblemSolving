package main

import (
	"fmt"
	"strings"
)

func main14() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
	fmt.Println(isPalindromeOptimized(s))
}

func isPalindrome(s string) bool {

	s = strings.ToLower(s)
	newstr := ""

	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			newstr += string(s[i])
		}
	}
	fmt.Println(newstr)
	j := len(newstr) - 1
	i := 0
	for i < len(newstr)/2 {
		if newstr[i] != newstr[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindromeOptimized(s string) bool {
	left, right := 0, len(s)-1

	lowercase := strings.ToLower(s)
	for left <= right {
		if !isValidChar(lowercase[left]) {
			left++
			continue
		}

		if !isValidChar(lowercase[right]) {
			right--
			continue
		}

		if lowercase[left] != lowercase[right] {
			return false
		} else {
			left++
			right--
		}
	}

	return true
}

func isValidChar(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

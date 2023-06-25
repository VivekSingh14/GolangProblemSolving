package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
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

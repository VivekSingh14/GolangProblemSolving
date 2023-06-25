package main

import (
	"fmt"
)

func main11() {
	haystack := "hello"
	needle := "ll"
	// it should return 1

	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	j := 0
	i := 0
	length := len(needle)
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			for j = 1; j < length; j++ {
				if needle[j] != haystack[i] {
					break
				}
				i = i + 1
			}
			if j == length {
				return i - length
			}
		} else {
			i = i + 1
		}
	}

	return -1
}

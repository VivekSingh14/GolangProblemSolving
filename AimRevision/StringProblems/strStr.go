package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "mississippi"
	needle := "issip"

	fmt.Println(strStr(haystack, needle))

	fmt.Println(Strstr(haystack, needle))

}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func Strstr(haystack string, needle string) int {
	size := len(needle)
	var res int
	for i := 0; i <= len(haystack)-size; i++ {
		if haystack[i] == needle[0] {
			res = i
			j := 0
			for j = 0; j < size; j++ {
				if haystack[i+j] != needle[j] {
					//&& ((len(haystack) - i) >= size)
					break
				}
			}
			if j == size {
				return res
			}
		}
	}
	return -1
}

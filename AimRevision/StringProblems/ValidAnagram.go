package main

import (
	"fmt"
	"sort"
)

func main7() {
	str1 := "listen"
	str2 := "silent"

	fmt.Println(IsAnagram(str1, str2))
}

func IsAnagram(str1 string, str2 string) bool {
	r1 := []rune(str1)
	r2 := []rune(str2)

	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})

	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})

	return string(r1) == string(r2)
}

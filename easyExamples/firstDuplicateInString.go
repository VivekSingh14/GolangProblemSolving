package main

import "fmt"

func main20() {

	str := "1abBcdefcghi1fxyAtB"
	fmt.Println(firstDuplicateChar(str))
}

func firstDuplicateChar(s string) string {
	var arr [256]int

	if len(s) == 1 {
		return s
	}

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			arr[s[i]] = arr[s[i]] + 1
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			arr[s[i]] = arr[s[i]] + 1
		} else {
			arr[s[i]] = arr[s[i]] + 1
		}
	}
	//fmt.Println(arr)

	for i := 0; i < len(s); i++ {
		if arr[s[i]] == 2 {

			return string(s[i])
		}
	}
	return ""
}

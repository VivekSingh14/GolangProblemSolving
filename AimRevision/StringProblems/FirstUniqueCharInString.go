package main

import "fmt"

func main2() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {
	var arr [26]int

	for i := 0; i < len(s); i++ {
		if arr[s[i]-97] == 0 {
			arr[s[i]-97] = 1
		} else {
			arr[s[i]-97] = arr[s[i]-97] + 1
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[s[i]-97] == 1 {
			return i
		}
	}
	return 0
}

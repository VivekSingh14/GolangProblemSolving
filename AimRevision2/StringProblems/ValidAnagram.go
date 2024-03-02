package main

import "fmt"

func main1() {

	t := "tan"
	p := "ana"
	fmt.Println(validAnagram(t, p))
}

// all test cases won't pass such as "ac", "bb"
func validAnagram(t string, p string) bool {
	if len(t) != len(p) {
		return false
	}

	count1 := 0
	count2 := 0

	for i := 0; i < len(t); i++ {
		count1 += int(t[i])
		count2 += int(p[i])
	}
	return count1 == count2
}

//all test case will pass
func validAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var arr [26]int
	var arr1 [26]int

	for i := 0; i < len(s); i++ {
		arr[s[i]-97] = arr[s[i]-97] + 1
		arr1[t[i]-97] = arr1[t[i]-97] + 1
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != arr1[i] {
			return false
		}
	}

	return true
}

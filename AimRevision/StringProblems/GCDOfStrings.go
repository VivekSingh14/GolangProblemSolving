package main

import "fmt"

func main() {
	str1 := "ABCABC"
	str2 := "ABC"

	fmt.Println(gcdOfStrings(str1, str2))
}

func gcdOfStrings(s1 string, s2 string) string {
	if s1+s2 != s2+s1 {
		return ""
	}
	x := gcd(len(s1), len(s2))
	return s1[:x]
}

func gcd(a, b int) int {
	var temp int
	for b != 0 {
		temp = a % b
		a = b
		b = temp
		//a, b = b, a%b
	}
	return a
}

package main

import "fmt"

func main24() {
	str := "hello"

	str = reverseVowels(str)

	fmt.Println(str)
}

func reverseVowels(s string) string {

	if len(s) < 2 {
		return s
	}

	var vowel []rune
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			vowel = append(vowel, rune(s[i]))
		}
	}

	res := ""
	j := len(vowel) - 1

	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			res += string(vowel[j])
			j--
		} else {
			res += string(s[i])
		}
	}
	return res
}

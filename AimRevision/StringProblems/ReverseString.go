package main

import "fmt"

func main1() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
}

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1

	for i < j {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
		i++
		j--
	}

	fmt.Println(string(s))
}

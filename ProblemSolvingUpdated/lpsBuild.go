package main

import "fmt"

func main() {
	str := "AAAAA"
	M := len(str)
	var lps [5]int

	length := 0
	i := 1
	lps[0] = 0

	for i < M {
		if str[i] == str[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = length
				i++
			}
		}
	}
	fmt.Println(lps)
}

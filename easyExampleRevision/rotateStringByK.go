package main

import (
	"fmt"
)

func main() {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseStr(s, k))
}

func reverseStr(s string, k int) string {
	runeArray := []rune(s)

	ln := len(runeArray)
	if ln == 1 {
		return s
	}
	if ln < k {
		k = k % ln
	}
	res1 := runeArray[ln-k:]
	res2 := runeArray[:ln-k]
	res1 = append(res1, res2...)
	return string(res1)
}

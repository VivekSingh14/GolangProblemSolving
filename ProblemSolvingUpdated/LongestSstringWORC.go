package main

import "fmt"

//TODO: SEQ NO: 3
//TODO: Longest substring without repeating characters

func main19() {
	str := "geeksforgeeks"
	longestUniqueSubsttr(str)
}

func longestUniqueSubsttr(str string) {
	len1 := len(str)
	st := 0
	var i int
	currlen := 0
	maxlen := 0
	start := 0
	map1 := make(map[rune]int)

	map1[rune(str[0])] = 0

	for i = 1; i < len1; i++ {
		if val, ok := map1[rune(str[i])]; ok {
			if val >= st {
				currlen = i - st
				if maxlen < currlen {
					maxlen = currlen
					start = st
				}
				st = val + 1
			}
		} else {
			map1[rune(str[i])] = i
		}
	}
	if maxlen < i-st {
		maxlen = i - st
		start = st
	}
	fmt.Println(start, " ", maxlen)
	fmt.Println(str[start:maxlen])
}

package main

import "fmt"

func main2() {

	S := "AAAABAB"
	K := 1
	fmt.Println(longestLength(S, K))
}

func longestLength(s string, k int) int {
	freq := map[byte]int{}
	maxFreq := 1
	maxLen := 0

	l := 0
	for r := l; r < len(s); r++ {

		freq[s[r]]++

		maxFreq = max(maxFreq, freq[s[r]])

		// if (r-l+1)-maxFreq > k {
		// 	freq[s[l]]--
		// 	l++
		// }

		//maxLen = max(maxLen, r-l+1)
	}

	fmt.Println(freq)
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

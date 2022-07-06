package main

import "fmt"

func main12() {
	text := "This is Text Text"
	pattern := "Text"
	search(text, pattern)
}

func search(txt string, pat string) {
	lenOftext := len(txt)
	lenofPat := len(pat)

	for i := 0; i <= lenOftext-lenofPat; i++ {
		var j int
		for j = 0; j < lenofPat; j++ {
			if txt[i+j] != pat[j] {
				break
			}
		}
		if j == lenofPat {
			fmt.Println("Found at index", i)
		}
	}
}

package main

//TODO: SEQ NO: 2
import "fmt"

const MAX = 256

//TODO: SEQ NO: 2
//TODO: Updated Rabin Karp Algo to find the anagrams into the text files.
/* TODO: Why it is called updated because rabin karp algo usually used to find the pattern from the
TODO: text and here we are trying to find the anagrams. */

func main20() {
	text := "BACDGABCDA"
	pattern := "ABCD"
	searchAna(text, pattern)
}

func searchAna(txt string, pat string) {
	lenT := len(txt)
	lenP := len(pat)

	var countP [MAX]rune
	countTW := [MAX]rune{}

	for i := 0; i < lenP; i++ {
		countP[pat[i]]++
		countTW[txt[i]]++
	}

	for i := lenP; i < lenT; i++ {
		if compare(countP, countTW) {
			fmt.Println("Anagram found at: ", i-lenP)
		}
		countTW[txt[i]]++
		countTW[txt[i-lenP]]--
	}
	if compare(countP, countTW) {
		fmt.Println("Anagram found at: ", lenT-lenP)
	}
}

func compare(arr1 [MAX]rune, arr2 [MAX]rune) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

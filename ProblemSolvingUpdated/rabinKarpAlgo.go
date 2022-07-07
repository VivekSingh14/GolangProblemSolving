package main

import "fmt"

const MAX1 = 256

func main() {
	txt := "BACDGABCDA"
	pat := "ABCD"

	searchAnagram(txt, pat)
}

func searchAnagram(txt string, pat string) {
	lenP := len(pat)
	lenT := len(txt)

	var countP [MAX1]rune
	var countTW [MAX1]rune

	for i := 0; i < lenP; i++ {
		countP[pat[i]]++
		countTW[txt[i]]++
	}

	for i := lenP; i < lenT; i++ {
		if compareOne(countP, countTW) {
			fmt.Println("Anangram is at", i-lenP)
		}
		countTW[txt[i]]++
		countTW[txt[i-lenP]]--

	}
	if compareOne(countP, countTW) {
		fmt.Println("Anagram found at: ", lenT-lenP)
	}
}

func compareOne(countP [MAX1]rune, countTW [MAX1]rune) bool {
	for i := 0; i < len(countP); i++ {
		if countP[i] != countTW[i] {
			return false
		}
	}
	return true
}

package main

//TODO:  Anagram of string and Bubble sort is covered.
import "fmt"

func main() {
	str1 := "silent"
	str2 := "listen"
	charArr1 := []rune(str1)
	charArr2 := []rune(str2)

	for i := 0; i < len(charArr1)-1; i++ {
		for j := i + 1; j < len(charArr1); j++ {
			if string(charArr1[i]) > string(charArr1[j]) {
				temp := charArr1[i]
				charArr1[i] = charArr1[j]
				charArr1[j] = temp
			}
		}
	}

	for i := 0; i < len(charArr2)-1; i++ {
		for j := i + 1; j < len(charArr2); j++ {
			if string(charArr2[i]) > string(charArr2[j]) {
				temp := charArr2[i]
				charArr2[i] = charArr2[j]
				charArr2[j] = temp
			}
		}
	}
	fmt.Println()
	if string(charArr1) == string(charArr2) {
		fmt.Println("Strings are armstrong")
	} else {
		fmt.Println("String are not armstrong")
	}

}

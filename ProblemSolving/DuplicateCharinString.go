package main

import "fmt"

func main() {
	str := "adcbaefedhhibh"
	var arr [26]int
	charStr := []rune(str)

	for i := 0; i < len(str); i++ {
		fmt.Print(charStr[i], " ")
		arr[charStr[i]-97] += 1
	}

	fmt.Println()
	for i := 0; i < len(arr); i++ {
		if arr[i] == 2 {
			fmt.Print(string(charStr[i]), " ")
		}
	}

	fmt.Println(arr)
}

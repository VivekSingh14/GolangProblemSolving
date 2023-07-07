package main

import (
	"fmt"
)

func main23() {
	//Help generate some jargon by writing a program that converts a long name like
	//Eg1: Three Letter Acronyms (TLA)
	//Eg2: Portable Network Graphics (PNG)

	//str := "Portable Network Graphics"
	str := "323443234 2Network :Graphics"

	//temp := strings.Split(str, " ")

	//var res []rune
	resstr := ""
	counter := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			counter = 0
		}
		if str[i] >= 'A' && str[i] <= 'Z' && counter == 0 {
			resstr += string(str[i])
			counter = 1
		}
	}

	// for i := 0; i < len(temp); i++ {
	// 	tempres := temp[i]
	// 	j := 0
	// 	//res = append(res, ContainsChar(tempres))
	// 	if (tempres[j] >= 'A' && tempres[j] <= 'Z') && j < len(tempres) {
	// 		resstr += string(tempres[j])
	// 		continue
	// 	} else {
	// 		j++
	// 	}

	// }

	//fmt.Println(string(res))

	fmt.Println("----------------")
	fmt.Println(resstr)

}

// func ContainsChar(str string) rune {
// 	//3Porta2ble

// 	for i := 0; i < len(str); i++ {
// 		if str[i] >= 'A' && str[i] <= 'Z' {
// 			return rune(str[i])
// 		}
// 	}
// 	return 0
// }

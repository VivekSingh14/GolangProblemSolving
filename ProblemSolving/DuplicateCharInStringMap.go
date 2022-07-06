package main

import "fmt"

func main() {
	str := "adcbaefedhhibh"
	var map1 = map[rune]int{}

	var map2 = map[rune]int{}

	map2['a'] = 3
	map2['b'] = 4
	charStr := []rune(str)

	var num = 3
	var num2 int
	num2 = 5

	fmt.Println(num2, " :::::: ", num, " ****** ", map2)

	for i := 0; i < len(str); i++ {
		_, ok := map1[charStr[i]]
		if !ok {
			map1[charStr[i]] = 1
		} else {
			count := map1[charStr[i]]
			count += 1
			map1[charStr[i]] = count
		}
	}

	for key, value := range map1 {
		if value == 2 {
			fmt.Print(string(key), " ")
		}
	}

	fmt.Println()

}

package main

import "fmt"

func main5() {
	str := "vivek"
	map1 := make(map[rune]int)
	arr := []rune(str)

	for _, r := range arr {
		if _, ok := map1[r]; ok {
			temp := map1[r]
			temp += 1
			map1[r] = temp
		} else {
			map1[r] = 1
		}
	}

	for r, i := range map1 {
		fmt.Printf("%c  %d \n", r, i)
	}

}

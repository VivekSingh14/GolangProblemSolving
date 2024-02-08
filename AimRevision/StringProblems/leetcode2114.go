package main

import (
	"fmt"
	"strings"
)

func main15() {
	arr := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	fmt.Println(mostWordsFound(arr))
}

func mostWordsFound(sentences []string) int {
	var count int

	for i := 0; i < len(sentences); i++ {
		temp := len(strings.Split(sentences[i], " "))
		if temp > count {
			count = temp
		}
	}
	return count
}

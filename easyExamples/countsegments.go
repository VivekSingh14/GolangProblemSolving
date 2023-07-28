package main

import (
	"fmt"
	"strings"
)

//asss
func main() {
	str := "Hello"
	fmt.Println(countSegments(str))
}

func countSegments(s string) int {

	if len(s) == 0 {
		return 0
	}

	arr := strings.Split(s, " ")

	return len(arr)
}

package main

import (
	"fmt"
	"strconv"
)

func main28() {
	details := []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}

	fmt.Println(countSeniors(details))
}

func countSeniors(details []string) int {
	var count int
	for i := 0; i < len(details); i++ {
		str := ""
		str = string(details[i][11])
		str += string(details[i][12])

		if ok, _ := strconv.Atoi(str); ok > 60 {
			count++
		}
	}
	return count
}

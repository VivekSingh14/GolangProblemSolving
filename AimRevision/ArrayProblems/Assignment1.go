package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main7() {
	a := []string{"801234567", "180234567", "0", "189234567", "891234567", "98", "9"}
	//a := []string{"5421", "245", "1452", "0345", "53", "354"}
	//a := []string{"039", "4", "14", "32", "", "34", "7"}
	fmt.Println(solution1(a))
}

func solution1(E []string) int {

	availability := make([]int, 10)
	for _, employee := range E {
		for _, day := range employee {
			availability[day-'0']++
		}
		//fmt.Println(availability)
	}

	fmt.Println(availability)

	maxEmployees := 0
	for i := 0; i < 10; i++ {
		if len(E) == availability[i] {
			return len(E)
		} else {
			max := -1
			maxIndex := -1
			for j := 0; j < len(availability); j++ {
				if availability[j] > max {
					max = availability[j]
					maxIndex = j
				}
			}
			check(maxIndex, E)
		}
	}
	return maxEmployees
}

func check(a int, E []string) int {
	var res []string
	for i := 0; i < len(E); i++ {
		if !strings.Contains(E[i], strconv.Itoa(a)) {
			res = append(res, E[i])
		}
	}
	fmt.Println(res)
	return 0
}

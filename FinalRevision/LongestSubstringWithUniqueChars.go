package main

import "fmt"

func main2() {

	str := "abbcdafeegh"
	var max, start, end int
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str)-1; j++ {
			repeat := calculatelongstring(str, i, j)
			if repeat != 0 {
				i = repeat
				break
			} else {
				tempmax := j - i + 1
				if tempmax > max {
					max = tempmax
					start = i
					end = j
				}
			}
		}
	}
	fmt.Println(max)
	for s := start; s <= end; s++ {
		fmt.Print(string(str[s]))
	}
	fmt.Println()
}

func calculatelongstring(str string, i int, j int) int {
	hashtable := make([]int, 26)

	for k := i; k <= j; k++ {
		if hashtable[str[k]-97] != 0 {
			return hashtable[str[k]-97]
		} else {
			hashtable[str[k]-97] = k
		}
	}

	return 0

}

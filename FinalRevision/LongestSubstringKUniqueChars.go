package main

import "fmt"

func main1() {
	str := "abcbdbdbbdcdabd"
	k := 3
	var max int
	var start int
	var end int
	for i := 0; i < len(str)-1; i++ {
		for j := i + 1; j < len(str); j++ {
			tempMax := display(str, i, j, k)
			if tempMax > max {
				max = tempMax
				start = i
				end = j
			}

		}
	}
	fmt.Println(max)
	for s := start; s <= end; s++ {
		fmt.Print(string(str[s]))
	}
	fmt.Println()
}

func display(str string, i int, j int, unique int) int {
	hashTable := make([]int, 28)
	for k := i; k <= j; k++ {
		if hashTable[str[k]-97] == 0 {
			hashTable[str[k]-97] = 1
		} else {
			num := hashTable[str[k]-97]
			num = num + 1
			hashTable[str[k]-97] = num
		}
	}
	var count int
	for l := 0; l < len(hashTable); l++ {
		if hashTable[l] > 0 {
			count++
		}
	}

	if count == unique {
		for s := i; s <= j; s++ {
			//fmt.Print(string(str[s]))
		}
		//fmt.Println()
		return j - i + 1
	}
	return 0
}

package main

// /-- Two sum using hash map(optimised)  --/

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}

	target := 9

	hashmap := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if value, ok := hashmap[target-arr[i]]; ok {
			fmt.Println(value, i)
		}
		hashmap[arr[i]] = i
	}
}

package main

import "fmt"

func main() {
	str := "aabacbebebe"
	count := make([]int, 26)

	for i := 0; i < len(str); i++ {
		count[str[i]-'a']++
	}
	for i := 0; i < len(str); i++ {
		//fmt.Print(count[i], " ")
		fmt.Printf("%d %c  ", count[i], str[i])
	}
	fmt.Println()

}

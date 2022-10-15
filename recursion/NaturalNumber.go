package main

import "fmt"

func main1() {
	n := 7
	PrintNatural(n)
}

func PrintNatural(n int) {
	if n > 0 {
		PrintNatural(n - 1)
		fmt.Println(n)
	}
}

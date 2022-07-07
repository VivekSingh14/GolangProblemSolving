package main

import "fmt"

func main12() {
	arr := [3][3]string{{"C#", "java", "golang"},
		{"C", "CPP", "JavaScript"},
		{"Hi", "Hello"}}
	//fmt.Println("Elements of Array 1")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(arr[i][j], " | ")
		}
		fmt.Println()
	}

}

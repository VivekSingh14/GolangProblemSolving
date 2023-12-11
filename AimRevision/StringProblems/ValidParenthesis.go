package main

import "fmt"

func main8() {
	str := "[()]{}{[()()]()}"
	fmt.Println(ValidParentheses(str))
}

func ValidParentheses(str string) bool {
	r1 := []rune(str)

	var stack []rune

	for i := 0; i < len(r1); i++ {
		if r1[i] == '(' || r1[i] == '{' || r1[i] == '[' {
			stack = append(stack, r1[i])
		} else if r1[i] == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else if r1[i] == '}' && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
		} else if r1[i] == ']' && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

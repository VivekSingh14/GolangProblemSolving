package main

import (
	"fmt"
)

func main() {
	day := "Mon"

	switch {
	case day == "Mon":
		fmt.Println("Monday")
		fallthrough
	case day == "Tue":
		fmt.Println("Tuesday")
		fallthrough
	case day == "Wed":
		fmt.Println("Wednesday")

	}
}

package main

import "fmt"

func main33() {
	var temp interface{} = "Vivek"

	switch t := temp.(type) {
	case int64:
		fmt.Println("integer: ", t)

	case string:
		fmt.Println("string: ", t)

	default:
		fmt.Println("nothing")
	}

}

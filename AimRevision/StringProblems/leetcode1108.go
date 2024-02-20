package main

import "fmt"

func main16() {
	address := "1.1.1.1"
	fmt.Println(defangIPaddr(address))

}

func defangIPaddr(address string) string {
	res := ""
	for i := 0; i < len(address); i++ {
		if string(address[i]) == "." {
			res += "["
			res += "."
			res += "]"
		} else {
			res += string(address[i])
		}
	}
	return res
}

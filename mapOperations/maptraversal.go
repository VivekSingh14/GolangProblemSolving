package main

import "fmt"

func main() {

	map1 := make(map[string]int)

	map1["ABC"] = 1
	map1["DEF"] = 2
	map1["GHI"] = 3
	map1["JKL"] = 4
	map1["MNO"] = 5
	map1["PQR"] = 6

	//map traverse
	for e, v := range map1 {
		fmt.Println(e, " : ", v)
	}

	//if key exists
	val, ok := map1["XYZ"]

	if ok {
		fmt.Println(val, "& and yes it exists.")
	} else {
		fmt.Println("it doesn't exist.")
	}

	//other way of check if key exists
	if val1, okay := map1["ABC"]; okay {
		fmt.Println(map1["ABC"], " : ", val1)
	}
}

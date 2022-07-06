package main

import "fmt"

func main() {
	//Nil Map with key as integer and value type as string
	//var mymap map[int]string

	//Empty map
	var mymap map[int]string = map[int]string{}

	//putting key value pairs into map
	mymap[0] = "Vivek"
	mymap[1] = "Singh"
	mymap[2] = "Jaipur"
	mymap[0] = "Pune"

	fmt.Println(mymap)

	var otherMap = make(map[int]string)

	fmt.Println(otherMap)

	otherMap[1] = "Hey"
	otherMap[2] = "This is just testing"

	fmt.Println(otherMap)

}

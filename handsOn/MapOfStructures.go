package main

import "fmt"

type Person1 struct {
	name string
	age  int
}

type Vehicle struct {
	vehicleName string
	age         int
}

func main() {
	map1 := make(map[string]interface{})

	person := Person1{name: "Vivek", age: 27}

	vehicle1 := Vehicle{vehicleName: "Alto", age: 3}

	hi := 3
	if hi == 3 {
		map1["person"] = person
	} else {
		map1["vehicle"] = vehicle1
	}

	fmt.Println(map1)
}

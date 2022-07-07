package main

import "fmt"

type person struct {
	Id      int
	name    string
	age     int
	phoneNo int
	gender  string
}

func main25() {
	var person1 []person = []person{
		{100, "Vivek", 27, 9782606429, "Male"},
		{101, "Alisha", 25, 9680437722, "Female"},
		{102, "Shubham", 23, 9782123456, "Male"},
		{103, "Harish", 27, 9782000002, "Male"},
	}

	var map1 map[int]person = map[int]person{}

	for i := 0; i < len(person1); i++ {
		map1[person1[i].Id] = person1[i]
	}

	fmt.Println("***********************")

	fmt.Println(map1)

	person1 = append(person1, person{106, "Shahnawaz", 27, 9782000005, "Male"})
	map1[person1[len(person1)-1].Id] = person1[len(person1)-1]

	fmt.Println("***********************")

	fmt.Println(map1)

}

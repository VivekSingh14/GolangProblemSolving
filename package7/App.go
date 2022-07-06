package main

import "fmt"

type Person struct {
	Id   string
	Name string
	City string
}
type PersonImpl struct {
	id    string
	test1 string
	test2 string
}

func main() {
	person1 := Person{
		Id:   "1",
		Name: "Vivek",
		City: "Jaipur",
	}
	personMap := map[Person]string{
		person1: person1.City,
	}
	fmt.Println(personMap)

}

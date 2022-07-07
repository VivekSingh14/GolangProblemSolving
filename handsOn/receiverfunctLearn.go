package main

import "fmt"

type employee struct {
	id         int
	name, city string
	standard   int
	age        int
}

func (e employee) display() {
	fmt.Println("Id: ", e.id)
	fmt.Println("Name: ", e.name)
	fmt.Println("City: ", e.city)
	fmt.Println("Standard: ", e.standard)
	fmt.Println("Age: ", e.age)
}

func main18() {
	employee1 := employee{
		id:       1,
		name:     "Vivek",
		city:     "Jaipur",
		standard: 8,
		age:      19,
	}
	employee1.display()

}

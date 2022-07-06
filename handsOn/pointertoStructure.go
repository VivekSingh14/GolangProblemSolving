package main

import "fmt"

type Employee struct {
	firstname, lastname string
	age, salary         int
}

func NewEmployee() *Employee {
	return &Employee{firstname: "Vivek",
		lastname: "Singh",
		age:      27,
		salary:   8000,
	}
}
func main() {
	var emp Employee = Employee{"Vivek", "Singh", 27, 9000}
	fmt.Println("Employee's first name is : ", emp.firstname, "and his/her salary is: ", emp.salary)

	emp2 := NewEmployee()
	fmt.Println(emp2.firstname)
}

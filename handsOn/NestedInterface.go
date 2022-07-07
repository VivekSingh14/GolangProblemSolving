package main

import "fmt"

//Parent Interface
type Person2 interface {
	Employee1
	Employee2
}

type Employee1 interface {
	CreateSalary(employeeName string)
}

type Employee2 interface {
	FixedSalary(employeeId int)
}

type Employee1Impl struct {
	basic int
}

type Employee2impl struct {
	minwage int
}

//receiver functions
func (e Employee1Impl) CreateSalary(name string) {
	fmt.Println(name, " : ", e.basic*3)
}

func (e1 Employee2impl) FixedSalary(Id int) {
	fmt.Println(Id, " : ", e1.minwage*2)
}

func main13() {
	employee1impl := Employee1Impl{basic: 3000}
	employee1impl.CreateSalary("Vivek")

	employee2impl := Employee2impl{minwage: 3000}
	employee2impl.FixedSalary(201)
}

package p1

import "fmt"

type PersonRepository interface {
	GetUser(int) int
}

func (n NameConfig) GetUser(age int) int {
	fmt.Println(n.Name, " ", age)
	return age
}

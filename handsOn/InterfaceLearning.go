package main

import "fmt"

type standard interface {
	display() int
}

type person1 struct {
	name string
	age  int
}

func (t person1) display() int {
	fmt.Println(t.name, " : ", t.age)
	return 1
}
func main() {
	var t standard
	t = person1{"Vivek", 27}
	fmt.Println(t.display())
}

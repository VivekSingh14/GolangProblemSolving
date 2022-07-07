package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewBlogPost() *User {
	return &User{
		"Vivek", 23,
	}
}

func (u *User) GetDetails() {
	fmt.Println(u.Name, " ", u.Age)
}

func main26() {
	//var u1 *User
	u1 := NewBlogPost()
	fmt.Println(u1.Name, " ", u1.Age)
	u1.GetDetails()

	var a1 = 10
	var b1 *int = &a1
	fmt.Println(*b1, " ", b1)

	fmt.Println(*u1, " ")
}

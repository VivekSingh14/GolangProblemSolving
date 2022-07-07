package main

import "fmt"

func main15() {

	var a1 int = 3
	var a2 *int = &a1

	fmt.Println(a2, " ", *a2, " ", a1, " ", &a1)

	//pass by value and tried to override the value
	sum1(a1)
	fmt.Println("*******************")
	fmt.Println(a2, " ", *a2, " ", a1, " ", &a1)

	//pass by reference
	sum2(&a1)
	fmt.Println("*******************")
	fmt.Println(a2, " ", *a2, " ", a1, " ", &a1)

	var v = 7

	//other way to declare the pointer and assigning the address of variable.
	var p = &v

	fmt.Println("*******************")
	fmt.Println(p, " ", *p, " ", v)

	v1 := 4

	//third way to declare the pointer and assigning the address of variable.
	p1 := &v1

	fmt.Println("*******************")
	fmt.Println(p1, " ", *p1, " ", v1)

}

func sum1(a int) {
	a = 8
}

func sum2(a *int) {
	*a = 8
}

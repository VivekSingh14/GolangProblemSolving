package main

import "fmt"

type Circle struct {
	x int
	y int
	z int
}
type Square struct {
	a int
}

func circleArea(c Circle) int {
	c.x = 5
	return 3 * c.x * c.z
}

func circleArea2(c *Circle) int {
	c.x = 6
	return 3 * c.x * c.z
}

func (c *Circle) circleArea3() int {
	c.x = 6
	return 3 * c.x * c.z
}

func (r *Square) circleArea3() int {
	r.a = 7
	return r.a * r.a
}
func main() {
	var c = Circle{10, 20, 30}

	//Passing the var c as a value which will be copied into the method.
	fmt.Println(circleArea(c), "  ", c.x)
	fmt.Println("------------------")
	//Passing the address of c but the type of c is Circle
	fmt.Println(circleArea2(&c), "  ", c.x)

	//new is used to return the pointer to type you can check the by hovering on it.
	var c1 = new(Circle)
	c1.x = 10
	c1.y = 20
	c1.z = 30
	fmt.Println("------------------")

	fmt.Println(circleArea2(c1), "  ", c1.x)
	//Now Using receiver, we can pass any type of variable whether it is value type or pointer type as a receiver to method. go automatically converts it.
	//var c2 = Circle{10, 20, 30}
	var c2 = new(Circle)
	c2.x = 10
	c2.y = 20
	c2.z = 30
	fmt.Println("------------------")

	fmt.Println(c2.circleArea3(), "  ", c2.x)

	var r2 = new(Square)
	r2.a = 2
	fmt.Println("------------------")

	fmt.Println(r2.circleArea3(), "  ", r2.a)

}

package main

import "fmt"

type Triangle struct {
	base, height float32
}

type Square1 struct {
	length float32
}

type Rectangle struct {
	length, breadth float32
}

type Area interface {
	Area() float32
}

func (square Square1) Area() float32 {
	return square.length * square.length
}

func (rectangle Rectangle) Area() float32 {
	return rectangle.length * rectangle.breadth
}

func (triangle Triangle) Area() float32 {
	return 0.5 * triangle.base * triangle.height
}
func main() {
	triangleObject := Triangle{
		base:   20,
		height: 10,
	}
	squareObject := Square1{
		length: 10,
	}
	rectangleObject := Rectangle{
		length:  5,
		breadth: 10,
	}

	var area Area
	fmt.Printf("area type = %T\n", area)
	area = triangleObject
	res := area.Area()
	fmt.Println("------------------")
	fmt.Printf("area type = %T\n", area)
	fmt.Println(" area of triangle: ", res)

	fmt.Println("------------------")
	area = squareObject
	res = area.Area()
	fmt.Printf("area type = %T\n", area)
	fmt.Println(" area of square: ", res)

	fmt.Println("------------------")
	area = rectangleObject
	res = area.Area()
	fmt.Printf("area type = %T\n", area)
	fmt.Println(" area of rectangle: ", res)

}

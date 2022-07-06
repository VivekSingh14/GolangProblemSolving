package main

import (
	"demo1/package4/p2"
	"fmt"
)

type Tank interface {
	Area() float64
}

type Radius struct {
	radi float64
}

func (r Radius) Area() float64 {
	return 3.14 * r.radi * r.radi
}

func main() {
	var t Tank

	t = Radius{}

	fmt.Println(" ", t.Area(), " ")

	temp := &p2.RepoService{}
	temp.GetNamesByService()

}

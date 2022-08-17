package main

import "fmt"

type Impl interface {
	Display()
}

type Car struct {
	ModelName string
	CarType   string
	Average   int
}

type Vehicle struct {
	RTONO       string
	LaunchYear  int
	VehicleType Car
}

func (v Vehicle) Display() {
	fmt.Println(v.RTONO, " : ", v.LaunchYear, " : ", v.VehicleType.CarType, " : ", v.VehicleType.ModelName, " : ", v.VehicleType.Average, " : ")
}

func main() {
	c1 := Car{"Kushaq", "SUV", 14}
	v1 := Vehicle{"RJ14", 2020, c1}
	show(v1)
}

func show(i Impl) {
	i.Display()
}

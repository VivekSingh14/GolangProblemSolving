package main

import "fmt"

type purchase interface {
	sell()
}
type display interface {
	show()
}
type salesman interface {
	purchase
	display
}

type game struct {
	name, price    string
	gameCollection []string
}

func (g game) sell() {
	fmt.Println("-------------------------------")
	fmt.Println("Name: ", g.name)
	fmt.Println("Price", g.price)
	fmt.Println("-------------------------------")
	for _, s := range g.gameCollection {
		fmt.Print(s, " : ")
	}
	fmt.Println()
}

func (g game) show() {
	fmt.Println("-------------------------------")
}

func shop(s salesman) {
	fmt.Println("Accessing from shop.")
	s.show()
	s.sell()
}

func main() {
	collection := []string{
		"Clash of Clans", "Contra", "IGI", "ViceCity",
	}
	//var game1 salesman
	game1 := game{"Vivek", "35.90", collection}
	//game1.sell()
	shop(game1)
}

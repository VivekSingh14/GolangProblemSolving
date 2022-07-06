package main

import "fmt"

type details struct {
	genre       string
	genreRating string
	reviews     string
}

type game struct {
	name  string
	price float64
	details
}

func (d details) ShowDetails() {
	fmt.Println("Genre ", d.genre, " Genre Rating ", d.genreRating, " reviews ", d.reviews)
}

func (g game) Show() {
	fmt.Println("Name ", g.name, " Price ", g.price)
	g.ShowDetails()
}

func main() {
	action := details{
		genre:       "war",
		genreRating: "18+ ",
		reviews:     "Average",
	}
	action.ShowDetails()
	action2 := details{
		genre:       "crime",
		genreRating: "12+ ",
		reviews:     "Good",
	}
	action2.ShowDetails()

	game1 := game{
		"Clash of Clans", 35.50, action2,
	}
	game1.Show()
}

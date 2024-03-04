package main

import "fmt"

func main12() {
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println(highestAltitude(gain))

}

func highestAltitude(gain []int) int {
	res := 0
	currentres := 0
	for i := 1; i < len(gain); i++ {
		currentres += gain[i]
		if currentres > res {
			res = currentres
		}
	}

	return res
}

func largestAltitude(gain []int) int {
	maxAltitude := 0
	currAltitude := 0

	for _, altitudeGain := range gain {
		currAltitude += altitudeGain
		if currAltitude > maxAltitude {
			maxAltitude = currAltitude
		}
	}

	return maxAltitude
}

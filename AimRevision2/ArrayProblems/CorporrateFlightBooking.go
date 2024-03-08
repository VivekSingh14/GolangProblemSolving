package main

import "fmt"

func main16() {
	bookings := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	n := 5
	//bookings := [][]int{{1, 2, 10}, {2, 2, 15}}
	//n := 2

	//fmt.Println(corpFlightBookings(bookings, n))
	fmt.Println(corpFlightBookingsOpt(bookings, n))
}

func corpFlightBookings(bookings [][]int, n int) []int {
	var flight [][]int
	for i := 0; i < len(bookings); i++ {
		temp := make([]int, n+1)
		for j := bookings[i][0]; j <= bookings[i][1]; j++ {
			temp[j] = bookings[i][2]
		}
		flight = append(flight, temp)
	}
	var res []int

	for i := 0; i < len(flight[0]); i++ {
		count := 0
		for j := 0; j < len(flight); j++ {
			count += flight[j][i]
		}
		res = append(res, count)
	}
	return res[1:]
}

func corpFlightBookingsOpt(bookings [][]int, n int) []int {

	ans := make([]int, n+1)
	for _, book := range bookings {
		ans[book[0]-1] += book[2]
		ans[book[1]] -= book[2]
	}
	fmt.Println(ans)

	for i := 1; i < n; i++ {
		ans[i] += ans[i-1]
	}
	return ans[:n]
}

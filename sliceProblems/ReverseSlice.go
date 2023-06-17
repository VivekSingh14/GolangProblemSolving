package main

import (
	"fmt"
)

func main5() {

	s := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(s, " : ", len(s), " : ", cap(s))

	//remove the middle of slice

	middle := len(s) / 2
	var result []int
	for i := 1; i <= middle; i++ {
		temp1 := s[0:i]
		result = append(result, temp1[i-1])
	}

	for i := middle + 1; i < len(s); i++ {
		temp1 := s[i : i+1]
		result = append(result, temp1[0])
	}
	fmt.Println(result)

	fmt.Println("-------------------------")
	//other method to reverse the slice

	t := []int{1, 2, 3, 4, 5, 6, 7}
	middle1 := len(t) / 2
	var result2 []int

	copy(result2, t[:middle1])
	copy(result2, t[middle1+1:len(s)])

	fmt.Println(result2)

	fmt.Println("------------------------")

	//reverse methods are over here

	//temp1 := s[middle+1 : len(s)]
	//s = s[:len(s)-1]

	//fmt.Println(s, " : ", len(s), " : ", cap(s))

	// a := s[:3]
	// fmt.Println(a, " : ", len(a), " : ", cap(a))
	// b := s[3:]
	// fmt.Println(b, " : ", len(b), " : ", cap(b))
	// var data int
	// for data = range s {
	// 	temp := s[len(s)-1 : len(s)]
	// 	s = s[:len(s)-1]
	// 	fmt.Println(temp)
	// }
	// log.Println(data)
}

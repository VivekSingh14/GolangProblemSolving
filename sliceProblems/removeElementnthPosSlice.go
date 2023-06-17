package main

import "fmt"

func main() {
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

	//result2 = remove(t, 3)

	result2 = append(t[:middle1], t[middle1+1:]...)

	fmt.Println(result2)

}

// func remove(slice []int, s int) []int {
// 	return append(slice[:s], slice[s+1:]...)
// }

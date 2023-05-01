package main

import "fmt"

// 2nd coding question
// ==========================
// There is a slice. do two subtasks:
// (1). Add value of 2nd position to 1st position;
//      Add value of 4th position to 3th position;
//      Add value of 6th position to 5th position;
//      ...
//      Add value of "2N" position to "2N-1" position;
// (2). Then remove 2N (even) position
// Please do NOT use EXTRA space to implement two subtasks.

// Example:
// (1). Original slice is: 1, 2, 3, 4, 5, 6, 7
// (2). After finish first subtask, the slice should be: 3, 2, 7, 4, 11, 6, 7
//      Because 3 (1st position) == 1 + 2; 7 (3rd position) == 3 + 4; 11 (5th position) == 5 + 6
//      NOTE: 7 (7th postion) is the last digit. So keep it same
// (3). After finish second subtask, the slice should be: 3, 7, 11, 7
// 3,2,7,4,11,6,7
//    0 1 2 3 4  5 6
//      1
// // 3 7 4 11 6 7
//    0 1 2  3 4 5
//      2
// // 3 7 11 6 7
//    0 1 2  3 4
//      3
// // 3 7 11 7
//    0 1  2 3

func main1() {
	var slice1 []int

	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)
	slice1 = append(slice1, 3)
	slice1 = append(slice1, 4)
	slice1 = append(slice1, 5)
	slice1 = append(slice1, 6)
	slice1 = append(slice1, 7)

	fmt.Println(slice1)

	for i := 0; i < len(slice1); i++ {
		if i%2 != 0 {
			slice1[i-1] = slice1[i] + slice1[i-1]
		}
	}

	fmt.Println(slice1)
	fmt.Println("*****************")

	for i := 1; i < len(slice1); i++ {

		slice1 = append(slice1[:i], slice1[i+1:]...)

	}
	fmt.Println(slice1)

	// fmt.Println("----------------------------")
	// // range method to traverse a slice
	// for temp, element := range slice1 {
	// 	fmt.Println(element, temp)
	// }
	// //
	// fmt.Println("----------------------------")
	// for _, element := range slice1 {
	// 	fmt.Println(element)
	// }

}

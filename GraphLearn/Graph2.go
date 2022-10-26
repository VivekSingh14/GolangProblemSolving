package main

import "fmt"

//the representation of graph using  adjacency matrics
// graph representation

//     8
// 1 ----> 2
//  /\    |
// | \    |
// |3 \6  | 9
// |   \  |
// |    \ |
// \|/ 7 \|/
// 4 ---- >3

// directed weighted Graph
// |V| = 4
// |E| = 5

func main() {
	//NoOfVertices := 5
	adjMat := [5][5]int{}

	// for i := range adjMat {
	// 	adjMat[i] = make([]int, NoOfVertices)
	// }
	// adding the comment
	adjMat[1][2] = 8
	adjMat[1][4] = 3
	adjMat[2][3] = 9
	adjMat[3][1] = 6
	adjMat[4][3] = 7

	for i := 0; i < len(adjMat); i++ {
		for j := 0; j < len(adjMat[i]); j++ {
			fmt.Print(adjMat[i][j], " ")
		}
		fmt.Println()
	}

}

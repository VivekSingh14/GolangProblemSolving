package main

import "fmt"

//the representation of graph using  adjacency matrics
// graph representation
// 1 ---- 2
// |\     |
// | \    |
// |  \   |
// |   \  |
// |    \ |
// |     \|
// 4 ---- 3
// \      /
//  \    /
//   \  /
//     5

// Undirected Graph
// |V| = 5
// |E| = 7

func main1() {
	//NoOfVertices := 5
	adjMat := [6][6]int{}

	// for i := range adjMat {
	// 	adjMat[i] = make([]int, NoOfVertices)
	// }

	adjMat[1][2] = 1
	adjMat[2][1] = 1
	adjMat[1][4] = 1
	adjMat[4][1] = 1
	adjMat[2][3] = 1
	adjMat[3][2] = 1
	adjMat[1][3] = 1
	adjMat[3][1] = 1
	adjMat[4][3] = 1
	adjMat[3][4] = 1
	adjMat[4][5] = 1
	adjMat[5][4] = 1
	adjMat[3][5] = 1
	adjMat[5][3] = 1

	for i := 0; i < len(adjMat); i++ {
		for j := 0; j < len(adjMat[i]); j++ {
			fmt.Print(adjMat[i][j], " ")
		}
		fmt.Println()
	}

}

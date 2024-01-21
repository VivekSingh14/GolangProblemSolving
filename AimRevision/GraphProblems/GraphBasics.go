package main

import "fmt"

type Graph1 struct {
	vertices    int
	adjmatrices [][]int
}

func NewGraph(vertices int) *Graph1 {
	matrix := make([][]int, vertices+1)

	for i := 0; i <= vertices; i++ {
		matrix[i] = make([]int, vertices+1)
	}

	return &Graph1{
		vertices:    vertices,
		adjmatrices: matrix,
	}

}

func (g *Graph1) AddEdge(source, destination int) {
	g.adjmatrices[source][destination] = 1
	g.adjmatrices[destination][source] = 1
}

func (g *Graph1) DFSGraph(vertex int) {
	visited := make([]int, g.vertices)
	g.DFSUtil(vertex, visited)
}

func (g *Graph1) DFSUtil(vertex int, visited []int) {
	visited[vertex] = 1

	fmt.Println(vertex)

	for i := 0; i < g.vertices; i++ {
		if g.adjmatrices[vertex][i] == 1 && visited[i] == 0 {
			g.DFSUtil(i, visited)
		}
	}
}

func main() {

	g := NewGraph(5)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)

	for i := 1; i < len(g.adjmatrices); i++ {
		for j := 1; j < len(g.adjmatrices); j++ {
			fmt.Print(g.adjmatrices[i][j], " ")
		}
		fmt.Println()
	}

	g.DFSGraph(1)

}

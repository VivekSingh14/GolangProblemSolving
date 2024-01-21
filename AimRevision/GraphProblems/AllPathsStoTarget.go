package main

import "fmt"

func main1() {
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}

func allPathsSourceTarget(graph [][]int) [][]int {

	ans := [][]int{}

	q := [][]int{}
	q0 := []int{0}

	q = append(q, q0)

	fmt.Println(len(q))

	for len(q) > 0 {

		q0 = q[0]
		q = q[1:]

		lastIdx := q0[len(q0)-1]
		for _, v := range graph[lastIdx] {

			item := make([]int, len(q0))
			_ = copy(item, q0)
			item = append(item, v)

			if v == len(graph)-1 {
				ans = append(ans, item)
				continue
			}
			q = append(q, item)
		}
	}

	return ans
}

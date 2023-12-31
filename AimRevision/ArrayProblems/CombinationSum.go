package main

import "fmt"

func main29() {
	arr := []int{2, 4, 6, 8}
	sum := 8
	fmt.Println(combinationSum(arr, sum))
}

func combinationSum(candidates []int, target int) [][]int {
	out := make([][]int, 0)
	for i := 0; i < len(candidates); i++ {
		if target-candidates[i] < 0 {
			continue
		} else if target-candidates[i] == 0 {
			out = append(out, []int{candidates[i]})
		} else {
			tmp := combinationSum(candidates[i:], target-candidates[i])
			for j := 0; j < len(tmp); j++ {
				tmp[j] = append(tmp[j], candidates[i])
			}
			out = append(out, tmp...)
		}
	}
	return out
}

package main

import "fmt"

type data struct {
	str []string
}

func main9() {
	arr := []string{"cat", "dog", "tac", "god", "act"}
	groupAnagram(arr)
}

func groupAnagram(arr []string) {

	map1 := make(map[int]data)
	for i := 0; i < len(arr); i++ {
		r := []rune(arr[i])
		temp := 0
		for j := 0; j < len(r); j++ {
			temp += int(r[j])
		}
		if _, ok := map1[temp]; !ok {
			var data1 []string
			data1 = append(data1, string(r))
			map1[temp] = data{data1}
		} else {
			res := map1[temp]
			res.str = append(res.str, string(r))
			map1[temp] = res
		}
	}

	for _, value := range map1 {
		fmt.Println(value)
	}

}

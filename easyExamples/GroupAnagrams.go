package main

import (
	"fmt"
	"sort"
	"strings"
)

func main45() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)

	fmt.Println(res)
}

func groupAnagrams(strs []string) [][]string {

	var res [][]string

	for i := 0; i < len(strs); i++ {
		j := len(strs) - 1

		temp := make([]string, 1)
		for j > i {

			if IsAnagram(strs[i], strs[j]) {
				temp = append(temp, strs[j])
				fmt.Println(temp)
				j--
				return nil
			} else {
				j--
			}
		}
		temp = append(temp, strs[i])
		fmt.Println(temp)
		res = append(res, temp)
	}
	return nil
}

func IsAnagram(str1 string, str2 string) bool {

	res1 := strings.Split(str1, "")
	res2 := strings.Split(str2, "")

	sort.Slice(res1, func(i, j int) bool {
		return res1[i] < res1[j]
	})
	sort.Slice(res2, func(i, j int) bool {
		return res2[i] < res2[j]
	})
	if strings.Join(res1, "") == strings.Join(res2, "") {
		return true
	}

	return false
}

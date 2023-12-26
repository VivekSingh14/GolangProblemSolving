package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	str := []string{"you", "lint", "code", "love", "you"}
	res := Encode(str)
	//fmt.Println(res)

	fmt.Println(Decode(res))
}

func Encode(strs []string) string {
	ans := &bytes.Buffer{}
	for _, s := range strs {
		t := fmt.Sprintf("%04d", len(s))
		ans.WriteString(t)
		ans.WriteString(s)
		fmt.Println(ans)
	}
	return ans.String()
}

// Decodes a single string to a list of strings.
func Decode(strs string) []string {
	ans := []string{}
	i, n := 0, len(strs)
	for i < n {
		t := strs[i : i+4]
		i += 4
		size, _ := strconv.Atoi(t)
		ans = append(ans, strs[i:i+size])
		i += size
	}
	return ans
}

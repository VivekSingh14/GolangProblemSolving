package main

import "fmt"

func main() {
	str := "ABC"
	permute(str, 0, len(str)-1)

}

func permute(str string, l int, r int) {

	if l == r {
		fmt.Println(str)
	} else {
		for i := l; i <= r; i++ {
			str = swap(str, l, i)
			permute(str, l+1, r)
			str = swap(str, l, i)
		}
	}
}

func swap(str string, l int, i int) string {
	r := []rune(str)
	temp := r[l]
	r[l] = r[i]
	r[i] = temp
	return string(r)
}

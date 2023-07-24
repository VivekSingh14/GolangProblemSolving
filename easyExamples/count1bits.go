package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var n uint32 = 00000000000000000000000000001011
	fmt.Println(hammingWeight(n))
	fmt.Println(hammingWeight2(n))
}

func hammingWeight(num uint32) int {

	return bits.OnesCount32(num)

}

func hammingWeight2(num uint32) int {
	var out int

	for num > 0 {
		if num%2 == 1 {
			out++
		}

		num /= 2
	}

	return out
}

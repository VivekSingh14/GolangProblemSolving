package main

import "fmt"

func main29() {
	n := 3
	fmt.Println(countPrimes(n))
}

func countPrimes(n int) int {
	primes := make([]bool, n)
	for i := 0; i < len(primes); i++ {
		primes[i] = true
	}
	for i := 2; i*i < len(primes); i++ {
		if primes[i] {
			for j := i; j*i < len(primes); j++ {
				primes[i*j] = false
			}
		}
	}
	count := 0
	for i := 2; i < len(primes); i++ {
		if primes[i] {
			count++
		}
	}
	return count
}

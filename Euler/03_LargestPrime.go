package main

import (
	"fmt"
)

func main() {
	res := PrimeFactors(600851475143) // [71 839 1471 6857]

	// largest
	fmt.Println(res[len(res)-1]) // 6857
}

func PrimeFactors(n int) []int {
	primes := []int{}

	for n%2 == 0 {
		primes = append(primes, 2)
		n = n / 2
	}

	for i := 3; i <= n; i += 2 {
		for n%i == 0 {
			primes = append(primes, i)
			n = n / i
		}
	}

	if n > 2 {
		primes = append(primes, n)
	}

	return primes
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var tm, m1, m2 int

	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			p := i * j

			if tm > p {
				break
			}

			if isPalindrome(int64(p)) {
				tm = p
				m1 = i
				m2 = j
			}
		}
	}

	fmt.Println("Result:", tm)
	fmt.Println("Product of:", m1, "and", m2)
}

func isPalindrome(n int64) bool {
	ns := strconv.FormatInt(n, 10)
	bound := (len(ns) / 2) + 1

	for i := 0; i < bound; i++ {
		if ns[i] != ns[len(ns)-1-i] {
			return false
		}
	}

	return true
}

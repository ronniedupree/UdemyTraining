package main

import (
	"fmt"
)

// 2520 is the smallest number that can be divided by each of the numbers
// from 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by all
// of the numbers from 1 to 20?

func main() {
	r := 1
	for i := 2; i <= 20; i++ {
		r = LCM(r, i)
	}
	fmt.Println(r)
}

func LCM(x, y int) int {
	return x * y / GCD(x, y)
}

func GCD(x, y int) int {
	for y != 0 {
		t := y
		y = x % y
		x = t
	}

	return x
}

package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := big.NewInt(100)
	t := fibonacci(n)
	fmt.Println(t)
}

func fibonacci(n *big.Int) *big.Int {

	total := big.NewInt(0)
	m := big.NewInt(0)
	f2 := big.NewInt(0)
	f1 := big.NewInt(1)

	if n.Cmp(big.NewInt(1)) == 0 {
		return f2
	}

	if n.Cmp(big.NewInt(2)) == 0 {
		return f1
	}

	for i := 3; n.Cmp(big.NewInt(int64(i))) >= 0; i++ {
		next := big.NewInt(0)
		next.Add(f2, f1)
		m = m.Mod(next, big.NewInt(2))

		if m.Cmp(big.NewInt(0)) == 0 && next.Cmp(big.NewInt(4000000)) == -1 {
			total = total.Add(total, next)
		}

		f2 = f1
		f1 = next
	}

	return total
}

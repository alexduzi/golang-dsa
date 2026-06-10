package main

import (
	"cmp"
	"fmt"
	"slices"
)

// Dado um valor V e um conjunto de N moedas, ache o número mínimo de moedas
// que representa V. Considere que sempre é possível representar V com o conjunto, e
// que podemos tomar um número infinito de moedas.
func main() {
	fmt.Println(minimumCoins(37, 4, []int{25, 10, 5, 1}))
}

func minimumCoins(v, n int, coins []int) int {
	// O(n log n)
	// ordenar decrescente, para pegar as maiores moedas primeiro
	slices.SortFunc(coins, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	ans := 0

	// O(n)
	for i := 0; i < n; i++ {
		// retirando a moeda atual enquanto for possivel
		for (v - coins[i]) >= 0 {
			v -= coins[i]
			ans++
		}
	}

	return ans
}

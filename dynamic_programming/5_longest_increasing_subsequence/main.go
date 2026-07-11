package main

import "fmt"

func listLength(s []int) int {
	n := len(s)
	lis := make([]int, n)
	for i := 0; i < n; i++ {
		lis[i] = 1
	}

	ans := 1

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			// para ser crescente o valor anterior deve ser menor que o atual
			if s[j] < s[i] {
				lis[i] = max(lis[i], 1+lis[j])
			}
		}
		ans = max(ans, lis[i])
	}

	return ans
}

// Maior Subsequência Crescente (LIS - Longest Increasing Subsequence)
// Dada uma sequência S com n valores inteiros, encontre o tamanho da maior
// subsequência de S que seja estritamente crescente.
// Subsequência: é obtida removendo quaisquer caracteres da sequência
// original, mantendo a ordem dos restantes.
// Qual o tamanho da maior subsequência crescente para a
// sequência abaixo?
// S = [3, 1, 8, 2, 5]
// Resposta: O tamanho é 3, pois a LIS é [1, 2, 5]
func main() {
	fmt.Println(listLength([]int{3, 1, 8, 2, 5}))
	fmt.Println(listLength([]int{10, 4, 5, 6, 5, 15, 2, 3}))
}

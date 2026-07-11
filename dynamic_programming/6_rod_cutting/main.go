package main

import "fmt"

// solução recursiva padrão
func rodCutting(n int, p []int) int {
	if n == 0 {
		return 0
	}

	ans := -1

	for i := 1; i <= n; i++ {
		if n-i >= 0 {
			ans = max(ans, rodCutting(n-i, p)+p[i-1])
		}
	}

	return ans
}

// solução usando PD top-down
func rodCutting2(n int, p, memo []int) int {
	if n == 0 {
		return 0
	}

	if memo[n] != -1 {
		return memo[n]
	}

	ans := -1

	for i := 1; i <= n; i++ {
		if n-i >= 0 {
			ans = max(ans, rodCutting2(n-i, p, memo)+p[i-1])
		}
	}

	memo[n] = ans

	return memo[n]
}

func rodCuttingMemo(n int, p []int) int {
	memo := make([]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = -1
	}

	return rodCutting2(n, p, memo)
}

// solução bottom-up
func rodCutting3(n int, p []int) int {
	memo := make([]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = -1
	}

	memo[0] = 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i-j >= 0 {
				memo[i] = max(memo[i], memo[i-j]+p[j-1])
			}
		}
	}

	return memo[n]
}

// Problema do Corte de Hastes (Rod Cutting)
// Deseja cortar um bastão de n metros em pedaços de comprimentos inteiros. É dada uma tabela de
// preços que indica, para cada comprimento de i metros do bastão, seu valor p(i). Qual a soma
// máxima em dinheiro possível de se obter cortando o bastão de n metros?
func main() {
	fmt.Println(rodCutting(10, []int{3, 4, 8, 10, 10, 11, 23, 23, 24, 25}))
	fmt.Println(rodCuttingMemo(10, []int{3, 4, 8, 10, 10, 11, 23, 23, 24, 25}))
	fmt.Println(rodCutting3(10, []int{3, 4, 8, 10, 10, 11, 23, 23, 24, 25}))
}

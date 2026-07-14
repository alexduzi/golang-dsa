package main

import "fmt"

func maxProfit(n int, values []int) int {
	memo := make([]int, n+1)
	for i := 1; i < len(memo); i++ {
		memo[i] = -1
	}
	return maxProfitUtil(n, values, memo)
}

func maxProfitUtil(n int, values, memo []int) int {
	if n == 0 {
		return 0
	}

	if memo[n] != -1 {
		return memo[n]
	}

	for i := 1; i <= n; i++ {
		if n-1 >= 0 {
			memo[n] = max(memo[n], maxProfitUtil(n-i, values, memo)+values[i-1])
		}
	}

	return memo[n]
}

// Problema "precious_stones"
// Imagine que você é um lapidário e possui um bloco bruto de uma pedra preciosa de N gramas.
// Para lapidar a pedra em gemas, você pode quebrar esse bloco em tamanhos inteiros de gramas.
// Cada gema tem um valor associado que depende apenas da sua massa, e sabemos de antemão qual o valor values[i] de uma gema de massa 1, 2, …, n.

// Por exemplo, suponha que o bloco bruto tenha 4 gramas e values = [2, 5, 7, 9], isso significa que:
// uma gema de 1 grama tem valor 2
// uma gema de 2 gramas tem valor 5
// uma gema de 3 gramas tem valor 7
// uma gema de 4 gramas tem valor 9
// Nessa situação, para maximizar o lucro, a melhor decisão é lapidar duas gemas de 2 gramas, gerando valor final de 5 + 5 = 10.
// Sabendo disso, dado um valor N de gramas do bloco bruto e o valor das gemas de acordo com a quantidade de gramas,
// determine o lucro máximo que pode ser obtido, dividindo o bloco bruto de forma ótima.

func main() {
	fmt.Println(maxProfit(4, []int{2, 5, 7, 9}))
	fmt.Println(maxProfit(8, []int{1, 5, 8, 9, 10, 17, 17, 20}))
	fmt.Println(maxProfit(8, []int{3, 5, 8, 9, 10, 17, 17, 20}))
}

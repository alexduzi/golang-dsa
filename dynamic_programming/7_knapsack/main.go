package main

import "fmt"

// solução recursiva padrão
func knapsack(n, w int, weights, values []int) int {
	if n == 0 || w == 0 {
		return 0
	}

	if weights[n-1] > w {
		return knapsack(n-1, w, weights, values)
	} else {
		includeItem := values[n-1] + knapsack(n-1, w-weights[n-1], weights, values)
		excludeItem := knapsack(n-1, w, weights, values)
		return max(includeItem, excludeItem)
	}
}

// solução com memoização top-down
func knapsack2(n, w int, weights, values []int, memo [][]int) int {
	if n == 0 || w == 0 {
		return 0
	}

	if memo[n][w] != -1 {
		return memo[n][w]
	}

	if weights[n-1] > w {
		return knapsack2(n-1, w, weights, values, memo)
	} else {
		includeItem := values[n-1] + knapsack2(n-1, w-weights[n-1], weights, values, memo)
		excludeItem := knapsack2(n-1, w, weights, values, memo)

		memo[n][w] = max(includeItem, excludeItem)
		return memo[n][w]
	}
}

func knapsack2Helper(n, w int, weights, values []int) int {
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, w+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return knapsack2(n, w, weights, values, memo)
}

// solução com tabulação bottom-up
func knapsack3(n, w int, weights, values []int) int {
	grid := make([][]int, n+1)
	for i := range grid {
		grid[i] = make([]int, w+1)
		for j := range grid[i] {
			grid[i][j] = 0
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if weights[i-1] <= j {
				grid[i][j] = max(grid[i-1][j], values[i-1]+grid[i-1][j-weights[i-1]])
			} else {
				grid[i][j] = grid[i-1][j]
			}
		}
	}

	return grid[n][w]
}

// Problema da Mochila 0/1
// O Problema da Mochila 0-1 é um clássico desafio de otimização combinatória onde, dado um conjunto de itens
// (cada um com peso e valor fixos) e uma mochila com capacidade máxima, o objetivo é maximizar o valor total dos itens sem exceder o limite de peso.
// A restrição "0-1" significa que cada item deve ser pego inteiro ou deixado para trás (não pode ser fracionado)
func main() {
	// fmt.Println(knapsack(3, 5, []int{1, 2, 3}, []int{6, 10, 12}))
	// fmt.Println(knapsack2Helper(3, 5, []int{1, 2, 3}, []int{6, 10, 12}))
	fmt.Println(knapsack3(3, 5, []int{1, 2, 3}, []int{6, 10, 12}))
}

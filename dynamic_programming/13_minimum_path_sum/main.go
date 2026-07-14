package main

import "fmt"

func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	// inicializando a tabulação
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, m)
	}

	memo[0][0] = grid[0][0]

	// casos base para coluna
	for i := 1; i < n; i++ {
		memo[i][0] = memo[i-1][0] + grid[i][0]
	}

	// casos base para linha
	for j := 1; j < m; j++ {
		memo[0][j] = memo[0][j-1] + grid[0][j]
	}

	// aplicando tabulação
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			memo[i][j] = min(memo[i-1][j], memo[i][j-1]) + grid[i][j]
		}
	}

	return memo[n-1][m-1]
}

// Problema "minimum_path_sum" (Adaptado de Leetcode 64)
// Empresas: Goldman Sachs, Google, Facebook, Amazon
// Dado um grid m x n preenchido com números não negativos,
// encontre um caminho da parte superior esquerda para a parte inferior
// direita que minimize a soma de todos os números ao longo desse caminho.

// Observação: Você só pode se mover para baixo ou para a direita.

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))

	fmt.Println(minPathSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}

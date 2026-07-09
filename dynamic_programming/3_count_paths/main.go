package main

import "fmt"

func init() {
	memo = make([][]int, 4)
	for i := 0; i < 4; i++ {
		memo[i] = make([]int, 4)
	}
}

// contando caminhos únicos em um grid de maneira recursiva
// solução recursiva básica, sem aplicar o conceito de PD
func countPaths(i, j int) int {
	if i == 0 || j == 0 {
		return 1
	}

	return countPaths(i-1, j) + countPaths(i, j-1)
}

var memo [][]int

// código otimizado utilizando abordagem PD
// com memoização (abordagem topdown)
func countPaths2(i, j int) int {
	if memo[i][j] != 0 {
		return memo[i][j]
	}
	if i == 0 || j == 0 {
		return 1
	}

	memo[i][j] = countPaths(i-1, j) + countPaths(i, j-1)

	return memo[i][j]
}

// código otimizado utilizando abordagem PD
// com tabulação (abordagem bottomup)
func countPaths3(destI, destJ int) int {
	grid := make([][]int, 4)
	for i := 0; i < 4; i++ {
		grid[i] = make([]int, 4)
	}

	// caso base: primeira linha e primeira coluna têm apenas 1 caminho
	for i := 0; i < len(grid); i++ {
		grid[i][0] = 1
	}
	for j := 0; j < len(grid[0]); j++ {
		grid[0][j] = 1
	}

	// preenche o restante da tabela a partir dos casos base
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[destI][destJ]
}

func main() {
	fmt.Println(countPaths3(3, 3))
}

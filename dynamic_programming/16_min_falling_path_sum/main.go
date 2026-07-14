package main

import (
	"fmt"
	"math"
)

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
	}

	// casos base
	for j := 0; j < n; j++ {
		memo[0][j] = matrix[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			pathA, pathB, pathC := math.MaxInt, math.MaxInt, math.MaxInt

			if isValid(i-1, j-1, n) {
				pathA = matrix[i][j] + memo[i-1][j-1]
			}
			if isValid(i-1, j, n) {
				pathB = matrix[i][j] + memo[i-1][j]
			}
			if isValid(i-1, j+1, n) {
				pathC = matrix[i][j] + memo[i-1][j+1]
			}

			memo[i][j] = min(pathA, pathB, pathC)
		}
	}

	ans := math.MaxInt
	for j := 0; j < n; j++ {
		ans = min(ans, memo[n-1][j])
	}

	return ans
}

func isValid(i, j, n int) bool {
	return i >= 0 && i < n && j >= 0 && j < n
}

// Problema "min_falling_path_sum" (Adaptado de Leetcode 931)
// Empresas: Amazon, Apple, Google, Bloomberg, Microsoft, Facebook
// Dado uma matriz n x n de inteiros chamada matrix, retorne a soma mínima de qualquer caminho descendente através da matriz.
// Um caminho descendente começa em qualquer elemento da primeira linha e escolhe o elemento
// na próxima linha que está diretamente abaixo ou diagonalmente à esquerda/direita.
// Especificamente, o próximo elemento a partir da posição
// (linha, coluna) será (linha + 1, coluna - 1), (linha + 1, coluna) ou (linha + 1, coluna + 1).
func main() {
	fmt.Println(minFallingPathSum([][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}))

	fmt.Println(minFallingPathSum([][]int{
		{-19, 57},
		{-40, -5},
	}))
}

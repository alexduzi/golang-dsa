package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	grid := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		// amount+1 nunca é uma resposta válida, então serve como "infinito" sem risco de overflow
		grid[i] = amount + 1
	}

	// caso base
	grid[0] = 0

	subProblem := 0

	// preenchendo a tabela de pd
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			subProblem = i - coin
			if subProblem >= 0 {
				grid[i] = min(grid[i], grid[subProblem]+1)
			}
		}
	}

	if grid[amount] > amount {
		return -1
	}

	return grid[amount]
}

// Problema "coins" (Adaptado de Leetcode 322)
// Empresas: Amazon, TikTok, Google, Apple, Microsoft, Facebook
// Dado um array de inteiros coins representando moedas de diferentes denominações e um inteiro amount
// representando um valor total de dinheiro, retorne o menor número possível de moedas necessário para formar esse valor.
// Se não for possível formar o valor com nenhuma combinação das moedas fornecidas, retorne -1.
// Você pode assumir que há um número infinito de cada tipo de moeda disponível.
func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1, 2, 3, 7, 11}, 10000))
}

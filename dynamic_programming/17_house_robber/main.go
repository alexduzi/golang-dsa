package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	memo := make([]int, len(nums))
	for i := 0; i < n; i++ {
		memo[i] = -1
	}

	var robUtil func([]int, int) int

	robUtil = func(i1 []int, i2 int) int {
		if i2 < 0 {
			return 0
		}

		if memo[i2] != -1 {
			return memo[i2]
		}

		result := max(robUtil(i1, i2-1), robUtil(i1, i2-2)+i1[i2])
		memo[i2] = result
		return memo[i2]
	}

	return robUtil(nums, n-1)
}

// Problema "house_robber" (Adaptado de Leetcode 198)
// Empresas: Amazon, Apple, Microsoft, Google, Facebook
// Você é um ladrão profissional planejando roubar casas ao longo de uma rua.
// Cada casa tem uma certa quantia de dinheiro guardada.
// A única restrição que impede você de roubar cada uma delas é que casas adjacentes possuem sistemas de segurança conectados,
// que automaticamente alertarão a polícia se duas casas adjacentes forem invadidas na mesma noite.
// Dado um array de inteiros nums representando a quantidade de dinheiro em cada casa,
// retorne o valor máximo de dinheiro que você pode roubar esta noite sem alertar a polícia.

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

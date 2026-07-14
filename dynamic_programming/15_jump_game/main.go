package main

import "fmt"

func canJump(nums []int) bool {
	// inicializando o memo
	memo := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = -1
	}
	memo[len(nums)-1] = 1

	var canJumpFromPosition func(int) bool

	canJumpFromPosition = func(position int) bool {
		if memo[position] != -1 {
			return memo[position] == 1
		}

		furtherJump := min(position+nums[position], len(nums)-1)

		for i := position + 1; i <= furtherJump; i++ {
			if canJumpFromPosition(i) {
				memo[position] = 1
				return true
			}
		}

		memo[position] = 0
		return false
	}

	return canJumpFromPosition(0)
}

// Problema "jump_game" (Adaptado de Leetcode 55)
// Empresas: Amazon, Apple, Google, Bloomberg, Microsoft, Facebook
// Você recebe um array de inteiros chamado nums.
// Inicialmente, você está posicionado no primeiro índice do array,
// e cada elemento no array representa o comprimento máximo de salto possível naquela posição.
// Retorne true se você consegue chegar no último índice do array, ou false caso contrário.

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

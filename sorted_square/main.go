package main

import (
	"fmt"
	"sort"
)

// Dado um array de números inteiros nums ordenado em ordem crescente,
// retorne um array com os quadrados de cada número, também ordenado de forma crescente.
// entrada: nums = [-4,-1,0,3,10]
// saída: [0,1,9,16,100]
// Explicação: Após elevar ao quadrado, temos como resultado o array [16, 1, 0, 9, 100].
// Em seguida, após ordenar os valores do array, temos [0, 1, 9, 16, 100].
func main() {
	fmt.Printf("%v\n", sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Printf("%v\n", sortedSquares([]int{-7, -3, 2, 3, 11}))
}

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	sort.Ints(nums)
	return nums
}

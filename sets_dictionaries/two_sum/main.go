package main

import "fmt"

// Problema "two-sum" (ref: Leetcode two-sum)
// Dado um array de inteiros nums e um inteiro target, retorne os índices dos dois números de tal forma que eles somem target.
// Você pode assumir que cada entrada terá exatamente uma solução,
// e você não pode usar o mesmo elemento duas vezes. Você pode retornar a resposta em qualquer ordem.
// este desafio também dá para resolver com dois loops (brute force), porém a solução fica O(n^2)
func main() {
	fmt.Println(twoSum([]int{8, 2, 7, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 5, 7, 2, 4, 8, 1, 6}, 15))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for idx, val := range nums {
		if mIdx, ok := m[val]; ok {
			return []int{mIdx, idx}
		}
		m[target-val] = idx
	}

	return []int{}
}

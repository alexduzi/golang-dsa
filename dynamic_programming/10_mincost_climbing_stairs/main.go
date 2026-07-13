package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	return minCostClimbingStairsUtil(len(cost), cost)
}

func minCostClimbingStairsUtil(n int, cost []int) int {
	if n == 0 || n == 1 {
		return 0
	}

	return min(minCostClimbingStairsUtil(n-1, cost)+cost[n-1],
		minCostClimbingStairsUtil(n-2, cost)+cost[n-2])
}

// solução com memoization top-down
func minCostClimbingStairs2(cost []int) int {
	memo := make(map[int]int)
	return minCostClimbingStairsUtil2(len(cost), cost, memo)
}

func minCostClimbingStairsUtil2(n int, cost []int, memo map[int]int) int {
	if n == 0 || n == 1 {
		return 0
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = min(minCostClimbingStairsUtil2(n-1, cost, memo)+cost[n-1],
		minCostClimbingStairsUtil2(n-2, cost, memo)+cost[n-2])

	return memo[n]
}

// Problema "mincost_climbing_stairs" (Adaptado de Leetcode 746)
// Empresas: Google, Amazon, Microsoft

// É dada uma matriz inteira cost em que cost[i]
// é o custo do i-ésimo degrau em uma escada. Depois de pagar o custo, você pode subir um ou dois degraus.
// Você pode começar no degrau com índice 0 ou no degrau com índice 1.
// Retorne o custo mínimo para chegar ao topo do andar.
func main() {
	// fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	// fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))

	fmt.Println(minCostClimbingStairs2([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs2([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

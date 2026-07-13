package main

import "fmt"

// solução recursiva comum
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

// solução com memoization top-down
func climbStairs2(n int) int {
	memo := make(map[int]int)
	return climbStairsHelper(n, memo)
}

func climbStairsHelper(n int, memo map[int]int) int {
	if n <= 2 {
		return n
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = climbStairsHelper(n-1, memo) + climbStairsHelper(n-2, memo)
	return memo[n]
}

// solução bottom-up
func climbStairs3(n int) int {
	if n <= 2 {
		return n
	}
	prev2, prev1 := 1, 2
	for i := 3; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}
	return prev1
}

// Problema "climbing_stairs" (Adaptado de Leetcode 1971)
// Empresas: Amazon, Google, Facebook, Apple, Microsoft

// Você está subindo uma escada. São necessários n degraus para chegar ao topo.
// A cada vez, você pode subir 1 ou 2 degraus. De quantas maneiras distintas você pode subir até o topo?
func main() {
	// fmt.Println(climbStairs(2))
	// fmt.Println(climbStairs(3))
	fmt.Println(climbStairs2(45))
	fmt.Println(climbStairs3(45))
}

package main

import "fmt"

// solução recursiva padrão
func lcsUtil(ind1, ind2 int, s1, s2 string) int {
	if ind1 < 0 || ind2 < 0 {
		return 0
	}

	if s1[ind1] == s2[ind2] {
		return 1 + lcsUtil(ind1-1, ind2-1, s1, s2)
	} else {
		return max(lcsUtil(ind1-1, ind2, s1, s2), lcsUtil(ind1, ind2-1, s1, s2))
	}
}

func findLCS(s1, s2 string) int {
	n := len(s1)
	m := len(s2)

	return lcsUtil(n-1, m-1, s1, s2)
}

// solução com memoization top-down
func lcsUtil2(ind1, ind2 int, s1, s2 string, memo [][]int) int {
	if ind1 < 0 || ind2 < 0 {
		return 0
	}

	if memo[ind1][ind2] != -1 {
		return memo[ind1][ind2]
	}

	if s1[ind1] == s2[ind2] {
		memo[ind1][ind2] = 1 + lcsUtil2(ind1-1, ind2-1, s1, s2, memo)
		return memo[ind1][ind2]
	} else {
		memo[ind1][ind2] = max(lcsUtil2(ind1-1, ind2, s1, s2, memo), lcsUtil2(ind1, ind2-1, s1, s2, memo))
		return memo[ind1][ind2]
	}
}

func findLCS2(s1, s2 string) int {
	n := len(s1)
	m := len(s2)

	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, m+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return lcsUtil2(n-1, m-1, s1, s2, memo)
}

// solução tabulada bottom-up
func lcsUtil3(s1, s2 string) int {
	n := len(s1)
	m := len(s2)

	grid := make([][]int, n+1)
	for i := range grid {
		grid[i] = make([]int, m+1)
		for j := range grid[i] {
			grid[i][j] = 0
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				grid[i][j] = 1 + grid[i-1][j-1]
			} else {
				grid[i][j] = max(grid[i][j-1], grid[i-1][j])
			}
		}
	}

	return grid[n][m]
}

// Maior Subsequência Comum (LCS - Longest Common Subsequence)
// Dadas duas strings S1 e S2, respectivamente de tamanhos n e m, qual é a maior
// subsequência comum entre elas?
// ● O que é uma subsequência?
// Uma subsequência de uma string é uma lista de caracteres da string original,
// tal que alguns caracteres podem ser deletados, mantendo-se a ordem original.
func main() {
	// fmt.Println(findLCS("abcde", "ace"))
	// fmt.Println(findLCS("abc", "def"))
	// fmt.Println(findLCS("abcdef", "abcdef"))

	fmt.Println(lcsUtil3("abcde", "ace"))
	fmt.Println(lcsUtil3("abc", "def"))
	fmt.Println(lcsUtil3("abcdef", "abcdef"))
	fmt.Println(lcsUtil3("aaaaaaaaaaaaaaaaaaaaaaaaabcdef", "abcdefaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
}

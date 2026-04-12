package main

import "fmt"

func main() {
	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	m2 := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	m3 := matrixMultiply(m1, m2)

	fmt.Printf("%v\n", m3)
}

// exemplo de algorítimo de ordem cúbica
func matrixMultiply(m1, m2 [][]int) [][]int {
	// matrix MxN
	M := len(m1)    // linhas da matriz m1
	N := len(m2[0]) // colunas da matriz m2
	P := len(m2)    // == len(m1[0])

	// alocando as linhas
	m3 := make([][]int, M)

	// alocando colunas para cada linha
	for i := range m3 {
		m3[i] = make([]int, N)
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < P; k++ {
				m3[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}

	return m3
}

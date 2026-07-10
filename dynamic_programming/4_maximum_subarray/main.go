package main

import "fmt"

// solução de força bruta + sliding window (janela deslizante)
// Para cada índice i (início da janela), expande a janela j até o fim do array,
// acumulando a soma e comparando com o melhor resultado já visto.
// Testa todas as O(n²) subarrays possíveis, sem reaproveitar cálculos entre janelas.
// Complexidade: O(n²) tempo, O(1) espaço.
func maxSubArray(v []int) int {
	maxSum := 0
	for i := 0; i < len(v); i++ {
		currentSum := 0
		for j := i; j < len(v); j++ {
			currentSum += v[j]
			maxSum = max(maxSum, currentSum)
		}
	}
	return maxSum
}

// solução de programação dinâmica (top-down convertido em tabulação)
// O problema se resume em, a cada índice i, achar o máximo entre dois números:
// localMax[i] = melhor soma contígua que termina exatamente em i.
// Ou começamos uma nova subarray a partir de v[i], ou estendemos a subarray
// ótima que termina em i-1 somando v[i] a ela — escolhe-se o maior dos dois.
// maxSum guarda o melhor valor entre todos os localMax calculados.
// Complexidade: O(n) tempo, O(n) espaço (por causa do array localMax).
func maxSubArray2(v []int) int {
	localMax := make([]int, len(v))

	// caso base
	localMax[0] = v[0]
	maxSum := v[0]

	for i := 1; i < len(v); i++ {
		localMax[i] = max(v[i], v[i]+localMax[i-1])
		maxSum = max(maxSum, localMax[i])
	}
	return maxSum
}

// algorítimo de kadane
// Mesma lógica de maxSubArray2, mas em vez de guardar localMax[i] para todo i,
// mantém apenas o valor do índice anterior em currentSum (currentSum já é o
// "localMax[i-1]"), eliminando o array auxiliar.
// A cada passo: currentSum = max(começar nova subarray em v[i], estender a anterior).
// Complexidade: O(n) tempo, O(1) espaço — otimização de espaço da versão anterior.
func maxSubArray3(v []int) int {
	maxSum, currentSum := v[0], 0
	for i := 1; i < len(v); i++ {
		currentSum = max(v[i], v[i]+currentSum)
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

// Soma Contígua Máxima de um Array
// Dado um array de inteiros não-nulos, ache o maior valor possível de ser obter com
// uma soma contígua de elementos de s.
// Soma contígua: soma de todos os elementos entre os índices i e j
// Exemplo: soma contígua de 1 até 7
// s(1, 7) = -10 + 2 + 3 + 6 + (-5) + 7 + (-20) = -17
func main() {
	// fmt.Println(maxSubArray([]int{5, -10, 2, 3, 6, -5, 7, -20, 10}))

	// fmt.Println(maxSubArray2([]int{5, -10, 2, 3, 6, -5, 7, -20, 10}))

	fmt.Println(maxSubArray3([]int{5, -10, 2, 3, 6, -5, 7, -20, 10}))
}

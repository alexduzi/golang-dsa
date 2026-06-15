package main

import (
	"fmt"
	"sort"
)

// Problema "garrafas_oasis"
// Em sua jornada pelo deserto, a caravana de João Batista finalmente avistou um oásis e poderão
// parar para matar a sede e encher suas garrafas de água. Suponha que o poço do oásis tenha uma
// quantidade de água X, e que a caravana possui N garrafas de água, cada uma com uma capacidade
// máxima. Encontre a quantidade máxima de garrafas que a caravana poderá encher completamente.
// Entrada
// A entrada será dada no formato de um json contendo os seguintes campos:
// ● “n”: o número de garrafas
// ● “x”: a quantidade de água do oasis
// ● “bottles”: um vetor de tamanho n, onde bottles[i] diz a capacidade da i-ésima garrafa
// Saída
// Imprima o número máximo de garrafas que é possível encher.
func main() {
	fmt.Println(minimumBottles(5, 10, []int{8, 5, 4, 3, 2}))
	fmt.Println(minimumBottles(3, 10, []int{6, 3, 2}))
}

func minimumBottles(n, x int, boottles []int) int {
	sort.Ints(boottles)

	ans := 0

	for i := 0; i < n; i++ {
		if x > 0 && boottles[i] <= x {
			x -= boottles[i]
			ans++
		}
	}

	return ans
}

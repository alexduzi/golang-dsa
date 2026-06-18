package main

import (
	"fmt"
	"math"
)

// Problema "police_and_thieves"
// Dado um vetor de tamanho N(1 ≤ N ≤ 10ହ) tal que cada elemento contém ‘P’ para policial
// ou ‘T’ para ladrão. Ache o número máximo de ladrões que podem ser pegos pela polícia, satisfazendo
// as seguintes condições:
// 1. Cada policial consegue pegar só um ladrão
// 2. Um policial não consegue pegar um ladrão que está a mais de K unidades dele.
// Entrada
// A entrada será dada no formato de um json contendo os seguintes campos:
// ● "n": um inteiro N, indicando o tamanho do vetor
// ● "k": um inteiro K, indicando o alcance dos policiais
// ● "arr": uma lista de N caracteres separados por espaço, indicando se há um policial ou um
// ladrão naquela posição
// Saída
// Imprima o máximo de ladrões que podem ser pegos.
func main() {
	fmt.Println(catchThieves(5, 1, []rune{'P', 'T', 'T', 'P', 'T'}))      // output 2
	fmt.Println(catchThieves(6, 2, []rune{'T', 'T', 'P', 'P', 'T', 'P'})) // output 3
}

func catchThieves(n, k int, positions []rune) int {
	// criando os indices para percorrer de forma eficiente
	thieves := make([]int, 0)
	polices := make([]int, 0)

	// montando os arrays auxiliares
	for i := 0; i < n; i++ {
		if positions[i] == 'T' {
			thieves = append(thieves, i)
		} else if positions[i] == 'P' {
			polices = append(polices, i)
		}
	}

	// começamos mais a esquerda e tentamos parear
	t, p, ans := 0, 0, 0

	for t < len(thieves) && p < len(polices) {
		dist := int(math.Abs(float64(thieves[t] - polices[p])))

		// policial atual consegue pegar nessa distancia
		if dist <= k {
			t++
			p++
			ans++
		} else if thieves[t] < thieves[p] { // ladraoi está muito a esquerda
			t++
		} else {
			p++
		}
	}

	return ans
}

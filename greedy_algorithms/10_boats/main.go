package main

import (
	"fmt"
	"sort"
)

// Problema "boats”
// O vulcão da ilha Pempoia está prestes a entrar em erupção e será preciso evacuar todas as
// pessoas, no entanto, os barcos não são muito resistentes e aguentam um peso de no máximo K. Além
// disso, cada barco consegue carregar no máximo duas pessoas, desde que a soma de seus pesos esteja
// dentro do limite.
// Dada uma lista dos pesos dos habitantes da ilha, diga qual o número mínimo de barcos para que
// ela possa ser evacuada imediatamente. É garantido que toda pessoa poderá ser salva.
// Entrada
// A entrada será dada no formato de um json contendo os seguintes campos:
// ● "people": uma lista de inteiros representando o peso dos habitantes da ilha
// ● "limit": um inteiro K, indicando o limite de peso que cada barco pode carregar
// Saída
// Imprima o número mínimo de barcos para carregar todas as pessoas.
// Explicação: 4 barcos: (3), (3), (4) e (5)
func main() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	i, j, boats := 0, len(people)-1, 0

	for i <= j {
		if people[i]+people[j] <= limit {
			i++
			j--
		} else if people[j] <= limit {
			j--
		} else {
			i++
		}
		boats++
	}

	return boats
}

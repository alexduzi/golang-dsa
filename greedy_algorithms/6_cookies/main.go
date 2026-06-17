package main

import (
	"fmt"
	"sort"
)

// Problema "cookies"
// Suponha que você é um pai incrível e quer dar às suas crianças alguns cookies. Mas você dará
// a cada criança no máximo um cookie.
// Cada criança i tem um fator guloso g[i], que é o tamanho mínimo de cookie para ela ficar
// satisfeita, e cada cookie j tem um tamanho s[j]. Se o tamanho do cookie for maior ou igual ao fator
// guloso da criança, então ela ficará satisfeita.
// Distribuindo os cookies de maneira ótima, qual é o número máximo de crianças satisfeitas?
// Entrada
// A entrada será dada no formato de um json contendo os seguintes campos:
// ● "g": uma lista de valores inteiros, onde g[i] representa o fator guloso da i-ésima criança
// ● "s": uma lista de valores inteiros, onde s[i] representa o tamanho do i-ésimo cookie
// Saída
// Imprima a quantidade máxima de crianças satisfeitas, ao distribuir os cookies de maneira ótima.
func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{1, 2}, []int{2, 3, 1}))
}

func findContentChildren(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	j, ans := 0, 0

	// percorrer e distribuir cookies
	for i := 0; i < len(g); i++ {
		// ainda existe cookie e cookie não satisfaz a criança
		for j < len(s) && s[j] < g[i] {
			j++
		}

		// não consegue satisfazer mais crianças / acabaram os cookies
		if j >= len(s) {
			break
		} else { // achamos o cookie
			ans++
			j++
		}
	}

	return ans
}

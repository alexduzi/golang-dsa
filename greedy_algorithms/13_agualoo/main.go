package main

import (
	"fmt"
	"sort"
)

// Certo dia o Reino de Agualoo foi atacado por um dragão cuspidor de fogo de várias cabeças,
// ameaçando queimar o reino até as cinzas. Preocupado, o rei chamou seus cavaleiros para derrotar o
// dragão e salvar o reino.
// Os cavaleiros explicaram: “Para derrotar o dragão, temos que cortar todas as suas cabeças. Cada
// cavaleiro consegue cortar só uma das cabeças do dragão. A ordem dos cavaleiros requer que, por
// cortar uma cabeça, um cavaleiro seja pago uma recompensa igual a uma moeda de ouro para cada
// centímetro da altura do cavaleiro.
// Haverá cavaleiros o suficiente para derrotar o dragão? O rei chamou seus conselheiros para ajudar a
// decidir quantos e quais cavaleiros contratar, minimizando o custo de derrotar o dragão. Como um de
// seus conselheiros, você deve ajudar o rei. Isto é muito sério: se você falhar, o reino inteiro será
// reduzido a cinzas!
// Entrada
// n: o número de cabeças do dragão.
// m: o número de cavaleiros no reino.
// diameters: uma lista de inteiros que representam o diâmetro das cabeças do dragão, em
// centímetros.
// heights: uma lista de inteiros que especifica a altura dos cavaleiros de Agualoo, também em
// centímetros.
// Saída
// Imprima uma linha contendo o número mínimo de moedas de ouro que o rei precisará pagar para
// derrotar o dragão. Se não for possível que os cavaleiros de Agualoo derrotem o dragão, imprima a
// linha “Agualoo esta condenada!”.

func main() {
	agualoo(2, 3, []int{5, 4}, []int{7, 8, 4})

	agualoo(2, 1, []int{5, 5}, []int{10})

	agualoo(2, 4, []int{7, 2}, []int{4, 3, 1, 2})

	agualoo(2, 4, []int{7, 2}, []int{2, 1, 8, 5})

	agualoo(2, 10, []int{1234567, 2345}, []int{12345610, 1, 123, 23564,
		123456, 123, 2, 3, 2, 1})
}

func agualoo(n, m int, diameters, heights []int) {
	if n > m {
		fmt.Println("Agualoo esta condenada!")
		return
	}

	sort.Ints(diameters)
	sort.Ints(heights)

	coins, dIdx, hIdx := 0, 0, 0

	for (n > 0 && m > 0) && (dIdx < len(diameters)) {
		if heights[hIdx] >= diameters[dIdx] {
			n--
			m--
			coins += heights[hIdx]
			hIdx++
			dIdx++
		} else { // se o cavaleiro atual não der conta, pula para o prox cavaleiro
			hIdx++
			m--
		}
	}

	if n > 0 {
		fmt.Println("Agualoo esta condenada!")
	} else {
		fmt.Println(coins)
	}
}

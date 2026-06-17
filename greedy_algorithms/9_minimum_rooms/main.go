package main

import (
	"fmt"
	"sort"
)

// Problema "minimum_rooms"
// Você é o responsável por organizar a infraestrutura de um evento que contará com muitas
// palestras interessantes e elas já estão programadas.
// Dado o tempo de começo e de fim de cada palestra, você deve determinar qual o número
// mínimo de salas que devem ser reservadas para que todas as palestras aconteçam no horário
// programado, considerando que duas palestras não ocorrem na mesma sala. Se uma palestra começar
// no momento que outra acabar, não podemos usar a mesma sala. Nesses casos, precisamos de outra
// reserva.
// Entrada
// Entrada A entrada será dada no formato de um json contendo os seguintes campos:
// ● "start": uma lista de inteiros representando os tempos de início das palestras
// ● "end": uma lista de inteiros representando os tempos de fim das palestras
// ● "n": um inteiro N, indicando o número de palestras no cronograma do evento
// Saída
// Imprima o número mínimo de salas que devem ser reservadas para que as palestras ocorram no
// horário programado.
func main() {
	fmt.Println(minimumRooms([]int{900, 940, 950, 1100, 1500, 1800}, []int{910, 1200, 1120, 1130, 1900, 2000}, 6)) // output 3
	fmt.Println(minimumRooms([]int{900, 1100, 1235}, []int{1000, 1200, 1240}, 3))                                  // output 1
}

func minimumRooms(start, end []int, n int) int {
	sort.Ints(start)
	sort.Ints(end)

	rooms, ans, i, j := 0, 0, 0, 0

	for i < n || j < n {
		// decidir se alocamos ou desalocamos uma sala
		if i < n && (j >= n || start[i] <= end[j]) {
			rooms++
			i++
		} else {
			rooms--
			j++
		}

		ans = max(ans, rooms)
	}

	return ans
}

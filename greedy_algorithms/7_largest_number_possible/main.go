package main

import (
	"fmt"
)

// Problema "largest_number_possible"
// Dados dois inteiros N e S, ache o maior número que pode ser formado com N dígitos e cuja
// soma desses dígitos seja igual a S. Imprima -1 se não for possível.
// Entrada
// A entrada será dada no formato de um json contendo os seguintes campos:
// ● "n": um inteiro N representando a quantidade de dígitos do número a ser formado
// ● "s": um inteiro S representando a soma dos dígitos do número
// Saída
// Imprima o maior número possível.
func main() {
	fmt.Println(findLargest(2, 9))  // output 90
	fmt.Println(findLargest(3, 20)) // output 992
	fmt.Println(findLargest(4, 0))  // output -1
}

func findLargest(n, s int) string {

	// impossível: soma maior que o máximo (9*n), ou soma zero com mais de 1 dígito (geraria "00...")
	if (9*n) < s || (s == 0 && n > 1) {
		return "-1"
	}

	ans := ""

	// estratégia gulosa: em cada posição, coloca o maior dígito possível (9)
	// para maximizar o número, priorizamos os dígitos mais significativos
	for s > 0 {
		digit := ""

		if (s - 9) >= 0 {
			// ainda há soma suficiente para colocar 9 nesta posição
			s -= 9
			digit = "9"
		} else {
			// soma restante é menor que 9: usa tudo o que sobrou nesta posição
			digit = string(rune('0') + int32(s))
			s = 0
		}
		ans += digit
	}

	// preenche com zeros à direita até atingir n dígitos
	for len(ans) < n {
		ans += "0"
	}

	return ans
}

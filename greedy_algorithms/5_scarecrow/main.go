package main

import "fmt"

// Problema "scarecrow"
// Ritinha tem uma propriedade muito grande onde plantará várias sementes, porém, a área está
// cheia de corvos, e Ritinha teme que eles irão comer suas sementes. Por isso, ela decidiu colocar
// espantalhos em algumas localizações do campo.
// O campo pode ser visto como um grid de 1 × N posições. Algumas áreas, marcadas por uma
// cerquilha (#), são inférteis e não se pode plantar nelas, enquanto as áreas férteis são marcadas por um
// ponto (. ). Quando um espantalho é colocado em uma posição, ele protege aquela posição e as
// adjacentes.
// Dada a descrição do campo, qual é o número mínimo de espantalhos que precisam ser colocados
// para proteger toda a área fértil?
// Entrada
// A entrada será dada no formato de um json contendo os seguintes campos:
// ● "n": um inteiro N (0 < N < 100) representando o comprimento do campo
// ● "field": uma string de N caracteres que descrevem o campo. Um ponto (.) indica uma área
// fértil e uma cerquilha (#) indica uma área infértil.
// Saída
// Imprima o número mínimo de espantalhos que precisam ser colocados.
func main() {
	fmt.Println(minimumScarecrows(3, ".#."))
	fmt.Println(minimumScarecrows(11, "...##....##"))
	fmt.Println(minimumScarecrows(2, "##"))
}

func minimumScarecrows(n int, field string) int {

	ans := 0

	// percorrer o campo
	for i := 0; i < n; i++ {
		// achar ponto fértil
		if field[i] == '.' {
			ans++
			i += 2
		}
	}

	return ans
}

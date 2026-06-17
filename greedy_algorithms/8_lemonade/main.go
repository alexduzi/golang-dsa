package main

import "fmt"

// Problema "lemonade"
// Em uma barraquinha de limonada, cada limonada custa 5 reais. Os clientes estão em fila
// para comprar, um de cada vez. Cada cliente comprará apenas uma limonada e pagará com uma
// nota de 5, 10 ou 20 reais. Você deve dar o troco correto para que o cliente seja cobrado apenas
// 5 reais.
// No entanto, no começo você não tem nenhuma nota para dar o troco. Você só poderá
// dar notas que conseguir com os clientes anteriores.
// Dado um vetor de inteiros de tamanho N, onde cada elemento nessa ordem é a nota que
// o i-ésimo consumidor pagará, imprima “Verdadeiro”, se você consegue dar o troco correto para
// todos os clientes, ou “Falso” caso contrário.
// Entrada
// Entrada A entrada será dada no formato de um json contendo os seguintes campos:
// ● "bills": uma lista de N inteiros, cada um indicando a nota que o i-ésimo cliente pagará
// Saída
// Imprima “Verdadeiro”, se for possível dar o troco correto para todos os clientes, ou
// “Falso” caso contrário.
// Explicação: Dos primeiros dois clientes na fila, coletamos duas notas de 5 reais. Dos próximos dois
// clientes na ordem, coletamos uma nota de 10 e damos uma nota de 5. Para o último cliente, não
// podemos dar 15 reais de volta, porque temos apenas duas notas de 10.
func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20}))
}

func lemonadeChange(bills []int) bool {
	// caixa: quantidade de notas disponíveis para troco
	myBills := map[int]int{
		20: 0,
		10: 0,
		5:  0,
	}

	notes := []int{20, 10, 5}

	for i := 0; i < len(bills); i++ {
		bill := bills[i]
		myBills[bill]++ // recebe a nota do cliente

		change := bill - 5 // troco necessário (limonada custa 5)

		// ATENÇÃO: iterar sobre map em Go não garante ordem.
		// O ideal seria iterar das notas maiores para as menores (guloso),
		// mas aqui a ordem aleatória pode produzir resultados incorretos.
		for _, value := range notes {
			// usa notas de valor k enquanto tiver disponível e o troco precisar
			for myBills[value] > 0 && change >= value {
				myBills[value]--
				change -= value
			}
		}

		// se ainda resta troco a dar, não foi possível atender este cliente
		if change > 0 {
			return false
		}
	}

	return true
}

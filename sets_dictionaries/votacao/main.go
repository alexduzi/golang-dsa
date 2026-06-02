package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Na contagem de votos de uma eleição, são gerados vários registros de votação contendo o nome do candidato
// e a quantidade de votos (formato .csv) que ele obteve em uma urna de votação.
// Você deve fazer uma função para receber todos registros de votação das urnas,
// e daí retornar um consolidado com os totais de cada candidato. O resultado pode ser mostrado em qualquer ordem.
func main() {
	votes := []string{
		"Alex Blue,15",
		"Maria Green,22",
		"Bob Brown,21",
		"Alex Blue,30",
		"Bob Brown,15",
		"Maria Green,27",
		"Maria Green,22",
		"Bob Brown,25",
		"Alex Blue,31",
	}

	result := counting(votes)

	fmt.Printf("Result: %v\n", result)
}

func counting(votes []string) []string {
	m := make(map[string]int)

	for _, v := range votes {
		data := strings.Split(v, ",")
		name := data[0]
		vote, _ := strconv.Atoi(data[1])

		m[name] += vote
	}

	result := make([]string, 0, len(m))
	for key, value := range m {
		result = append(result, fmt.Sprintf("%s,%d", key, value))
	}
	return result
}

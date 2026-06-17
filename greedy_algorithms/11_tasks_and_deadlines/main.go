package main

import (
	"fmt"
	"slices"
)

// Problema "tasks_and_deadlines"
// Trabalhando como freelancer, você conseguiu N(1 ≤ N ≤ 2 ∗ 10ହ) projetos. Cada projeto tem
// uma duração e um prazo, e você deverá escolher uma ordem para os executar. Sua recompensa por
// cada projeto é dada por d − f, onde d é seu prazo e f é o seu tempo de término. (No começo o tempo
// é 0, e você terá que fazer todos os projetos, mesmo que tenha recompensa negativa.)
// Qual será sua recompensa ao fim de todos os projetos, executando-os de forma ótima?
// Entrada
// A entrada será dada no formato de um json contendo os seguintes campos:
// ● "tasks": uma lista de listas, onde cada sublista contém dois inteiros representando a duração e
// o prazo de cada projeto
// Saída
// Imprima qual a recompensa máxima que pode ser obtida.
func main() {
	fmt.Println(maximumReward([]Task{
		{6, 10},
		{8, 15},
		{5, 12},
	}))

	fmt.Println(maximumReward([]Task{
		{3, 47},
		{5, 11},
		{1, 70},
		{2, 100},
		{2, 41},
		{2, 66},
		{5, 80},
		{4, 84},
		{5, 81},
		{5, 40},
	}))

	fmt.Println(maximumReward([]Task{
		{80, 55},
		{29, 46},
		{58, 5},
		{92, 80},
		{62, 68},
		{64, 20},
		{78, 56},
		{41, 66},
		{62, 44},
		{32, 80},
	}))
}

type Task struct {
	first  int
	second int
}

func maximumReward(tasks []Task) int {
	// ordena por duração crescente: tarefas mais curtas primeiro minimizam o tempo acumulado
	slices.SortFunc(tasks, func(a, b Task) int {
		return a.first - b.first
	})

	time, reward := 0, 0

	for i := 0; i < len(tasks); i++ {
		time += tasks[i].first // tempo de término da tarefa atual

		// recompensa = prazo - tempo de término (pode ser negativa)
		reward += tasks[i].second - time
	}

	return reward
}

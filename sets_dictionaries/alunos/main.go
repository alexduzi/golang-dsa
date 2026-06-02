package main

import "fmt"

// Em um portal de cursos online, cada aluno possui um código único, representado por um número inteiro.

// Cada instrutor do portal pode ter vários cursos, sendo que um mesmo aluno pode se matricular em quantos cursos quiser.

// Assim, o número total de alunos de um instrutor não é simplesmente a soma dos alunos de todos os cursos que ele possui,
// pois pode haver alunos repetidos em mais de um curso.

// Você deve criar uma função que, dada a lista de alunos de todos os cursos de um instrutor, a função deve retornar a
// quantidade total de alunos deste instrutor.
func main() {
	courses := [][]int{
		{15, 21, 80, 42},
		{21, 80, 47},
		{12, 21, 47, 35},
	}

	amountStudents := studentsCount(courses)

	fmt.Printf("Amount of students: %d\n", amountStudents)
}

func studentsCount(courses [][]int) int {
	set := NewSet()
	for _, line := range courses {
		for _, val := range line {
			set.Add(val)
		}
	}
	return set.Size()
}

// Set representa um conjunto de elementos únicos usando um mapa internamente
type Set struct {
	elements map[any]struct{}
}

// NewSet cria e retorna um novo conjunto vazio
func NewSet() Set {
	return Set{
		elements: make(map[any]struct{}),
	}
}

// Add insere um elemento no conjunto (ignorado se já existir)
func (set *Set) Add(elem any) {
	set.elements[elem] = struct{}{}
}

// Remove exclui um elemento do conjunto
func (set *Set) Remove(elem any) {
	delete(set.elements, elem)
}

// IsEmpty retorna true se o conjunto não contiver nenhum elemento
func (set *Set) IsEmpty(elem any) bool {
	return set.Size() == 0
}

// Size retorna a quantidade de elementos no conjunto
func (set *Set) Size() int {
	return len(set.elements)
}

// List retorna todos os elementos do conjunto em um slice
func (set *Set) List() (list []any) {
	for k := range set.elements {
		list = append(list, k)
	}
	return
}

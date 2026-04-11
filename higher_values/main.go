package main

import "fmt"

func main() {
	fmt.Printf("%v\n", higherValues([]int{7, 3, 8, 7, 5}))
}

// Função que recebe um vetor de números, e retorna um novo vetor dizendo quantos elementos
// maiores existem no vetor, para cada elemento do vetor
func higherValues(arr []int) []int {
	newArr := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] > arr[i] {
				newArr[i]++
			}
		}
	}

	return newArr
}

package main

import "fmt"

func main() {
	nums := []int{20, 9, 86, -2, 16, 13, 34, 4}

	fmt.Printf("%v\n", insertionSort(nums))
}

// algoritmo mais eficiente que merge sort e quick sort
// esse algorítimo sempre verifica os elementos a esquerda e move para a direita
func insertionSort(nums []int) []int {
	N := len(nums)
	// começa na segunda posição
	for i := 1; i < N; i++ {
		aux := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > aux {
			nums[j+1] = nums[j]
			j -= 1
		}
		nums[j+1] = aux
	}
	return nums
}

package main

import "fmt"

func main() {
	nums := []int{20, 9, 86, -2, 16, 13, 34, 4}
	fmt.Printf("%v\n", selectionSort(nums))
}

func selectionSort(nums []int) []int {
	N := len(nums)

	for i := 0; i < N-1; i++ {
		minIndex := i
		for j := i + 1; j < N; j++ {
			if nums[j] < nums[minIndex] { // achou um menor elemento que i
				minIndex = j
			}
		}
		if minIndex != i { // se o minIndex for diferente de i é pq achou um elemento menor, então nesse caso faz a troca
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}

	return nums
}

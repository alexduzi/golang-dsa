package main

import "fmt"

// Dados dois arrays, calcule o seu produto escalar.
// entrada: nums1 = [1,0,0,2,3], nums2 = [0,3,0,4,0]
// saída: 8
// Explicação: O produto escalar dos arrays acima pode ser encontrado pela expressão: (1 * 0) + (0 * 3) + (0 * 0) + (2 * 4) + (3 * 0) = 8
// entrada: nums1 = [0,1,0,0,0], nums2 = [0,0,0,0,2]
// saída: 0
func main() {
	fmt.Printf("%d\n", dotProduct([]int{1, 0, 0, 2, 3}, []int{0, 3, 0, 4, 0}))
	fmt.Printf("%d\n", dotProduct([]int{0, 1, 0, 0, 0}, []int{0, 0, 0, 0, 2}))
}

func dotProduct(nums1, nums2 []int) int {
	product := 0
	for i := 0; i < len(nums1); i++ {
		product += nums1[i] * nums2[i]
	}
	return product
}

package main

import (
	"fmt"
	"strconv"
)

// Dado um array nums de inteiros, retorne quantos deles contêm um número par de dígitos.
// entrada: nums = [12,345,2,6,7896]
// saída: 2
// Explicação:
// 12 contém 2 dígitos (número par de dígitos).
// 345 contém 3 dígitos (número ímpar de dígitos).
// 2 contém 1 dígito (número ímpar de dígitos).
// 6 contém 1 dígito (número ímpar de dígitos).
// 7896 contém 4 dígitos (número par de dígitos).
// Portanto, apenas 12 e 7896 contêm um número par de dígitos.
// entrada: nums = [555,901,482,1771]
// saída: 1
func main() {
	fmt.Printf("%d\n", findNumbers([]int{12, 345, 2, 6, 7896}))
	fmt.Printf("%d\n", findNumbers([]int{555, 901, 482, 1771}))

	fmt.Printf("%d\n", findNumbers2([]int{12, 345, 2, 6, 7896}))
	fmt.Printf("%d\n", findNumbers2([]int{555, 901, 482, 1771}))
}

func findNumbers(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if len(strconv.Itoa(nums[i]))%2 == 0 { // método mais prático é converter o número em string
			result++
		}
	}
	return result
}

func findNumbers2(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if countDigits(nums[i])%2 == 0 {
			result++
		}
	}
	return result
}

func countDigits(num int) (result int) {
	for num > 0 {
		num /= 10 // outro metodo é contar os números pela divisão sucessiva por 10
		result++
	}
	return
}

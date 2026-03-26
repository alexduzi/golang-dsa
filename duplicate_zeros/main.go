package main

import "fmt"

// Dado um array de inteiros arr, duplique cada ocorrência de zero, deslocando os elementos restantes para a direita (shifting right).
// Observe que elementos além do tamanho do array original não são escritos.
// Utilize a abordagem "In-place" na qual a modificação é feita diretamente no array.
// entrada: nums = [1,0,2,3,0,4,5,0]
// saída: [1,0,0,2,3,0,0,4]
// entrada: nums = [1,2,3]
// saída: [1,2,3]
func main() {
	nums := []int{1, 0, 2, 3, 0, 4, 5, 0}
	nums2 := []int{1, 2, 3}
	duplicateZeros(nums)
	duplicateZeros(nums2)
	fmt.Printf("%v", nums)
	fmt.Printf("%v", nums2)
}

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			for j := len(arr) - 2; j > i; j-- { // começa da penúltima posição, esse for irá percorrer o sentido inverso
				arr[j+1] = arr[j] // deslocando os elementos para a direita
			}
			arr[i+1] = arr[i] // coloca o zero na prox posição
			i = i + 1
		}
	}
}

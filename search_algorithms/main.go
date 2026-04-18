package main

import "fmt"

func main() {
	numbers := []int{1, 20, 30, 5, 10, 17, 15}
	index := search(numbers, 10)
	fmt.Printf("%v", index)
}

// busca sequencial
// complexidade O(n)
func search(arr []int, e int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == e {
			return i
		}
	}
	return -1
}

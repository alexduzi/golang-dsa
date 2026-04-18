package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := []int{1, 20, 30, 5, 10, 17, 15}
	index := binarySearchIter(numbers, 10)
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

// busca binária
func binarySearchIter(arr []int, e int) int {
	low := 0
	middle := 0
	high := len(arr) - 1

	for low <= high {
		middle = int(math.Floor(float64((low + high) / 2)))
		if e < arr[middle] {
			high = middle - 1
		} else if e > arr[middle] {
			low = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

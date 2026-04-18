package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := []int{1, 20, 30, 5, 10, 17, 15}
	index := binarySearchIter(numbers, 10)
	fmt.Printf("%v\n", index)

	index2 := binarySearchRecursive(numbers, 10, 0, len(numbers)-1)
	fmt.Printf("%v\n", index2)
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
// O(long n)
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

func binarySearchRecursive(arr []int, e, low, high int) int {

	if low > high {
		return -1
	}

	middle := int(math.Floor(float64((low + high) / 2)))

	if e == arr[middle] {
		return middle
	} else if e < arr[middle] {
		return binarySearchRecursive(arr, e, low, middle-1)
	} else {
		return binarySearchRecursive(arr, e, middle+1, high)
	}
}

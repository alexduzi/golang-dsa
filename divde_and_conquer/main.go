package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	counter := countRecursive(arr)
	fmt.Printf("counter: %v\n", counter)

	sum := sumRecursive(arr)
	fmt.Printf("sum: %v\n", sum)

	arr2 := []int{3, 6, 8, 10, 1, 2, 1}
	sorted := quickSort(arr2)
	fmt.Println(sorted)

	arr3 := []int{3, 6, 8, 10, 1, 2, 1}
	mergeS := mergeSort(arr3)
	fmt.Println(mergeS)
}

func countRecursive(arr []int) int {
	size := len(arr)
	if size == 0 || size == 1 {
		return 1
	}
	return 1 + countRecursive(arr[1:])
}

func sumRecursive(arr []int) int {
	size := len(arr)
	if size == 0 {
		return 0
	}
	if size == 2 {
		return arr[0] + arr[1]
	}
	return arr[0] + sumRecursive(arr[1:])
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less, high []int

	for _, value := range arr[1:] {
		if value <= pivot {
			less = append(less, value)
		} else {
			high = append(high, value)
		}
	}

	result := quickSort(less)
	result = append(result, pivot)
	result = append(result, quickSort(high)...)
	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

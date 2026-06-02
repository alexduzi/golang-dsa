package main

import (
	"fmt"
)

// ref: Leetcode intersection-of-two-arrays
// Dados dois arrays num1 e num2, retorne um array contendo sua interseção.
// Cada elemento do array resultante deve ser único, e você pode apresentar os elementos em qualquer ordem.
func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	result := intersection(nums1, nums2)

	fmt.Printf("Result 1: %v\n", result)

	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}

	result = intersection(nums1, nums2)

	fmt.Printf("Result 1: %v\n", result)
}

func intersection(nums1, nums2 []int) []int {
	setNums1 := make(map[int]struct{})
	setNums2 := make(map[int]struct{})
	setResult := make(map[int]struct{})

	for _, v := range nums1 {
		setNums1[v] = struct{}{}
	}

	for _, v := range nums2 {
		setNums2[v] = struct{}{}
	}

	for key := range setNums1 {
		if _, ok := setNums2[key]; ok {
			setResult[key] = struct{}{}
		}
	}

	result := make([]int, 0, len(setResult))
	for key := range setResult {
		result = append(result, key)
	}

	return result
}

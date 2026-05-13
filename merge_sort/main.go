package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{20, 9, 86, -2, 16, 13, 34, 4}
	fmt.Printf("%v\n", mergeSort(nums, 0, 7))
}

func mergeSort(nums []int, left, right int) []int {
	if left < right {
		middle := int(math.Floor(float64(left+right) / 2))
		mergeSort(nums, left, middle)
		mergeSort(nums, middle+1, right)
		merge(&nums, left, middle, right)
	}
	return nums
}

func merge(nums *[]int, left, middle, right int) {
	result := make([]int, right-left+1)
	length := right - left + 1
	i := left
	j := middle + 1
	k := 0

	for i <= middle && j <= right {
		if (*nums)[i] < (*nums)[j] {
			result[k] = (*nums)[i]
			k++
			i++
		} else {
			result[k] = (*nums)[j]
			k++
			j++
		}
	}

	for i <= middle {
		result[k] = (*nums)[i]
		k++
		i++
	}

	for j <= right {
		result[k] = (*nums)[j]
		k++
		j++
	}

	for i := 0; i < length; i++ {
		(*nums)[left+i] = result[i]
	}
}

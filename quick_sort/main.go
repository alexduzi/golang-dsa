package main

import (
	"fmt"
)

func main() {
	nums := []int{20, 9, 86, -2, 16, 13, 34, 4}
	fmt.Printf("%v\n", quickSort(nums, 0, 7))
}

func quickSort(nums []int, left, right int) []int {
	if left < right {
		pivot := partition(nums, left, right)
		quickSort(nums, left, pivot-1)
		quickSort(nums, pivot+1, right)
	}
	return nums
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

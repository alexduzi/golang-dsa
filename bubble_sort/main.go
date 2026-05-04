package main

import "fmt"

func main() {
	fmt.Printf("%v\n", bubbleSort([]int{10, 20, 5, 3, 1}))
}

func bubbleSort(nums []int) []int {
	N := len(nums)

	for i := 0; i < N; i++ {
		for j := 0; j < N-1; j++ {
			if nums[j] > nums[j+1] {
				// aux := nums[j]
				// nums[j] = nums[j+1]
				// nums[j+1] = aux
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

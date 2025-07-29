package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
}

func searchInsert(nums []int, target int) int {
	var result int
	for low, high := 0, len(nums)-1; low <= high; {
		mid := (high + low) / 2
		n := nums[mid]
		switch {
		case n < target:
			result = mid + 1
			low = mid + 1
		case n > target:
			result = mid
			high = mid - 1
		default:
			result = mid
			return result
		}
	}
	return result
}

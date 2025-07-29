package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	writeIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[writeIndex] = nums[i]
			writeIndex++
		}

	}
	return writeIndex
}

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d\n", findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	fmt.Printf("%d\n", findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	maxOnes := 0   // maior sequencia de 1s
	countOnes := 0 // sequencia de 1s atual
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			countOnes++
		}
		if nums[i] != 1 && countOnes > 0 {
			if countOnes > maxOnes {
				maxOnes = countOnes
			}
			countOnes = 0
		}
	}
	if countOnes > maxOnes {
		maxOnes = countOnes
	}
	return maxOnes
}

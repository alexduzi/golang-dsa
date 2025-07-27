package main

import "fmt"

func main() {
	fmt.Println(string(10))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	hasher := map[int]int{}
	for index, value := range nums {
		if idxM, ok := hasher[value]; ok {
			return []int{idxM, index}
		}
		hasher[target-value] = index
	}
	return []int{}
}

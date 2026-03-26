package main

import (
	"fmt"
	"sort"
)

// Dado um array de números inteiros nums, retorne true se houver
// valores repetidos ou false se não houver repetição de valores no array.
// entrada: nums = [1,2,3,1]
// saída: true
// entrada: nums = [1,2,3,4]
// saída: false
// entrada: nums = [1,1,1,3,3,4,3,2,4,2]
// saída: true
func main() {
	fmt.Printf("%v\n", containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Printf("%v\n", containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Printf("%v\n", containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

func containsDuplicate(nums []int) bool {
	hash := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			return true
		}
		hash[nums[i]] += 1
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	hash := make(map[int]int, len(nums))
	for _, num := range nums {
		if _, ok := hash[num]; ok {
			return true
		}
		hash[num] += 1
	}
	return false
}

func containsDuplicate3(nums []int) bool {
	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			return true
		}
		hash[nums[i]] += 1
	}
	return false
}

func containsDuplicate4(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}

func containsDuplicate5(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

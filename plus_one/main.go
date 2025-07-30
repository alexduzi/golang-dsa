package main

import "fmt"

/*
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer.
The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.

Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
*/

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9, 9}))
	fmt.Println(plusOne([]int{9, 9, 9, 9, 9}))
}

func plusOne(digits []int) []int {
	size := len(digits)
	if size == 1 && digits[0] == 9 {
		return []int{1, 0}
	} else if size == 1 && digits[0] < 9 {
		sumNumber := digits[0] + 1
		return []int{sumNumber}
	} else {
		counterNine := 0
		for i := size - 1; i >= 0; i-- {
			if digits[i] == 9 {
				digits[i] = 0
				counterNine++
			} else {
				digits[i] += 1
				break
			}

			if counterNine > 1 && i == 0 {
				newSlice := append([]int{1}, digits...)
				return newSlice
			}
		}
		return digits
	}
}

package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 0, 3}

	result := missingNumber(arr)
	fmt.Println(result)
	// for _, value := range result {
	// 	fmt.Printf("%d \t", value)
	// }
}

func missingNumber(arr []int) int {
	n := len(arr)
	sum := n * (n + 1) / 2

	for _, value := range arr {
		sum -= value
	}

	return sum
}

func moveZeros(arr []int) []int {
	aux := 0
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		if arr[i] != 0 && arr[aux] == 0 {
			arr[i], arr[aux] = arr[aux], arr[i]
		}
		if arr[aux] != 0 {
			aux++
		}
	}
	return arr
}

func moveZeros2(arr []int) []int {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		if arr[i] == 0 {
			arr = append(arr[:i], arr[i+1:]...)
			arr = append(arr, 0)
			i -= 1
			arrLen -= 1
		}
	}
	return arr
}

func reverseArray(arr []int) []int {
	start := 0
	end := len(arr) - 1

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}

	return arr
}

func removeEvenNumbers(arr []int) (oddNumbers []int) {
	for i := 0; i < len(arr); i++ {
		if (arr[i] % 2) != 0 {
			oddNumbers = append(oddNumbers, arr[i])
		}
	}
	return
}

func printArray(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%d \t", arr[i])
	}
}

func printSlice(slc [6]int) {
	for _, value := range slc {
		fmt.Printf("%v \t", value)
	}
}

func findSum(n int) int {
	return n * (n + 1) / 2
}

func findSumInterate(n int) (sum int) {
	for i := 0; i < n; i++ {
		sum += i
	}
	return
}

func fat(n int) (fat int) {
	fat = 1
	if n == 0 || n == 1 {
		return
	}
	for i := 1; i <= n; i++ {
		fat *= i
	}
	return
}

func fatRecursive(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * fatRecursive(n-1)
}

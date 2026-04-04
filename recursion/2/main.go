package main

import "fmt"

func main() {
	fmt.Printf("%d\n", fat2(3))
	fmt.Printf("%d\n", fat2(5))
	fmt.Printf("%d\n", fat2(20))
	fmt.Printf("%d\n", fat2(17))
}

func fat(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * fat(n-1)
}

func fat2(n int) int {
	return fatTailRecursive(n, 1)
}

func fatTailRecursive(n, total int) int {
	if n == 0 {
		return total
	}

	return fatTailRecursive(n-1, n*total)
}

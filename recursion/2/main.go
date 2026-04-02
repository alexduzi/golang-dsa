package main

import "fmt"

func main() {
	fmt.Printf("%d\n", fat(3))
	fmt.Printf("%d\n", fat(5))
	fmt.Printf("%d\n", fat(20))
	fmt.Printf("%d\n", fat(17))
}

func fat(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * fat(n-1)
}

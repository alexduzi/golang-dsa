package main

import "fmt"

func main() {
	fmt.Printf("%d\n", fib2(47))
}

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func fib2(n int) int {
	return fibTailRecursive(n, 0, 1)
}

func fibTailRecursive(n, a, b int) int {
	if n == 0 {
		return a
	}

	if n == 1 {
		return b
	}

	return fibTailRecursive(n-1, b, a+b)
}

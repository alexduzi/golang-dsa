package main

import "fmt"

func main() {
	fmt.Printf("%d\n", sumNaturals(5))
	fmt.Printf("%d\n", sumNaturals(20))
	fmt.Printf("%d\n", sumNaturals(17))
}

func sumNaturals(n int) int {
	if n == 0 {
		return n
	}

	return n + sumNaturals(n-1)
}
